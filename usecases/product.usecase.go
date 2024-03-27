package usecases

import (
	"database/sql"
	"errors"

	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/repositories"
)

type (
	ProductUsecase interface {
		Create(requestCreateProduct *models.RequestCreateProduct) (*models.Product, error)
		GetAll() ([]models.Product, error)
		GetByID(id uint) (*models.Product, error)
		UpdateByID(product *models.Product, values map[string]interface{}) (*models.Product, error)
		DeleteByID(product *models.Product) (*models.Product, error)
		DeleteAll() error
	}

	productUsecase struct {
		ProductRepository repositories.ProductRepository
	}
)

func NewProductUsecase(repository repositories.ProductRepository) ProductUsecase {
	return &productUsecase{
		ProductRepository: repository,
	}
}

func (usecase productUsecase) Create(requestCreateProduct *models.RequestCreateProduct) (*models.Product, error) {
	product := models.Product{
		Title:       requestCreateProduct.Title,
		Qty:         requestCreateProduct.Qty,
		Price:       requestCreateProduct.Price,
		Description: sql.NullString{String: requestCreateProduct.Description, Valid: requestCreateProduct.Description != ""},
		Image:       sql.NullString{String: requestCreateProduct.Image, Valid: requestCreateProduct.Image != ""},
	}
	return usecase.ProductRepository.Create(&product)
}

func (usecase productUsecase) GetAll() ([]models.Product, error) {
	return usecase.ProductRepository.GetAll()
}

func (usecase productUsecase) GetByID(id uint) (*models.Product, error) {
	getByID, err := usecase.ProductRepository.GetByID(id)
	if err != nil {
		return nil, err
	} else if getByID.ID == 0 {
		return nil, errors.New("product is invalid, please try again")
	}
	return getByID, nil
}

func (usecase productUsecase) UpdateByID(product *models.Product, values map[string]interface{}) (*models.Product, error) {
	return usecase.ProductRepository.UpdateByID(product, values)
}

func (usecase productUsecase) DeleteByID(product *models.Product) (*models.Product, error) {
	return usecase.ProductRepository.DeleteByID(product)
}

func (usecase productUsecase) DeleteAll() error {
	return usecase.ProductRepository.DeleteAll()
}
