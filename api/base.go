package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/constants"
)

type BaseCtrl struct {
}

func (c *BaseCtrl) GetUserVo(ctx *gin.Context, adaptor adaptor.Adaptor, token string) *vo.UserVo {
	result := adaptor.Redis.Get(ctx, constants.UserTokenKey+token)
	if result == nil {
		return nil
	}
	val := result.Val()
	var userVo *vo.UserVo
	if err := json.Unmarshal([]byte(val), &userVo); err != nil {
		return nil
	}
	return userVo
}

func (c *BaseCtrl) GetAdminVo(ctx *gin.Context, adaptor adaptor.Adaptor, token string) *vo.UserVo {
	result := adaptor.Redis.Get(ctx, constants.AdminTokenKey+token)
	if result == nil {
		return nil
	}
	val := result.Val()
	var userVo *vo.UserVo
	if err := json.Unmarshal([]byte(val), &userVo); err != nil {
		return nil
	}
	return userVo
}
