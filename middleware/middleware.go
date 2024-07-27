package middleware

import (
	"github.com/gin-gonic/gin"
)
func LogRequest(ctx *gin.Context) {
	//request := *ctx.Request
//	logrus.Infof("%s \t%s \t%s", request.Method,request.URL.Path ,request.URL.Host)
	ctx.Next()
}