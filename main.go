package main

import (
	"fmt"
	"crypto/tls"
	"os"
	"github.com/quic-go/quic-go/http3"
	"github.com/sirupsen/logrus"
	"github.com/MyNameIsRaphi/web_proxy/forward"
)
const PORT = 8080




func main(){

	key_path, exists := os.LookupEnv("KEY_PATH")
	if !exists{
		logrus.Fatal("Failed to read KEY_PATH")
	} 
	cert_path, exists := os.LookupEnv("CERT_PATH")
	if !exists {
		logrus.Fatal("Failed to read CERT_PATH")
	}
	cert, err := tls.LoadX509KeyPair(cert_path, key_path)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to inialize tls")
	}
	var tlsConfig = tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.VerifyClientCertIfGiven}
	
	
	var addr string = fmt.Sprintf(":%d", PORT)
	handler := forward.Handler{}	
	handler.TlsConfig = &tlsConfig

	err = http3.ListenAndServe(addr, cert_path, key_path, handler)	

	if err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}



}



