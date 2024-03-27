package repositories

import (
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"

	"gorm.io/gorm"
)

type (
	CartRepository interface {
		Create(cart *models.Cart) (*models.Cart, error)
		GetAll() ([]models.Cart, error)
		GetAllByUser(userID uint) ([]models.Cart, error)
		GetByID(id uint, userID uint) (*models.Cart, error)
		GetByProductID(productID uint, userID uint) (*models.Cart, error)
		UpdateByID(cart *models.Cart) (*models.Cart, error)
		DeleteByID(cart *models.Cart) (*models.Cart, error)
		DeleteAll() error
	}

	cartRepository struct {
		DB *gorm.DB
	}
)

func NewCartRepositoy(db *gorm.DB) CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (repository cartRepository) Create(cart *models.Cart) (*models.Cart, error) {
	if err := repository.DB.Create(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (repository cartRepository) GetAll() ([]models.Cart, error) {
	var carts []models.Cart
	if err := repository.DB.Preload("User").Preload("Product").Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (repository cartRepository) GetAllByUser(userID uint) ([]models.Cart, error) {
	var carts []models.Cart
	if err := repository.DB.Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (repository cartRepository) GetByID(id uint, userID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := repository.DB.Preload("User").Preload("Product").Where("ID = ?", id).Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (repository cartRepository) GetByProductID(productID uint, userID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := repository.DB.Preload("User").Preload("Product").Where("product_id = ?", productID).Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (repository cartRepository) UpdateByID(cart *models.Cart) (*models.Cart, error) {
	if err := repository.DB.Preload("User").Preload("Product").Where("ID = ?", cart.ID).Updates(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (repository cartRepository) DeleteByID(cart *models.Cart) (*models.Cart, error) {
	if err := repository.DB.Preload("User").Preload("Product").Where("ID = ?", cart.ID).Delete(cart, cart.ID).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (repository cartRepository) DeleteAll() error {
	if err := repository.DB.Where("1 = 1").Delete(&models.Cart{}).Error; err != nil {
		return err
	}
	return nil
}
