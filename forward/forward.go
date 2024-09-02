package forward

import (
	"io"
	"net"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleRequest(ctx *gin.Context) {
	req := *ctx.Request 

	var method string = req.Method
	if method == "CONNECT" {
		tunnelRequest(ctx)
	}
}

func tunnelRequest(ctx *gin.Context) {
	destConn, err := net.Dial("tcp", ctx.Request.Host)
	logrus.Info(ctx.Request.Host)
	
	if err != nil {
		logrus.Error(err)
		ctx.AbortWithError(http.StatusInternalServerError,err)
		return 
	}
	defer destConn.Close()
	ctx.Status(http.StatusOK)
	hijacker, ok := ctx.Writer.(http.Hijacker) // check if writer supports the hijacking interface
	if !ok {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Hijacking not supported"))
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		logrus.Error(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer clientConn.Close()
	// TODO use net/http library
	go func() {
		_, err := io.Copy(destConn, clientConn)
		if err != nil {
			logrus.Info(err)
		}
	}()
	_, err = io.Copy(clientConn, destConn)		
	if err != nil {
		logrus.Info(err)
	}
}

