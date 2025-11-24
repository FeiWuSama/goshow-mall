package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func NewResultWithOk(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Result{
		Code: OK.Code,
		Data: data,
		Msg:  OK.Msg,
	})
}

func NewResultWithError(ctx *gin.Context, data any, err *BusinessError) {
	ctx.JSON(http.StatusOK, Result{
		Code: err.Code,
		Data: data,
		Msg:  err.Msg,
	})
}
