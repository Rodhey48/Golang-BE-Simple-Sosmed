package configs

import (
	"fmt"
	"os"
	entityPost "simple_sosmed/models/posts/entity"
	entityUser "simple_sosmed/models/users/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}
	Migration()
}

func Migration() {
	err := DB.AutoMigrate(entityUser.User{}, entityPost.Posts{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	fmt.Println("? Connected Successfully to the Database")

	// Execute the SQL statement to create the "uuid-ossp" extension if it doesn't exist
	result := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if result.Error != nil {
		panic(result.Error)
	}
}
