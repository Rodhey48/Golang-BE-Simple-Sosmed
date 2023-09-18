package controllers

import (
	"errors"
	"net/http"
	"simple_sosmed/configs"
	"simple_sosmed/middlewares"
	"simple_sosmed/models/base"
	"simple_sosmed/models/posts/dto"
	postsEntity "simple_sosmed/models/posts/entity"
	usersEntity "simple_sosmed/models/users/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateTags		Posts
// @Summary			Get user post
// @Description		Get user post
// @Produce			application/json
// @Tags			Posts
// @Success			200 {object} base.BaseRespose{}
// @Router			/posts [get]
// @Security 		BearerAuth
func GetPost(c echo.Context) error {
	user := middlewares.ClaimsToken(c)
	var userEntity usersEntity.User

	result := configs.DB.
		Where("id = ?",
			user.Id).Select("id, name, email").First(&userEntity)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	var posts []postsEntity.Posts

	foundPost := configs.DB.Where("posts.user_id = ? AND posts.is_active = true", userEntity.Id).Preload("User").Find(&posts)

	if errors.Is(foundPost.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "Post not found",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, base.BaseRespose{
		Status:  true,
		Message: "Ok",
		Data:    posts,
	})
}

// CreateTags		Posts
// @Summary			Create user post
// @Description		Create user post
// @Produce			application/json
// @Tags			Posts
// @Success			201 {object} base.BaseRespose{}
// @Router			/posts [POST]
// @Security 		BearerAuth
func CreatePostingController(c echo.Context) error {
	var postUser dto.PostUserDto
	c.Bind(&postUser)

	user := middlewares.ClaimsToken(c)
	var userEntity usersEntity.User

	result := configs.DB.
		Where("id = ?",
			user.Id).Select("id, name, email").First(&userEntity)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	var postsEntity postsEntity.Posts

	postsEntity.MapFromDto(postUser, userEntity)

	saved := configs.DB.Create(&postsEntity)

	if saved.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseRespose{
			Status:  false,
			Message: "Internal Server Error",
			Data:    result.Error.Error(),
		})
	}

	return c.JSON(http.StatusCreated, base.BaseRespose{
		Status:  true,
		Message: "Success posting",
		Data:    postUser,
	})
}

// CreateTags		Posts
// @Summary			Update user post
// @Description		Update user post
// @Produce			application/json
// @Param			idPost path string true "update post by id"
// @Tags			Posts
// @Success			201 {object} base.BaseRespose{}
// @Router			/posts/{idPost} [PUT]
// @Security 		BearerAuth
func EditPostUserController(c echo.Context) error {
	idPost := c.Param("id")

	var postUser dto.PostUserDto
	c.Bind(&postUser)

	user := middlewares.ClaimsToken(c)
	var userEntity usersEntity.User

	result := configs.DB.
		Where("id = ?",
			user.Id).Select("id, name, email").First(&userEntity)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	var postsEntity postsEntity.Posts

	foundPost := configs.DB.Where("posts.id = ? AND posts.is_active = true", idPost).Preload("User").First(&postsEntity)

	if errors.Is(foundPost.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "Post not found",
			Data:    nil,
		})
	}

	if userEntity.Id != postsEntity.UserId {
		return c.JSON(http.StatusForbidden, base.BaseRespose{
			Status:  false,
			Message: "You can't edit posts that don't belong to you",
			Data:    nil,
		})
	}

	postsEntity.Message = postUser.Message
	postsEntity.PicUrl = postUser.PicUrl

	err := configs.DB.Save(&postsEntity)

	if err.Error != nil {
		return c.JSON(http.StatusForbidden, base.BaseRespose{
			Status:  false,
			Message: "Internal Server Error",
			Data:    err.Error.Error(),
		})
	}

	return c.JSON(http.StatusCreated, base.BaseRespose{
		Status:  true,
		Message: "Success edit posting",
		Data:    postsEntity,
	})

}

// CreateTags		Posts
// @Summary			Delete user post
// @Description		Delete user post
// @Produce			application/json
// @Param			idPost path string true "delete posts by id"
// @Tags			Posts
// @Success			200 {object} base.BaseRespose{}
// @Router			/posts/{idPost} [DELETE]
// @Security 		BearerAuth
func DeletePostController(c echo.Context) error {
	idPost := c.Param("id")

	user := middlewares.ClaimsToken(c)
	var userEntity usersEntity.User

	result := configs.DB.
		Where("id = ?",
			user.Id).Select("id, name, email").First(&userEntity)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	}

	var postsEntity postsEntity.Posts

	foundPost := configs.DB.Where("posts.id = ? AND posts.is_active = true", idPost).Preload("User").First(&postsEntity)

	if errors.Is(foundPost.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseRespose{
			Status:  false,
			Message: "Post not found",
			Data:    nil,
		})
	}

	if userEntity.Id != postsEntity.UserId {
		return c.JSON(http.StatusForbidden, base.BaseRespose{
			Status:  false,
			Message: "You can't delete posts that don't belong to you",
			Data:    nil,
		})
	}

	postsEntity.IsActive = false

	err := configs.DB.Save(&postsEntity)

	if err.Error != nil {
		return c.JSON(http.StatusForbidden, base.BaseRespose{
			Status:  false,
			Message: "Internal Server Error",
			Data:    err.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.BaseRespose{
		Status:  true,
		Message: "Success delete posting",
		Data:    nil,
	})
}
