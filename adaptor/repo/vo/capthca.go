package vo

type SlideCaptchaVo struct {
	Key              string `json:"key"`
	ImageBase64      string `json:"ImageBase64"`
	TitleImageBase64 string `json:"TitleImageBase64"`
	TitleHeight      int    `json:"TitleHeight"`
	TitleWidth       int    `json:"TitleWidth"`
	TitleX           int    `json:"TitleX"`
	TitleY           int    `json:"TitleY"`
}

type SlideCaptchaCheckVo struct {
	Ticket string `json:"ticket"`
	Expire int64  `json:"expire"`
}
