package vo

type UserVo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
