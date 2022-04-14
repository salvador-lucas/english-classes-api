package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Enabled  bool   `json:"active"`
}
