package entity

import (
	"simple_sosmed/models/base"
	"simple_sosmed/models/posts/dto"
	"simple_sosmed/models/users/entity"
)

type Posts struct {
	base.BaseEntity
	Message string      `json:"message" gorm:"not null"`
	PicUrl  string      `json:"picUrl"`
	UserId  string      `json:"userId"`
	User    entity.User `gorm:"foreignKey:UserId;references:Id"`
}

func (posts *Posts) MapFromDto(postsDto dto.PostDto, user entity.User) {
	posts.Message = postsDto.Message
	posts.PicUrl = postsDto.PicUrl
	posts.UserId = user.Id
}
