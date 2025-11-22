package router

import "github.com/gin-gonic/gin"

func AuthMiddleware(filter func(ctx *gin.Context) bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		if filter != nil && !filter(context) {
			context.Next()
			return
		}
		// 鉴权中间件
		context.Next()
	}
}
