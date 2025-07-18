package db

import (
	"blogApp/models"
	"fmt"
	"log"
	"os"
	
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDb() {
	var err error
	err = godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	//connStr := "user= password= dbname= port= sslmode=disable"
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println("Connection string:", connStr)

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//migrate the schema
	err = DB.AutoMigrate(models.User{}, &models.Post{})
	if err != nil {
		log.Fatal("failed to migrate schema", err)
	}
	fmt.Println("connected to database successfully")
}