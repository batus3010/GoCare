package patientModel

import "GoCare/common"

const EntityName = "Patient"

type Patient struct {
	common.SQLModel
	FirstName string `json:"first-name" gorm:"column:first_name"`
	LastName  string `json:"last-name" gorm:"column:last_name"`
	Gender    string `json:"gender" gorm:"column:gender"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Email     string `json:"email" gorm:"column:email"`
	Address   string `json:"address" gorm:"column:address"`
}

func (Patient) TableName() string { return "patients" }
