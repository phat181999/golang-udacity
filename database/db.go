package database

import (
	"github/phat/product/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	DB *gorm.DB
}
var DB *gorm.DB
func (s *Server) Init() *gorm.DB {	
	var err error
	var db *gorm.DB
	dns := "host=localhost user=postgres password=tanphat99 dbname=customers port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	s.DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("Connect Database Success")
	}

	s.DB.Migrator().DropTable(&model.Customers{})
	s.DB.AutoMigrate(&model.Customers{})
	db = s.DB
	return db
}

func  GetDb() *gorm.DB {
	return DB
}