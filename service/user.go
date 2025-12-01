package service

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/vo"
)

type IUserService interface {
	SMobileLogin(context context.Context, dto interface{}) (*vo.UserVo, error)
}
