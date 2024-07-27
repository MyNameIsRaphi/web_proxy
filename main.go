package main

import (
	"fmt"
	"os"

	"sync"

	"github.com/MyNameIsRaphi/web_proxy/forward"
	"github.com/MyNameIsRaphi/web_proxy/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
const PORT = 8080
const HTTPS_PORT = 4040



func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go runHTTPProxy(&wg)	
	
	wg.Wait()	

}
func runHTTPSProxy(wg *sync.WaitGroup){
	key_path, exists := os.LookupEnv("KEY_PATH")
	if !exists{
		logrus.Fatal("Failed to read KEY_PATH")
	} 
	cert_path, exists := os.LookupEnv("CERT_PATH")
	if !exists {
		logrus.Fatal("Failed to read CERT_PATH")
	}
	//TODO add TLS proxy
	server := gin.Default()
	var addr string = fmt.Sprintf(":%d", HTTPS_PORT)
	server.Use(forward.HandleRequest)
	err := server.RunTLS(addr, cert_path, key_path)
	if err != nil {
		logrus.Fatal(err)
	}
	wg.Done()	
}

func runHTTPProxy(wg *sync.WaitGroup) {
	var err error
	server := gin.Default()
	server.Use(middleware.LogRequest, forward.HandleRequest)
	logrus.Infof("Starting server on port %d", PORT)
	err = server.Run()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
	wg.Done()
}
