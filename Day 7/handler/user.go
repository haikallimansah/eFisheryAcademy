package handler

import (
	"clean_arch/entity"
	"clean_arch/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUseCase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUseCase}
}

func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.UserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := handler.userUsecase.CreateUser(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create user failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "User created success",
		Data:    user,
	})
}

func (handler UserHandler) GetAllUser(c echo.Context) error {
	users, err := handler.userUsecase.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get users Failed",
			Error:   err.Error(),
		})
	}

	if len(users) <= 0 {
		return c.JSON(http.StatusNotFound, entity.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Get users Failed. No users found",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Users found",
		Data:    users,
	})
}

func (handler UserHandler) GetUserById(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userUsecase.GetUserById(intId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get user Failed",
			Error:   err.Error(),
		})
	}

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, entity.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Get user Failed. No user found",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User found",
		Data:    user,
	})
}

func (handler UserHandler) UpdateUser(c echo.Context) error {
	req := entity.UpdateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error:   err.Error(),
		})
	}

	intId, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userUsecase.UpdateUser(req, intId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update user failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User updated successfully",
		Data:    user,
	})
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	intId, _ := strconv.Atoi(c.Param("id"))
	err := handler.userUsecase.DeleteUser(intId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete user failed",
			Error:   err.Error(),
		}))
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User deleted successfully",
	})
}
