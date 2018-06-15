package models

import "github.com/jinzhu/gorm"

// PostComment ...
type PostComment struct {
	gorm.Model
	Content string
}
