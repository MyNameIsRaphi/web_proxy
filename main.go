package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
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
	logrus.Infof("Started server on port %d", PORT)
	cert, err := tls.LoadX509KeyPair(cert_path, key_path)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to inialize tls")
	}
	var tlsConfig = tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.VerifyClientCertIfGiven}
	var port_string = fmt.Sprintf(":%d", PORT)
	server, err := net.Listen("tcp", port_string)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to create server on port %s", port_string) 
	}
	for {
		conncetion, conn_err := server.Accept()
		if conn_err != nil {
			logrus.WithError(conn_err).Error("Failed to accept incoming conntecion")
			continue
		}
		defer conncetion.Close()
		tlsConnection := tls.Client(conncetion, &tlsConfig)
		handshake_err := tlsConnection.Handshake()
		if handshake_err != nil {
			logrus.WithError(handshake_err).Error("TLS handshake error skipping request")
			continue
		}
		// TODO: forward request
		// TODO: response to request



	}

}



