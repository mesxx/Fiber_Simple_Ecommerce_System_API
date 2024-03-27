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
	OrderHandler interface {
		Create(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		GetAllByUser(c *fiber.Ctx) error
		GetByID(c *fiber.Ctx) error
		UpdateByID(c *fiber.Ctx) error
		DeleteByID(c *fiber.Ctx) error
		DeleteAll(c *fiber.Ctx) error
	}

	orderHandler struct {
		OrderUsecase   usecases.OrderUsecase
		ProductUsecase usecases.ProductUsecase
	}
)

func NewOrderHandler(orderUsecase usecases.OrderUsecase, productUsecase usecases.ProductUsecase) OrderHandler {
	return &orderHandler{
		OrderUsecase:   orderUsecase,
		ProductUsecase: productUsecase,
	}
}

func (handler orderHandler) Create(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	// START request
	var requestCreateOrder models.RequestCreateOrder
	if err := c.BodyParser(&requestCreateOrder); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	requestCreateOrder.UserID = userSigned.ID

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestCreateOrder); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	//START calculate product orders
	totalPrice := 0
	for i, v := range requestCreateOrder.ProductOrders {
		// START get product
		product, err := handler.ProductUsecase.GetByID(v.ProductID)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END get product

		if v.Qty > product.Qty {
			return fiber.NewError(fiber.StatusBadRequest, "not enough product")
		}

		requestCreateOrder.ProductOrders[i].TotalPrice = product.Price * v.Qty
		totalPrice += int(requestCreateOrder.ProductOrders[i].TotalPrice)
	}
	//END calculate product orders

	// START midtrans
	orderData := helpers.GenerateSnapReq(userSigned.Name, "", userSigned.Email, totalPrice)
	paymentUrl, err1 := helpers.CreateSnapTransactionURL(orderData)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}
	// END midtrans

	requestCreateOrder.Status = "pending"
	requestCreateOrder.PaymentID = orderData.TransactionDetails.OrderID
	requestCreateOrder.TotalPrice = uint(totalPrice)

	createdData, err2 := handler.OrderUsecase.Create(&requestCreateOrder)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	res := fiber.Map{
		"order":       createdData,
		"payment_url": paymentUrl,
	}
	return c.Status(fiber.StatusCreated).JSON(helpers.GetResponseData(fiber.StatusCreated, "success", res))
}

func (handler orderHandler) GetAll(c *fiber.Ctx) error {
	orders, err := handler.OrderUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", orders))
}

func (handler orderHandler) GetAllByUser(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	orders, err := handler.OrderUsecase.GetAllByUser(userSigned.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", orders))
}

func (handler orderHandler) GetByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	order, err2 := handler.OrderUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", order))
}

func (handler orderHandler) UpdateByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	order, err2 := handler.OrderUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	if order.Status == "pending" {
		status, err := helpers.GetCoreAPITransactionData(order.PaymentID)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		order.Status = status.TransactionStatus
	}

	updateByID, err3 := handler.OrderUsecase.UpdateByID(order)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler orderHandler) DeleteByID(c *fiber.Ctx) error {
	userSigned := c.Locals("user").(*models.CustomClaims)

	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	order, err2 := handler.OrderUsecase.GetByID(uint(value), userSigned.ID)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	deleteByID, err3 := handler.OrderUsecase.DeleteByID(order)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", deleteByID))
}

func (handler orderHandler) DeleteAll(c *fiber.Ctx) error {
	if err := handler.OrderUsecase.DeleteAll(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponse(fiber.StatusOK, "success"))
}
