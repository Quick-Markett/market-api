package database

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"main.go/models"
)

var (
	PostgresInstance *gorm.DB
)

func setup(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Market{}, &models.Order{}, &models.OrderItem{}, &models.Review{}, &models.Category{}); err != nil {
		log.Fatalf("Erro ao migrar as tabelas: %v", err)
	}
	log.Println("Migração realizada com sucesso!")
}

func ConnectWithDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	appEnv := os.Getenv("APP_ENV")
	databaseConnection := ""

	if appEnv == "PROD" {
		databaseConnection = os.Getenv("DATABASE_URL")
		if databaseConnection == "" {
			log.Fatal("Database url is not set in the environment.")
		}
	} else {
		databaseConnection = "host=localhost user=postgres password=password dbname=database port=5432 sslmode=disable"
	}

	PostgresInstance, _ = gorm.Open(postgres.New(postgres.Config{
		DSN: databaseConnection,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	setup(PostgresInstance)

	database, err := PostgresInstance.DB()
	if err != nil {
		log.Fatalf("Failed to get raw SQLDB from GORM: %v", err)
	}

	database.SetConnMaxIdleTime(15)
	database.SetMaxIdleConns(10)
	database.SetConnMaxLifetime(time.Minute * 10)

	if err := database.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	} else {
		log.Println("Database connection established successfully.")
	}
}
