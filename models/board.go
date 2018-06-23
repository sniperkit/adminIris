package models

import "github.com/jinzhu/gorm"

// Board ...
type Board struct {
	gorm.Model
	Title   string
	Content string
	writer  string
	viewcnt uint
}
