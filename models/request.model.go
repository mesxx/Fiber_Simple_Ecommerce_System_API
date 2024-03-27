package models

type (
	RequestCreateUser struct {
		Name     string `json:"name" validate:"required,min=5,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestLoginUser struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestCreateProduct struct {
		Title       string `json:"title" validate:"required"`
		Qty         uint   `json:"qty" validate:"required"`
		Price       uint   `json:"price" validate:"required"`
		Description string `json:"description"`
		Image       string `json:"image"`
	}

	RequestUpdateProduct struct {
		Title       string  `json:"title"`
		Qty         *uint   `json:"qty"`
		Price       *uint   `json:"price"`
		Description *string `json:"description"`
	}

	RequestCreateCart struct {
		UserID     uint `json:"user_id" validate:"required"`
		ProductID  uint `json:"product_id" form:"product_id" validate:"required"`
		Qty        uint `json:"qty" validate:"required"`
		TotalPrice uint `json:"total_price" validate:"required"`
	}

	RequestUpdateCart struct {
		Qty *uint `json:"qty"`
	}

	RequestCreateOrder struct {
		UserID        uint           `json:"user_id" validate:"required"`
		Status        string         `json:"status"`
		PaymentID     string         `json:"payment_id"`
		TotalPrice    uint           `json:"total_price"`
		ProductOrders []ProductOrder `json:"product_orders" validate:"required"`
	}
)
