package admin

import (
	"github.com/gin-gonic/gin"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/result"
)

type Ctrl struct {
	adaptor *adaptor.Adaptor
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor: adaptor,
	}
}

func (c *Ctrl) HelloWorld(ctx *gin.Context) {
	result.NewResultWithOk(ctx, "Hello World")
}
