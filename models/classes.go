package models

import "github.com/jinzhu/gorm"

type Classes struct {
	gorm.Model
	StudentID uint   `json:"studenId" gorm:"column:id_student;not null"`
	Day       string `json:"day" gorm:"column:day;not null"`
	Enabled   bool   `json:"enabled" gorm:"column:enabled;not null;default:true"`
	Hour      string `json:"hour" gorm:"column:hour;"`
}
