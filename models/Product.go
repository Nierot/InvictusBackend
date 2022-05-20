package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `json:"Name"`
	Volume  int     `json:"Volume"`
	Alcohol float32 `json:"Alcohol"`
}

type ProductInput struct {
	Name    string  `json:"Name" binding:"required"`
	Volume  int     `json:"Volume" binding:"required"`
	Alcohol float32 `json:"Alcohol" binding:"required"`
}
