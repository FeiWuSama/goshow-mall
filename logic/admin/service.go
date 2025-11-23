package admin

import (
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/dao"
)

func NewService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		adapter:     adaptor,
		adminMapper: dao.NewAdminDao(*adaptor),
	}
}
