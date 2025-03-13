package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qujing226/quantum_post/config"
	"github.com/qujing226/quantum_post/internal/event"
)

type App struct {
	server    *gin.Engine
	consumers []event.Consumer
}

func main() {
	config.InitViper()

}
