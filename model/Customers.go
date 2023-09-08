package model

import (
	"gorm.io/gorm"
)

type Customers struct {
	Name      string
	Email     string
	Role      string
	Phone     int
	Contacted bool

	gorm.Model
}
