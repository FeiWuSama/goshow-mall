package dto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type SlideCaptchaDto struct {
	Once string `url:"once"`
	Time int64  `url:"ts"`
	Sign string `url:"sign"`
}

func (c *SlideCaptchaDto) CheckSign() bool {
	data := fmt.Sprintf("%s%s%d", c.Once, "fewiwusama2015", c.Time)
	hash := sha256.Sum256([]byte(data))
	return c.Sign == hex.EncodeToString(hash[:])
}

type SlideCaptchaCheckDto struct {
	Key    string `json:"key"`
	SlideX int    `json:"slideX"`
	SlideY int    `json:"slideY"`
}
