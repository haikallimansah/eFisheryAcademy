package routes

import (
	"clean_arch/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/user", userHandler.CreateUser)
}
