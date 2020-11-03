package database

import (
	"log"
	"os"

	"github.com/hooneun/go-postgres-auth/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB a Database instance
var DB *gorm.DB

// ConnectToDB Connects the database
func ConnectToDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file\n", err)
	}

	// dns := fmt.Sprintf("host=127.0.0.1 user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
	// os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))
	dns := "host=127.0.0.1 user=postgres password= dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Seoul"

	log.Print("Connectiong to PostgreSQL DB...")
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)

	}
	log.Println("connected")
	// turned on the loger on info mode
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Print("Running the migrations...")
	DB.AutoMigrate(&models.User{}, &models.Claims{})
}
