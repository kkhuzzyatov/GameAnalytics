package entities

import "gorm.io/gorm"

type Project struct {
	Name    string   `json:"name"`
	Players []Player `json:"players"`
	gorm.Model
}
