package main

import (
	// "log"
	// "welcome-to-costco/pkg/storage"

	"fmt"
	"log"
	"os"
	Types "welcome-to-costco/migrations/internal/storage/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_SSL_MODE"),
	)

	log.Println("!!!!!!!!poop!!!!!!!")
	log.Println(dsn)
	log.Println("!!!!!!!!!!!!!!!!!!!")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Types.Customer{})

	return events.APIGatewayProxyResponse{
		Body:       "migration",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
