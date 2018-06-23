package thread

import (
	"github.com/senseoki/adminIris/models"

	"github.com/jinzhu/gorm"
)

// ThreadVO ...
type ThreadVO struct {
	Thread *models.Thread
	DB     *gorm.DB
}
