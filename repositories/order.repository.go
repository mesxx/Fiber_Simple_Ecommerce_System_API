package repositories

import (
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"

	"gorm.io/gorm"
)

type (
	OrderRepository interface {
		Create(order *models.Order) (*models.Order, error)
		GetAll() ([]models.Order, error)
		GetAllByUser(userID uint) ([]models.Order, error)
		GetByID(id uint, userID uint) (*models.Order, error)
		UpdateByID(order *models.Order) (*models.Order, error)
		DeleteByID(order *models.Order) (*models.Order, error)
		DeleteAll() error
	}

	orderRepository struct {
		DB *gorm.DB
	}
)

func NewOrderRepositoy(db *gorm.DB) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (repository orderRepository) Create(order *models.Order) (*models.Order, error) {
	if err := repository.DB.Create(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (repository orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	if err := repository.DB.Preload("User").Preload("ProductOrders").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (repository orderRepository) GetAllByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := repository.DB.Preload("User").Preload("ProductOrders").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (repository orderRepository) GetByID(id uint, userID uint) (*models.Order, error) {
	var order models.Order
	if err := repository.DB.Preload("User").Preload("ProductOrders").Where("ID = ?", id).Where("user_id = ?", userID).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repository orderRepository) UpdateByID(order *models.Order) (*models.Order, error) {
	if err := repository.DB.Preload("User").Preload("ProductOrders").Where("ID = ?", order.ID).Updates(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (repository orderRepository) DeleteByID(order *models.Order) (*models.Order, error) {
	if err := repository.DB.Preload("User").Preload("ProductOrders").Where("ID = ?", order.ID).Delete(order, order.ID).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (repository orderRepository) DeleteAll() error {
	if err := repository.DB.Where("1 = 1").Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}
