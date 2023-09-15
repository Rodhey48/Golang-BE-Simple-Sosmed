package response

import (
	"simple_sosmed/middlewares"
	"simple_sosmed/models/users/entity"
)

type UserResponse struct {
	Id    int    `json:"id" gorm:"primaryKey autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (userResponse *UserResponse) MapFromDatabase(userDatabase entity.User) {
	userResponse.Id = userDatabase.Id
	userResponse.Name = userDatabase.Name
	userResponse.Email = userDatabase.Email
	userResponse.Token = middlewares.GenerateJWT(userDatabase)
}
