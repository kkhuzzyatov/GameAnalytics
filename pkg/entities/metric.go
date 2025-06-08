package entities

import "gorm.io/gorm"

type Metric struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	gorm.Model
}
