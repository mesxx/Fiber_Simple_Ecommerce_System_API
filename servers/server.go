package servers

import (
	"fmt"

	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/configs"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/middlewares"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"
	"github.com/mesxx/Fiber_Simple_Ecommerce_System_API/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func Server() *fiber.App {
	err1 := godotenv.Load()
	if err1 != nil {
		fmt.Println("error godotenv:", err1.Error())
	}

	db, err2 := configs.DatabaseConfig()
	if err2 != nil {
		fmt.Println("error gorm database config:", err2.Error())
	}
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.ProductOrder{}, &models.Cart{})

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorMiddleware,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	user := api.Group("/user")
	product := api.Group("/product")
	cart := api.Group("/cart")
	order := api.Group("/order")

	routes.UserRoute(user, db)
	routes.ProductRoute(product, db)
	routes.CartRoute(cart, db)
	routes.OrderRoute(order, db)

	return app
}
