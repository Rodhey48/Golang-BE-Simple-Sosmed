package controllers

import (
	"errors"
	"net/http"
	"simple_sosmed/configs"
	"simple_sosmed/helper"
	"simple_sosmed/middlewares"
	"simple_sosmed/models/base"
	"simple_sosmed/models/users/dto"
	"simple_sosmed/models/users/entity"
	"simple_sosmed/models/users/response"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterController(c echo.Context) error {
	var userRegister dto.UserRegister
	c.Bind(&userRegister)

	var userEntity entity.User
	userEntity.MapFromRegister(userRegister)

	userEntity.Email = strings.Trim(userEntity.Email, " ")
	userEntity.Email = strings.ToLower(userEntity.Email)

	validate := userEntity.ValidateData()

	if !validate.Status {
		return c.JSON(http.StatusBadRequest, validate)
	}

	var foundUser entity.User

	found := configs.DB.Find(&foundUser, "email = ?", userEntity.Email)

	if found.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, base.BaseRespose{
			Status:  false,
			Message: "Email already use",
			Data:    nil,
		})
	}

	result := configs.DB.Create(&userEntity)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseRespose{
			Status:  false,
			Message: "Internal Server Error",
			Data:    result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseRespose{
		Status:  true,
		Message: "Success register",
		Data:    userRegister,
	})
}

func LoginController(c echo.Context) error {
	var userLogin dto.UserLogin
	c.Bind(&userLogin)

	var userEntity entity.User

	result := configs.DB.
		Where("email = ?",
			userLogin.Email).First(&userEntity)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "Email user not found",
			Data:    nil,
		})
	}

	if !helper.CheckPasswordHash(userLogin.Password, userEntity.Password) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "Email or Password user is wrong",
			Data:    nil,
		})
	}

	var userResponse response.UserResponse
	userResponse.MapFromDatabase(userEntity)

	return c.JSON(http.StatusOK, base.BaseRespose{
		Status:  true,
		Message: "Success login",
		Data:    userResponse,
	})

}

func GetUsersLoggedController(c echo.Context) error {
	user := middlewares.ClaimsToken(c)
	var userEntity entity.User

	result := configs.DB.
		Where("id = ?",
			user.Id).Select("id, name, email").First(&userEntity)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "Email user not found",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, base.BaseRespose{
		Status:  true,
		Message: "Success get logged user info",
		Data:    userEntity,
	})
}
