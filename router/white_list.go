package router

var whiteList = map[string]bool{
	"/health":                   true,
	"/check":                    true,
	"/admin/captcha/slide":      true,
	"/user/captcha/slide":       true,
	"user/captcha/slide/verify": true,
}
