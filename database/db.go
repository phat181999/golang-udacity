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
	if err := SeedInitialData(db); err != nil {
        log.Fatal("Failed to seed initial data: ", err)
    }
	return db
}

func  GetDb() *gorm.DB {
	return DB
}

func SeedInitialData(db *gorm.DB) error {
    customers := []model.Customers{
        {Name: "John", Role: "Doe", Email: "john.doe@example.com", Phone: 123123, Contacted: true},
		{Name: "Phat", Role: "Doe", Email: "Danny@example.com", Phone: 123123, Contacted: false},
		{Name: "Danny", Role: "Doe", Email: "Danny@example.com", Phone: 123123, Contacted: true},
		{Name: "Nhi", Role: "Doe", Email: "Nhi@example.com", Phone: 123123, Contacted: true},
		{Name: "Tan", Role: "Doe", Email: "Tan@example.com", Phone: 123123, Contacted: false},
		{Name: "Phatht8", Role: "Doe", Email: "Phatht8@example.com", Phone: 123123, Contacted: true},
        // Add more customer data as needed
    }

    for _, customer := range customers {
        if err := db.Create(&customer).Error; err != nil {
            return err
        }
    }

    return nil
}
