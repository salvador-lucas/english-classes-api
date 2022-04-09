package models

type Student struct {
	ID       uint     `json:"id" gorm:"primary_key"`
	Name     string   `json:"name"`
	Lastname string   `json:"lastname"`
	Active   bool     `json:"active"`
	Classes  []string `json:"classes"`
}
