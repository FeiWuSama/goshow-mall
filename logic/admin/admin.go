package admin

import (
	"context"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/mapper"
)

type Service struct {
	adapter     *adaptor.Adaptor
	adminMapper mapper.AdminMapper
}

func (s *Service) HelloWorld(c context.Context) string {
	return s.adminMapper.HelloWorld(c)
}
