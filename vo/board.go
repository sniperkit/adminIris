package vo

import (
	"github.com/jinzhu/gorm"
	"github.com/senseoki/adminIris/models"
)

// Board ...
type Board struct {
	RDB *gorm.DB
	models.Board
}
