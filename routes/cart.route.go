package routes

import (
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/handlers"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/middlewares"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/repositories"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CartRoute(router fiber.Router, db *gorm.DB) {
	cartRepository := repositories.NewCartRepositoy(db)
	productRepository := repositories.NewProductRepositoy(db)

	cartUsecase := usecases.NewCartUsecase(cartRepository)
	productUsecase := usecases.NewProductUsecase(productRepository)

	handler := handlers.NewCartHandler(cartUsecase, productUsecase)

	router.Get("/", handler.GetAll)
	router.Delete("/delete/all", handler.DeleteAll)

	// Authorization
	router.Use(middlewares.RestrictedUser)

	router.Post("/", handler.Create)
	router.Get("/user", handler.GetAllByUser)
	router.Get("/:id", handler.GetByID)
	router.Patch("/:id", handler.UpdateByID)
	router.Delete("/:id", handler.DeleteByID)
	//
}
