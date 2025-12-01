package dto

type UserMobileLoginDto struct {
	Mobile string `json:"mobile"`
	Ticket string `json:"ticket"`
}

type UserMobilePasswordLoginDto struct {
	UserMobileLoginDto
	Password string `json:"password"`
}
