package constants

const (
	UserToken  = "token"
	AdminToken = "token"
)

const (
	UserTokenKey     = "goshow:user:token:"
	AdminTokenKey    = "goshow:admin:token:"
	SlideCaptchaKey  = "goshow:slide:captcha:"
	CaptchaTicketKey = "goshow:ticket:"

	TokenExpire         = 60 * 60 * 24 * 3
	CaptchaExpire       = 60 * 5
	CaptchaTicketExpire = 60
)

const (
	UserActiveStatus = 1
	UserBanStatus    = -1
)
