package router

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/utils/logger"
)

type responseWriterWrapper struct {
	gin.ResponseWriter
	writer io.Writer
}

func (w responseWriterWrapper) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

func GetResponseBody(c *gin.Context) string {
	if c.Request.Response == nil {
		return ""
	}
	data, _ := io.ReadAll(c.Request.Response.Body)
	return string(data)
}

func AccessLogMiddleware(filter func(ctx *gin.Context) bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		if filter != nil && !filter(context) {
			context.Next()
			return
		}
		//记录日志
		begin := time.Now()
		fields := []zap.Field{
			zap.String("ip", context.RemoteIP()),
			zap.String("method", context.Request.Method),
			zap.String("path", context.Request.URL.Path),
			zap.String("token", context.GetHeader(constants.UserToken)),
		}
		var responseBody bytes.Buffer
		multiWriter := io.MultiWriter(context.Writer, &responseBody)
		context.Writer = &responseWriterWrapper{
			ResponseWriter: context.Writer,
			writer:         multiWriter,
		}
		resp := responseBody.String()
		if len(resp) > 1024 {
			resp = resp[:1024]
		}
		context.Next()
		fields = append(fields, zap.Int64("cost", time.Since(begin).Microseconds()))
		fields = append(fields, zap.Int("status", context.Writer.Status()))
		fields = append(fields, zap.String("response", resp))
		logger.Info("记录日志", fields...)
	}
}
