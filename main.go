package main

import (
	"os"
	"simple_sosmed/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	// configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8000"
	}
	return ":" + port
}

// func loadEnv(){
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("Failed load env file")
// 	}
// }
