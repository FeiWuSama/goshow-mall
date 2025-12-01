package user

import (
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/redis"
	"workspace-goshow-mall/dao"
)

func NewService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		adapter:    adaptor,
		userMapper: dao.NewUserDao(*adaptor),
		verify:     redis.NewVerify(adaptor.Redis),
	}
}
