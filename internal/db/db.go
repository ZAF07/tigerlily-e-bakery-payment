package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Tiger-Coders/tigerlily-payment/internal/injection"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/env"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/logger"
	repos "github.com/Tiger-Coders/tigerlily-payment/internal/repository/checkout"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ORM *sql.DB

type Db struct {
	db *gorm.DB
}

func NewDB() *sql.DB {
	connectDB()
	return ORM
}

func connectDB() {

	logs := logger.NewLogger()
	db, err := sql.Open("postgres", env.GetDBEnv())
	if err != nil {
		logs.ErrorLogger.Printf("Couldn't connect to Database %+v", err)
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	logs.InfoLogger.Println("Successfully connected to Database")

	ORM = db
}

// LATEST IMPLEMENTATION
// func InitPostgresDB() *sql.DB {
func InitPostgresDB() repos.CheckoutDBInterface {
	logger := injection.GetApplicationConfig().DefaultLogger
	config := injection.GetGeneralConfig().PaymentDB
	sourceName := config.GetPostgresDBString()
	fmt.Println("DATABASE NAME : ---> ", sourceName)

	db, err := sql.Open("postgres", sourceName)
	if err != nil {
		logger.ErrorLogger.Printf("error connection to database : %+v\n", err)
	}

	// Calling the DB() function on the *gorm.DB instance returns the underlying *sql.DB instance
	db.SetMaxOpenConns(config.MaxConn)
	db.SetConnMaxIdleTime(time.Duration(config.MaxIdleConn))
	d := repos.NewCheckoutRepo(db)

	if pingErr := d.Ping(); pingErr != nil {
		logger.ErrorLogger.Printf("🚨database ping error : %+v\n", pingErr)
	}
	// logger.InfoLogger.Println("database connected ! 🎇")
	logger.InfoLogger.Println("database connected")
	return d
}
