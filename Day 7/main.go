package main

import (
	"clean_arch/config"
	"clean_arch/handler"
	"clean_arch/repository"
	"clean_arch/routes"
	"clean_arch/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.Routes(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
