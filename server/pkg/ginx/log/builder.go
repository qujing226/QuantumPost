package log

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type AccessLog struct {
	Method   string
	URL      string
	Duration string
	ReqBody  string
	RespBody string
	Status   int
}

type MiddlewareBuild struct {
	allowReqBody  bool
	allowRespBody bool
	loggerFunc    func(ctx context.Context, al *AccessLog)
}

func NewBuilder(loggerFunc func(ctx context.Context, al *AccessLog)) *MiddlewareBuild {
	return &MiddlewareBuild{
		loggerFunc: loggerFunc,
	}
}

func (b *MiddlewareBuild) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		al := &AccessLog{
			Method: ctx.Request.Method,
			URL:    ctx.Request.URL.String(),
		}
		if b.allowReqBody && ctx.Request.Body != nil {
			body, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
			// 这是一个很消耗CPU和内存的操作
			if len(body) > 1024 {
				al.ReqBody = string(body[:1024])
			} else {
				al.ReqBody = string(body)
			}
		}
		if b.allowRespBody {
			ctx.Writer = responseWriter{
				al:             al,
				ResponseWriter: ctx.Writer,
			}
		}
		defer func() {
			al.Duration = time.Since(start).String()
			b.loggerFunc(ctx, al)
		}()
		ctx.Next()
	}
}

type responseWriter struct {
	al *AccessLog
	gin.ResponseWriter
}

func (w responseWriter) WriteHeader(statusCode int) {
	w.al.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w responseWriter) Writer(data []byte) (int, error) {
	if len(data) > 1024 {
		w.al.RespBody = string(data[:1024])
	} else {
		w.al.RespBody = string(data)
	}
	return w.ResponseWriter.Write(data)
}

func (w responseWriter) WriteString(data string) (int, error) {
	if len(data) > 1024 {
		w.al.RespBody = data[:1024]
	} else {
		w.al.RespBody = data
	}
	return w.ResponseWriter.WriteString(data)
}

func (b *MiddlewareBuild) AllowReqBody() *MiddlewareBuild {
	b.allowReqBody = true
	return b
}

func (b *MiddlewareBuild) AllowRespBody() *MiddlewareBuild {
	b.allowRespBody = true
	return b
}
