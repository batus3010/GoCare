package patientModel

import (
	"GoCare/common"
	"strings"
)

type PatientCreate struct {
	common.SQLModel
	FirstName string `json:"first-name" gorm:"column:first_name"`
	LastName  string `json:"last-name" gorm:"column:last_name"`
	//Gender    string `json:"gender" gorm:"column:gender"`
	//Dob       string `json:"dob" gorm:"column:date_of_birth"`
	//Phone     string `json:"phone" gorm:"column:phone"`
}

func (PatientCreate) TableName() string {
	return Patient{}.TableName()
}

func (data PatientCreate) Validate() error {
	data.FirstName = strings.TrimSpace(data.FirstName)
	if data.FirstName == "" {
		return common.ErrFirstNameIsBlank
	}
	data.LastName = strings.TrimSpace(data.LastName)
	if data.LastName == "" {
		return common.ErrLastNameIsBlank
	}
	return nil
}
