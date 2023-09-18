package main

import (
	"os"
	"simple_sosmed/configs"
	_ "simple_sosmed/docs"
	"simple_sosmed/routes"

	"github.com/labstack/echo/v4"
)

// @title Simple Sosmed API
// @version 1.0
// @description This is a sample server Simple_Sosmed server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Simple Suppor
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":3000"
	}
	return ":" + port
}

// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("Failed load env file")
// 	}
// }
