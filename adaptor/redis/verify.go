package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"workspace-goshow-mall/constants"
)

type Verify struct {
	redis *redis.Client
}

func NewVerify(redis *redis.Client) *Verify {
	return &Verify{
		redis: redis,
	}
}

func (v *Verify) SaveCaptcha(ctx context.Context, key string, value string) error {
	return v.redis.Set(ctx, constants.SlideCaptchaKey+key, value, constants.CaptchaExpire*time.Second).Err()
}

func (v *Verify) GetCaptcha(ctx context.Context, key string) (string, error) {
	result, err := v.redis.Get(ctx, constants.SlideCaptchaKey+key).Result()
	if err != nil {
		v.redis.Del(ctx, constants.SlideCaptchaKey+key)
		return "", err
	}
	return result, nil
}

func (v *Verify) SaveCaptchaTicket(ctx context.Context, key string, value string) error {
	return v.redis.Set(ctx, constants.CaptchaTicketKey+key, value, constants.CaptchaExpire*time.Second).Err()
}

func (v *Verify) GetCaptchaTicket(ctx context.Context, key string) (int64, error) {
	result, err := v.redis.Exists(ctx, constants.CaptchaTicketKey+key).Result()
	if err != nil {
		v.redis.Del(ctx, constants.CaptchaTicketKey+key)
		return result, err
	}
	return result, nil
}

func (v *Verify) SaveUserToken(ctx context.Context, key string, value string) error {
	return v.redis.Set(ctx, constants.UserTokenKey+key, value, constants.TokenExpire*time.Second).Err()
}

func (v *Verify) GetUserToken(ctx context.Context, key string) (string, error) {
	result, err := v.redis.Get(ctx, constants.UserTokenKey+key).Result()
	if err != nil {
		v.redis.Del(ctx, constants.UserTokenKey+key)
		return "", err
	}
	return result, nil
}

func (v *Verify) IncrPasswordErrorCount(ctx context.Context, key string) (int64, error) {
	pipeline := v.redis.Pipeline()
	result, err := pipeline.Incr(ctx, constants.PasswordErrorKey+key).Result()
	if err != nil {
		return 0, err
	}
	if result == 1 {
		pipeline.Expire(ctx, constants.PasswordErrorKey+key, constants.PasswordErrorExpire*time.Minute)
	}
	_, err = pipeline.Exec(ctx)
	return result, err
}

func (v *Verify) DeletePasswordErrorCount(ctx context.Context, key string) error {
	return v.redis.Del(ctx, constants.PasswordErrorKey+key).Err()
}
