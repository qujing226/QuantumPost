//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/qujing226/quantum_post/ioc"
	"github.com/qujing226/quantum_post/pkg/log/zapLogger"
)

var thirdProvider = wire.NewSet(

	zapLogger.NewZapLogger,
)

func InitWebServer() *App {
	wire.Build(
		thirdProvider,

		ioc.InitWebServer,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}

func NewGinEngine() *gin.Engine {
	return gin.Default()
}
