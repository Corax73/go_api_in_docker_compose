package customDb

import (
	"prettyApi/customLog"
	"prettyApi/models"
)

// Init conducts initial migrations and populates test data. Returns true on success.
func Init() bool {
	var resp bool
	database := GetConnect()
	if database != nil {
		errProduct := database.AutoMigrate(&models.Product{})
		if errProduct == nil {
			var count int64
			database.Model(&models.Product{}).Count(&count)
			if count == 0 {
				SeedingProducts(database)
			}
			if count > 0 {
				resp = true
			}
		} else {
			customLog.Logging(errProduct)
		}
	}
	return resp
}
