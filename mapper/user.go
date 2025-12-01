package mapper

import (
	"context"
	"workspace-goshow-mall/adaptor/repo/model"
)

type UserMapper interface {
	GetUserByMobile(ctx context.Context) (*model.User, error)
}
