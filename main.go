package main

import (
	"os"
	"simple_sosmed/configs"
	"simple_sosmed/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}
