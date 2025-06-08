package entities

import "gorm.io/gorm"

type Player struct {
	Metrics []Metric `json:"metrics"`
	gorm.Model
}
