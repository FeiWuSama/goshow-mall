package vo

type UserVo struct {
	Id       int64  `json:"id"`
	Nickname string `json:"Nickname"`
	Sex      int32  `json:"sex"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}
