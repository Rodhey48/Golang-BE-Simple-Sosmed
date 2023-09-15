package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"simple_sosmed/configs"
	"simple_sosmed/models/base"
	"simple_sosmed/models/users/dto"
	"simple_sosmed/models/users/entity"
	"simple_sosmed/models/users/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddUsersController(c echo.Context) error {
	var userRegister dto.UserRegister
	c.Bind(&userRegister)
	fmt.Println("Daftar user userRegister", userRegister)

	if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, base.BaseRespose{
			Status:  false,
			Message: "Email still empty",
			Data:    nil,
		})
	}

	var userDatabase entity.User
	userDatabase.Name = userRegister.Name
	userDatabase.Email = userRegister.Email
	userDatabase.Password = userRegister.Password

	fmt.Println("Daftar user userDatabase", userDatabase)

	// result := configs.DB.Create(&userDatabase)
	result := configs.DB.Create(&userDatabase)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseRespose{
			Status:  false,
			Message: "Failed add data users",
			Data:    nil,
		})
	}

	var userResponse response.UserResponse
	userResponse.MapFromDatabase(userDatabase)

	return c.JSON(http.StatusCreated, base.BaseRespose{
		Status:  true,
		Message: "Success add data users",
		Data:    userResponse,
	})
}

func LoginController(c echo.Context) error {
	var userLogin dto.UserLogin
	c.Bind(&userLogin)

	var userDatabase entity.User
	userDatabase.MapFromLogin(userLogin)

	result := configs.DB.
		Where("email = ? AND password = ?",
			userDatabase.Email,
			userDatabase.Password).First(&userDatabase)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusUnauthorized, base.BaseRespose{
			Status:  false,
			Message: "Failed login check email and password",
			Data:    nil,
		})
	}

	var userResponse response.UserResponse
	userResponse.MapFromDatabase(userDatabase)

	return c.JSON(http.StatusOK, base.BaseRespose{
		Status:  true,
		Message: "Success login",
		Data:    userResponse,
	})

}

func GetUsersController(c echo.Context) error {
	var users []entity.User

	result := configs.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseRespose{
			Status:  false,
			Message: "Failed get data users",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseRespose{
		Status:  true,
		Message: "Success get data users",
		Data:    users,
	})
}
