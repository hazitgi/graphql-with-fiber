package config

import (
	"fmt"
	"log"
	"os"

	"github.com/hazitgi/graphql-with-fiber/models"
	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	Host := os.Getenv("DB_HOST")
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	Port := os.Getenv("DB_PORT")

	// Provide default values if environment variables are not set
	if Host == "" {
		Host = "database" // This should match the service name in docker-compose
	}
	if User == "" {
		User = "postgres"
	}
	if Password == "" {
		Password = "root"
	}
	if DbName == "" {
		DbName = "loomERP" // Make sure this matches your actual database name
	}
	if Port == "" {
		Port = "5432" // Default PostgreSQL port inside the container
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", Host, User, Password, DbName, Port)

	fmt.Println("Connecting to database...", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected successfully")

	if err := DB.AutoMigrate(
		&models.User{}); err != nil {
		log.Fatalln("failed to migrate tables")
	}
}

func GetDb() *gorm.DB {
	if DB == nil {
		log.Fatalln("Db connection isn't initialized!")
	}
	return DB
}
