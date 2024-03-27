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
	ProductHandler interface {
		Create(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		GetByID(c *fiber.Ctx) error
		GetImageByID(c *fiber.Ctx) error
		UpdateByID(c *fiber.Ctx) error
		DeleteImageByID(c *fiber.Ctx) error
		DeleteByID(c *fiber.Ctx) error
		DeleteAll(c *fiber.Ctx) error
	}

	productHandler struct {
		ProductUsecase usecases.ProductUsecase
	}
)

func NewProductHandler(usecase usecases.ProductUsecase) ProductHandler {
	return &productHandler{
		ProductUsecase: usecase,
	}
}

func (handler productHandler) Create(c *fiber.Ctx) error {
	// START request
	var requestCreateProduct models.RequestCreateProduct
	if err := c.BodyParser(&requestCreateProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START request file
	fileImage, err1 := c.FormFile("image")
	if err1 == nil {
		// START check type
		fileType := fileImage.Header.Get("Content-Type")
		if err := helpers.UploadSettingType(fileType); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END check type

		// START set filename
		fileName, err := helpers.UploadSettingName(fileImage.Filename)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END set filename

		// START save file
		destination := "./publics/images/" + fileName
		if err := c.SaveFile(fileImage, destination); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END save file

		requestCreateProduct.Image = fileName
	}
	// END request file

	// START validator
	validate := validator.New()
	if err := validate.Struct(requestCreateProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END validator

	res, err2 := handler.ProductUsecase.Create(&requestCreateProduct)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(helpers.GetResponseData(fiber.StatusCreated, "success", res))
}

func (handler productHandler) GetAll(c *fiber.Ctx) error {
	products, err := handler.ProductUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", products))
}

func (handler productHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	product, err2 := handler.ProductUsecase.GetByID(uint(value))
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", product))
}

func (handler productHandler) GetImageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	product, err2 := handler.ProductUsecase.GetByID(uint(value))
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START set path
	fileName := product.Image.String
	destination := "./publics/images/" + fileName
	// END set path

	return c.Status(fiber.StatusOK).SendFile(destination)
}

func (handler productHandler) UpdateByID(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	product, err2 := handler.ProductUsecase.GetByID(uint(value))
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START request
	var requestUpdateByIDProduct models.RequestUpdateProduct
	if err := c.BodyParser(&requestUpdateByIDProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END request

	// START map request
	var values = make(map[string]interface{})
	// END map request

	if requestUpdateByIDProduct.Title != "" {
		values["title"] = requestUpdateByIDProduct.Title
	}

	if requestUpdateByIDProduct.Qty != nil {
		values["qty"] = requestUpdateByIDProduct.Qty
	}

	if requestUpdateByIDProduct.Price != nil {
		values["price"] = requestUpdateByIDProduct.Price
	}

	if requestUpdateByIDProduct.Description != nil {
		if *requestUpdateByIDProduct.Description == "" {
			values["description"] = nil
		} else {
			values["description"] = requestUpdateByIDProduct.Description
		}
	}

	// START request file
	fileImage, err3 := c.FormFile("image")
	if err3 == nil {
		// START check type
		fileType := fileImage.Header.Get("Content-Type")
		if err := helpers.UploadSettingType(fileType); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END check type

		// START set filename
		fileName, err := helpers.UploadSettingName(fileImage.Filename)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END set filename

		// START delete old file
		if err := helpers.DeleteImage(product.Image.String); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END delete old file

		// START save new file
		destination := "./publics/images/" + fileName
		if err := c.SaveFile(fileImage, destination); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// END save new file

		values["image"] = fileName
	}
	// END request file

	updateByID, err4 := handler.ProductUsecase.UpdateByID(product, values)
	if err4 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err4.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler productHandler) DeleteImageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	product, err2 := handler.ProductUsecase.GetByID(uint(value))
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START delete file
	if err := helpers.DeleteImage(product.Image.String); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END delete file

	// START map request
	var values = make(map[string]interface{})
	values["image"] = nil
	// END map request

	updateByID, err3 := handler.ProductUsecase.UpdateByID(product, values)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", updateByID))
}

func (handler productHandler) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err1 := strconv.Atoi(id)
	if err1 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err1.Error())
	}

	product, err2 := handler.ProductUsecase.GetByID(uint(value))
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// START delete file
	if err := helpers.DeleteImage(product.Image.String); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END delete file

	deleteByID, err3 := handler.ProductUsecase.DeleteByID(product)
	if err3 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err3.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponseData(fiber.StatusOK, "success", deleteByID))
}

func (handler productHandler) DeleteAll(c *fiber.Ctx) error {
	// START delete all file
	if err := helpers.DeleteAllImage(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// END delete all file

	if err := handler.ProductUsecase.DeleteAll(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(helpers.GetResponse(fiber.StatusOK, "success"))
}
