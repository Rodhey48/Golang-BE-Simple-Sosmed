package routes

import (
	"os"
	"simple_sosmed/controllers"
	"simple_sosmed/middlewares"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.POST("/login", controllers.LoginController)
	e.POST("/register", controllers.RegisterController)

	eAuth := e.Group("")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middlewares.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("PRIVATE_KEY_JWT")),
	}
	eAuth.Use(echojwt.WithConfig(config))
	eAuth.GET("/me", controllers.GetUsersLoggedController)
	eAuth.GET("/posts", controllers.GetPost)
	eAuth.POST("/posts", controllers.CreatePostingController)
	eAuth.PUT("/posts/:id", controllers.EditPostUserController)
	eAuth.DELETE("posts/:id", controllers.DeletePostController)
}
