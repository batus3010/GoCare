package patientModel

import (
	"GoCare/common"
	"strings"
)

type PatientUpdate struct {
	FirstName *string `json:"first-name" gorm:"column:first_name"`
	LastName  *string `json:"last-name" gorm:"column:last_name"`
	Address   *string `json:"address" gorm:"column:address"`
	Status    *int    `json:"status" gorm:"column:status"`
}

func (PatientUpdate) TableName() string {
	return Patient{}.TableName()
}

func (u *PatientUpdate) Validate() error {
	if strPtr := u.FirstName; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return common.ErrFirstNameIsBlank
		}
		u.FirstName = &str
	}

	if strPtr := u.LastName; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return common.ErrLastNameIsBlank
		}
		u.LastName = &str
	}

	if strPtr := u.Address; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return common.ErrAddressIsBlank
		}
		u.Address = &str
	}
	return nil

}
