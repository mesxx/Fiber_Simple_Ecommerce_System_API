package usecases

import (
	"errors"

	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/repositories"
)

type (
	OrderUsecase interface {
		Create(requestCreateOrder *models.RequestCreateOrder) (*models.Order, error)
		GetAll() ([]models.Order, error)
		GetAllByUser(userID uint) ([]models.Order, error)
		GetByID(id uint, userID uint) (*models.Order, error)
		UpdateByID(order *models.Order) (*models.Order, error)
		DeleteByID(order *models.Order) (*models.Order, error)
		DeleteAll() error
	}

	orderUsecase struct {
		OrderRepository repositories.OrderRepository
	}
)

func NewOrderUsecase(repository repositories.OrderRepository) OrderUsecase {
	return &orderUsecase{
		OrderRepository: repository,
	}
}

func (usecase orderUsecase) Create(requestCreateOrder *models.RequestCreateOrder) (*models.Order, error) {
	order := models.Order{
		UserID:        requestCreateOrder.UserID,
		Status:        requestCreateOrder.Status,
		PaymentID:     requestCreateOrder.PaymentID,
		TotalPrice:    requestCreateOrder.TotalPrice,
		ProductOrders: requestCreateOrder.ProductOrders,
	}
	return usecase.OrderRepository.Create(&order)
}

func (usecase orderUsecase) GetAll() ([]models.Order, error) {
	return usecase.OrderRepository.GetAll()
}

func (usecase orderUsecase) GetAllByUser(userID uint) ([]models.Order, error) {
	return usecase.OrderRepository.GetAllByUser(userID)
}

func (usecase orderUsecase) GetByID(id uint, userID uint) (*models.Order, error) {
	getByID, err := usecase.OrderRepository.GetByID(id, userID)
	if err != nil {
		return nil, err
	} else if getByID.ID == 0 {
		return nil, errors.New("order ID is invalid, please try again")
	}
	return getByID, nil
}

func (usecase orderUsecase) UpdateByID(order *models.Order) (*models.Order, error) {
	return usecase.OrderRepository.UpdateByID(order)
}

func (usecase orderUsecase) DeleteByID(order *models.Order) (*models.Order, error) {
	return usecase.OrderRepository.DeleteByID(order)
}

func (usecase orderUsecase) DeleteAll() error {
	return usecase.OrderRepository.DeleteAll()
}
