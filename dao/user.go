package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/model"
	"workspace-goshow-mall/adaptor/repo/query"
	"workspace-goshow-mall/result"
)

func NewUserDao(adaptor adaptor.Adaptor) *UserDao {
	return &UserDao{
		db:          adaptor.Db,
		redisClient: adaptor.Redis,
	}
}

type UserDao struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func (u UserDao) GetUserByMobile(ctx context.Context) (*model.User, error) {
	qs1 := query.Use(u.db).MobileUser
	mobileUser, err := qs1.WithContext(ctx).Where(qs1.MobileAes.Eq("")).First()
	if err != nil {
		return nil, err
	}
	if mobileUser == nil {
		return nil, result.NewBusinessErrorWithMsg(result.ParamError, "手机号未注册")
	}
	qs2 := query.Use(u.db).User
	user, err := qs2.WithContext(ctx).Where(qs2.ID.Eq(mobileUser.UserID)).First()
	return user, nil
}
