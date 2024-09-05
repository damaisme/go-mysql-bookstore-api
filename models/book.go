package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// Books model
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
