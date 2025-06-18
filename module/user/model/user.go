package userModel

import (
	"GoCare/common"
	"errors"
)

const EntityUser = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	Role            string `json:"role" gorm:"column:role;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomErrorResponse(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
	ErrEmailExisted = common.NewCustomErrorResponse(
		errors.New("this email has already existed"),
		"this email has already existed",
		"ErrEmailExisted",
	)
)
