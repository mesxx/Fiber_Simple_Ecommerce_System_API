package handlers

import (
	"strconv"

	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/helpers"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/usecases"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	CartHandler interface {
		Create(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		GetAllByUser(c *fiber.Ctx) error
		GetByID(c *fiber.Ctx) error
		UpdateByID(c *fiber.Ctx) error
		DeleteByID(c *fiber.Ctx) error
		DeleteAll(c *fiber.Ctx) error
	}

	cartHandler struct {
		CartUsecase    usecases.CartUsecase
		ProductUsecase usecases.ProductUsecase
	}
)

func NewCartHandler(cartUsecase usecases.CartUsecase, productUsecase usecases.ProductUsecase) CartHandler {
	return &cartHandler{
		CartUsecase:    cartUsecase,
		ProductUsecase: productUsecase,
	}
}

func (handler cartHandler) Create(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	// START request
	var requestCreateCart models.RequestCreateCart
	if err := c.BodyParser(&requestCreateCart); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START check cart
	if _, err := handler.CartUsecase.GetByProductID(requestCreateCart.ProductID, userSigned.ID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END check cart

	// START get product
	product, err1 := handler.ProductUsecase.GetByID(requestCreateCart.ProductID)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}
	// END get product

	if requestCreateCart.Qty > product.Qty {
		return fiber.NewError(fiber.StatusBadRequest, "not enough product")
	}

	requestCreateCart.UserID = userSigned.ID
	requestCreateCart.TotalPrice = product.Price * requestCreateCart.Qty

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestCreateCart); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	res, err2 := handler.CartUsecase.Create(&requestCreateCart)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(helpers.GetResponseData(fiber.StatusCreated, "success", res))
}

func (handler cartHandler) GetAll(c *fiber.Ctx) error {
	carts, err := handler.CartUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", carts))
}

func (handler cartHandler) GetAllByUser(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	carts, err := handler.CartUsecase.GetAllByUser(userSigned.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", carts))
}

func (handler cartHandler) GetByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	cart, err2 := handler.CartUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", cart))
}

func (handler cartHandler) UpdateByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	cart, err2 := handler.CartUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START get product
	product, err3 := handler.ProductUsecase.GetByID(cart.ProductID)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	// END get product

	// START request
	var requestUpdateCartByID models.RequestUpdateCart
	if err := c.BodyParser(&requestUpdateCartByID); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	if requestUpdateCartByID.Qty != nil {
		if *requestUpdateCartByID.Qty > product.Qty {
			return fiber.NewError(fiber.StatusBadRequest, "not enough product")
		}

		cart.Qty = *requestUpdateCartByID.Qty
		cart.TotalPrice = product.Price * cart.Qty
	}

	updateByID, err4 := handler.CartUsecase.UpdateByID(cart)
	if err4 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err4.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler cartHandler) DeleteByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	cart, err2 := handler.CartUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	deleteByID, err3 := handler.CartUsecase.DeleteByID(cart)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", deleteByID))
}

func (handler cartHandler) DeleteAll(c *fiber.Ctx) error {
	if err := handler.CartUsecase.DeleteAll(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponse(fiber.StatusOK, "success"))
}
