package models

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	StudentID uint    `gorm:"column:id_student;not null"`
	Amount    float32 `gorm:"column:amount;not null"`
	Period    string  `gorm:"column:period;not null"`
}
