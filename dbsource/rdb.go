package dbsource

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/senseoki/adminIris/models"
)

// Storage ...
type Storage struct {
	db *gorm.DB
}

// NewRDB ...
func (s *Storage) NewRDB(url string, pool int) {
	log.Println("2222")
	db, err := gorm.Open("mysql", url)
	s.db = db
	if err != nil {
		panic("failed to connect database")
	}
	defer s.db.Close()

	s.db.DB().SetMaxIdleConns(10)

	s.db.DB().SetMaxOpenConns(100)

	pingErr := s.db.DB().Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("db connect ok")
	s.Migrate()
}

func (s *Storage) Migrate() {
	s.db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Post{}, &models.PostComment{}, &models.Channel{})
}
