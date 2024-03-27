package repositories

import (
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"

	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Create(product *models.Product) (*models.Product, error)
		GetAll() ([]models.Product, error)
		GetByID(id uint) (*models.Product, error)
		UpdateByID(product *models.Product, values map[string]interface{}) (*models.Product, error)
		DeleteByID(product *models.Product) (*models.Product, error)
		DeleteAll() error
	}

	productRepository struct {
		DB *gorm.DB
	}
)

func NewProductRepositoy(db *gorm.DB) ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (repository productRepository) Create(product *models.Product) (*models.Product, error) {
	if err := repository.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repository productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := repository.DB.Preload("ProductOrders").Preload("Carts").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repository productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := repository.DB.Preload("ProductOrders").Preload("Carts").Where("ID = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository productRepository) UpdateByID(product *models.Product, values map[string]interface{}) (*models.Product, error) {
	if err := repository.DB.Table("products").Preload("ProductOrders").Preload("Carts").Where("ID = ?", product.ID).Updates(values).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repository productRepository) DeleteByID(product *models.Product) (*models.Product, error) {
	if err := repository.DB.Preload("ProductOrders").Preload("Carts").Where("ID = ?", product.ID).Delete(product, product.ID).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repository productRepository) DeleteAll() error {
	if err := repository.DB.Where("1 = 1").Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}
