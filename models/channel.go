package models

import "github.com/jinzhu/gorm"

// Channel ...
type Channel struct {
	gorm.Model
	Name    string
	Members []User
}
