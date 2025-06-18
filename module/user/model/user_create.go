package userModel

import "GoCare/common"

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password_hash;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
	Role            string `json:"role" gorm:"column:role;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}
