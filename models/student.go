package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name     string `gorm:"column:name;not null`
	Lastname string `gorm:"column:lastname;not null`
	Enabled  bool   `gorm:"column:enabled;default true`
}
