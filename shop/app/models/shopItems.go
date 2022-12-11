package models

import "github.com/jinzhu/gorm"

type Products struct {
	gorm.Model
	Name      string
	Price     float32
	ImageName string
	Calories  float64
	Protein   float64
	Carbs     float64
	Fat       float64
}
