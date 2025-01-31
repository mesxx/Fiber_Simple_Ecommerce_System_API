package helpers

import "github.com/mesxx/Fiber_Simple_Ecommerce_System_API/models"

func GetResponse(code int, message string) *models.Response {
	return &models.Response{
		Code:    code,
		Message: message,
	}
}

func GetResponseData(code int, message string, data interface{}) *models.ResponseData {
	return &models.ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
