package adaptor

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"workspace-goshow-mall/config"
)

type Adaptor struct {
	config config.Config
	Db     *gorm.DB
	Redis  *redis.Client
}

// NewAdaptor 自动注入适配器
func NewAdaptor(conf config.Config, db *gorm.DB, redis *redis.Client) *Adaptor {
	return &Adaptor{
		config: conf,
		Db:     db,
		Redis:  redis,
	}
}
