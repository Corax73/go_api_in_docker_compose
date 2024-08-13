package repository

import (
	"fmt"
	"prettyApi/customDb"
	"prettyApi/customLog"
	"prettyApi/models"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ProductsRepository struct {
	Offset, Limit int
	Order         string
}

// ProductsRepository returns a pointer to the initiated repository instance.
func NewRepository() *ProductsRepository {
	rep := ProductsRepository{
		Offset: 0,
		Limit:  20,
		Order:  "price desc",
	}
	return &rep
}

// GetList returns lists of entities with the total number, if a model exists, with a limit (there is a default value), offset.
func (rep *ProductsRepository) GetList() *[]map[string]interface{} {
	data := []map[string]interface{}{}
	database := customDb.GetConnect()
	product := new(models.Product)
	database.Model(product).Limit(rep.Limit).Offset(rep.Offset).Order(rep.Order).Find(&data)
	return &data
}

// GetOne
func (rep *ProductsRepository) GetOne(uuid string) *[]map[string]interface{} {
	database := customDb.GetConnect()
	data := []map[string]interface{}{}
	product := new(models.Product)
	database.Model(product).Where("id = ?", uuid).First(&data)
	return &data

}

// Create
func (rep *ProductsRepository) Create(data map[string]interface{}) (*models.Product, error) {
	var (
		resp            *models.Product
		err             error
		title           string
		price, quantity int
	)

	if val, ok := data["title"]; ok {
		title = fmt.Sprint(val)
	}
	if val, ok := data["price"]; ok {
		priceReq := fmt.Sprint(val)
		price, err = strconv.Atoi(priceReq)
	}
	if val, ok := data["quantity"]; ok {
		quantityReq := fmt.Sprint(val)
		quantity, err = strconv.Atoi(quantityReq)
	}

	if title != "" && price > 0 && quantity > 0 {
		database := customDb.GetConnect()
		newId := uuid.New()
		product := models.Product{ID: newId, Title: title, Price: price, QuantityInStock: quantity, CreatedAt: time.Now()}
		tx := database.Begin()
		result := tx.Create(&product)
		if result.Error == nil {
			res := tx.Commit()
			if res.Error != nil {
				tx.Rollback()
				err = res.Error
				customLog.Logging(res.Error)
			} else {
				resp = &product
			}
		} else {
			tx.Rollback()
			err = result.Error
			customLog.Logging(result.Error)
		}
	}
	return resp, err
}
