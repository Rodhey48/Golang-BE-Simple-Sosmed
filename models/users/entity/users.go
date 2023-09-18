package entity

import (
	helper "simple_sosmed/helper"
	"simple_sosmed/models/base"
	"simple_sosmed/models/users/dto"

	"gorm.io/gorm"
)

type User struct {
	base.BaseEntity
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
}

func (user *User) MapFromLogin(userLogin dto.UserLogin) {
	user.Email = userLogin.Email
	user.Password = userLogin.Password
}

func (user *User) MapFromRegister(userRegister dto.UserRegister) {
	user.Name = userRegister.Name
	user.Email = userRegister.Email
	user.Password = userRegister.Password
}

func (user *User) ValidateData() base.BaseRespose {
	if user.Email == "" {
		return base.BaseRespose{
			Status:  false,
			Message: "Email cant empty",
			Data:    nil,
		}
	}

	if !helper.IsValidEmail(user.Email) {
		return base.BaseRespose{
			Status:  false,
			Message: "Invalid email format",
			Data:    nil,
		}
	}

	if user.Name == "" {
		return base.BaseRespose{
			Status:  false,
			Message: "Name cant empty",
			Data:    nil,
		}
	}

	if len(user.Name) < 3 {
		return base.BaseRespose{
			Status:  false,
			Message: "Name length must be min 3 character",
			Data:    nil,
		}
	}

	if user.Password == "" {
		return base.BaseRespose{
			Status:  false,
			Message: "Password cant empty",
			Data:    nil,
		}
	}

	if len(user.Password) < 6 {
		return base.BaseRespose{
			Status:  false,
			Message: "Password length must be min 6 character",
			Data:    nil,
		}
	}

	return base.BaseRespose{
		Status:  true,
		Message: "Data Oke",
	}
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	passworHash, _ := helper.HashPassword(user.Password)
	user.Password = passworHash
	return
}
