package dto

type AddAdminDto struct {
	Name     string `json:"username"`
	NickName string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Sex      int32  `json:"sex"`
}

type UpdateAdminDto struct {
	Id       int32  `json:"id"`
	Name     string `json:"username"`
	NickName string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Sex      int32  `json:"sex"`
}
