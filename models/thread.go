package models

import "github.com/jinzhu/gorm"

// Thread ...
type Thread struct {
	gorm.Model
	Topic string
	Name  string

	Threads []Thread
}
