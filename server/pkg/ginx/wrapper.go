package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qujing226/quantum_post/pkg/log"
	"net/http"
)

func WrapBody[T any, C any](l log.Logger, fn func(ctx context.Context, req T, uc C) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req T
		if err := ctx.Bind(&req); err != nil {
			return
		}
		val, ok := ctx.Get("user")
		if !ok {
			ctx.JSON(http.StatusOK, Result{
				Code: 4,
				Msg:  "user not login",
			})
		}
		uc, ok := val.(C)
		if !ok {
			ctx.JSON(http.StatusOK, Result{
				Code: 4,
				Msg:  "user not login",
			})
		}

		res, err := fn(ctx, req, uc)
		if err != nil {
			l.Error("Server Handler ERROR",
				log.String("Path: ", ctx.Request.URL.Path),
				log.String("Route:", ctx.FullPath()),
				log.Error(err))
			ctx.JSON(http.StatusOK, Result{
				Code: 5,
				Msg:  "server error",
			})
		}
		ctx.JSON(http.StatusOK, res)
	}
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}
