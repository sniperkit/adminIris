package dbsource

import (
	"log"

	"github.com/senseoki/adminIris/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Storage ...
type Storage struct {
	RDB *gorm.DB
}

// NewRDB ...
func (s *Storage) NewRDB() {
	var err error
	s.RDB, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/goteamDev?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	s.RDB.LogMode(true)
	s.RDB.DB().SetMaxIdleConns(10)
	s.RDB.DB().SetMaxOpenConns(100)

	log.Println("RDB Connection Established")
	s.migrate()
}

// Migrate ...
func (s *Storage) migrate() {
	s.RDB.AutoMigrate(&models.Board{})
}
