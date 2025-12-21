package model

import (
	"gorm.io/gorm"
)

type Number struct {
	gorm.Model

	Value int `json:"value" gorm:"not null"`
}
