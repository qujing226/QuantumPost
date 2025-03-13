package ioc

import "github.com/gin-gonic/gin"

func InitWebServer() *gin.Engine {
	return gin.Default()
}
