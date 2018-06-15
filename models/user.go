package models

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.Model
	Email    string
	Password string
	Username string
	Photo    string
	UserPost []Post
}
