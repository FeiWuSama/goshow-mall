package admin

import (
	"github.com/gin-gonic/gin"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/api"
	"workspace-goshow-mall/constants"
	"workspace-goshow-mall/logic/admin"
	"workspace-goshow-mall/result"
	"workspace-goshow-mall/service"
)

type Ctrl struct {
	api.BaseCtrl
	adaptor      *adaptor.Adaptor
	adminService service.IAdminService
}

func NewCtrl(adaptor *adaptor.Adaptor) *Ctrl {
	return &Ctrl{
		adaptor:      adaptor,
		adminService: admin.NewService(adaptor),
	}
}

func (c *Ctrl) CreateUser(ctx *gin.Context) {
	token := ctx.Request.Header.Get(constants.AdminToken)
	adminVo := c.GetAdminVo(ctx, *c.adaptor, token)
	if adminVo == nil {
		result.NewResultWithError(ctx, nil, result.NewBusinessError(result.Unauthorized))
		ctx.Abort()
		return
	}
	//helloWorld := c.adminService.CreateUser(ctx.Request.Context())
	result.NewResultWithOk(ctx, adminVo)
}
