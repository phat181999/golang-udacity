package seed

import (
	"github/phat/product/model"

	"gorm.io/gorm"
)

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
