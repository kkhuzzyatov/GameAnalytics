package entities

import "gorm.io/gorm"

type Developer struct {
	Login        string    `json:"login"`
	HashPassword int       `json:"hash_of_password"`
	Projects     []Project `json:"project"`
	gorm.Model
}
