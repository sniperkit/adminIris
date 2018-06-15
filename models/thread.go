package models

import "github.com/jinzhu/gorm"

// Thread ...
type Thread struct {
	gorm.Model
	Name  string
	Posts []Post
}
