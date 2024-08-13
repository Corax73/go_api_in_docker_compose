package customDb

import (
	"prettyApi/models"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedingProducts creates test entries in the products table.
func SeedingProducts(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		id := uuid.New()
		product := models.Product{ID: id, Title: "Test_product" + strconv.Itoa(i), CreatedAt: time.Now(), Price: 123 + i, QuantityInStock: 10 + i}
		db.Create(&product)
	}
}
