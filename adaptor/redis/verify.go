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

func (v *Verify) GetCaptchaTicket(ctx context.Context, key string) (string, error) {
	result, err := v.redis.Get(ctx, constants.CaptchaTicketKey+key).Result()
	if err != nil {
		v.redis.Del(ctx, constants.CaptchaTicketKey+key)
		return "", err
	}
	return result, nil
}
