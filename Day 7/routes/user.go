package routes

import (
	"clean_arch/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/user", userHandler.CreateUser)
	e.GET("/user", userHandler.GetAllUser)
	e.GET("/user/:id", userHandler.GetUserById)
	e.PUT("/user/:id", userHandler.UpdateUser)
	e.DELETE("/user/:id", userHandler.DeleteUser)
}
