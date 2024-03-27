package routes

import (
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/handlers"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/repositories"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewProductRepositoy(db)
	usecase := usecases.NewProductUsecase(repository)
	handler := handlers.NewProductHandler(usecase)

	router.Static("/image", "./publics/images")

	router.Post("/", handler.Create)
	router.Get("/", handler.GetAll)
	router.Get("/:id", handler.GetByID)
	router.Get("/image/:id", handler.GetImageByID)
	router.Patch("/:id", handler.UpdateByID)
	router.Delete("/image/:id", handler.DeleteImageByID)
	router.Delete("/:id", handler.DeleteByID)
	router.Delete("/delete/all", handler.DeleteAll)

}
