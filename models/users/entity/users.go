package entity

import (
	"simple_sosmed/models/users/dto"
)

type User struct {
	Id       int    `json:"id" gorm:"primaryKey autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
