package usermodel

type UserFilter struct {
	UserName string `form:"userName" json:"userName"`
}
