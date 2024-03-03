package models

import "gorm.io/gorm"

// Book holds information about different websites
type Book struct {
	gorm.Model
	Title string
	Price string
}
