package api

import (
	"context"
	"github.com/goccy/go-json"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/constants"
)

type BaseCtrl struct {
}

func (c *BaseCtrl) GetUserVo(ctx context.Context, adaptor *adaptor.Adaptor, token string) (*vo.UserVo, error) {
	re, err := adaptor.Redis.Get(ctx, constants.UserTokenKey+token).Result()
	if err != nil {
		return nil, err
	}
	var userVo *vo.UserVo
	if err := json.Unmarshal([]byte(re), &userVo); err != nil {
		return nil, err
	}
	return userVo, nil
}

func (c *BaseCtrl) GetAdminVo(ctx context.Context, adaptor *adaptor.Adaptor, token string) (*vo.UserVo, error) {
	re, err := adaptor.Redis.Get(ctx, constants.AdminTokenKey+token).Result()
	if err != nil {
		return nil, err
	}
	var userVo *vo.UserVo
	if err := json.Unmarshal([]byte(re), &userVo); err != nil {
		return nil, err
	}
	return userVo, nil
}
