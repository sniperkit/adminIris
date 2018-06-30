package models

import "github.com/jinzhu/gorm"

// Board ...
type Board struct {
	gorm.Model
	Title   string `form:"title"`
	Content string `form:"content"`
	Writer  string `form:"writer"`
	Viewcnt uint   `form:"viewcnt"`
}
