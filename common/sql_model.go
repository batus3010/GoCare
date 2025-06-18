package common

import "time"

type SQLModel struct {
	Id        int       `json:"-" gorm:"column:id;"`
	Status    int       `json:"status" gorm:"column:status;default:1;"` // add default to have default value
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
