package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"workspace-goshow-mall/adaptor"
)

type AdminDao struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewAdminDao(adaptor adaptor.Adaptor) *AdminDao {
	return &AdminDao{
		db:          adaptor.Db,
		redisClient: adaptor.Redis,
	}
}

func (a *AdminDao) HelloWorld(ctx context.Context) string {
	return "Hello World"
}
