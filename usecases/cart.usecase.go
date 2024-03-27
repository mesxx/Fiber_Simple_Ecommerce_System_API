package usecases

import (
	"errors"

	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/repositories"
)

type (
	CartUsecase interface {
		Create(requestCreateCart *models.RequestCreateCart) (*models.Cart, error)
		GetAll() ([]models.Cart, error)
		GetAllByUser(userID uint) ([]models.Cart, error)
		GetByID(id uint, userID uint) (*models.Cart, error)
		GetByProductID(productID uint, userID uint) (*models.Cart, error)
		UpdateByID(cart *models.Cart) (*models.Cart, error)
		DeleteByID(cart *models.Cart) (*models.Cart, error)
		DeleteAll() error
	}

	cartUsecase struct {
		CartRepository repositories.CartRepository
	}
)

func NewCartUsecase(repository repositories.CartRepository) CartUsecase {
	return &cartUsecase{
		CartRepository: repository,
	}
}

func (usecase cartUsecase) Create(requestCreateCart *models.RequestCreateCart) (*models.Cart, error) {
	cart := models.Cart{
		UserID:     requestCreateCart.UserID,
		ProductID:  requestCreateCart.ProductID,
		Qty:        requestCreateCart.Qty,
		TotalPrice: requestCreateCart.TotalPrice,
	}
	return usecase.CartRepository.Create(&cart)
}

func (usecase cartUsecase) GetAll() ([]models.Cart, error) {
	return usecase.CartRepository.GetAll()
}

func (usecase cartUsecase) GetAllByUser(userID uint) ([]models.Cart, error) {
	return usecase.CartRepository.GetAllByUser(userID)
}

func (usecase cartUsecase) GetByID(id uint, userID uint) (*models.Cart, error) {
	getByID, err := usecase.CartRepository.GetByID(id, userID)
	if err != nil {
		return nil, err
	} else if getByID.ID == 0 {
		return nil, errors.New("cart ID is invalid, please try again")
	}
	return getByID, nil
}

func (usecase cartUsecase) GetByProductID(productID uint, userID uint) (*models.Cart, error) {
	getByProductID, err := usecase.CartRepository.GetByProductID(productID, userID)
	if err != nil {
		return nil, err
	} else if getByProductID.ID > 0 {
		return nil, errors.New("already added to cart")
	}
	return getByProductID, nil
}

func (usecase cartUsecase) UpdateByID(cart *models.Cart) (*models.Cart, error) {
	return usecase.CartRepository.UpdateByID(cart)
}

func (usecase cartUsecase) DeleteByID(cart *models.Cart) (*models.Cart, error) {
	return usecase.CartRepository.DeleteByID(cart)
}

func (usecase cartUsecase) DeleteAll() error {
	return usecase.CartRepository.DeleteAll()
}
