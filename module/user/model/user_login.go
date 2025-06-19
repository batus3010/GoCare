package userModel

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password_hash" gorm:"column:password_hash;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}
