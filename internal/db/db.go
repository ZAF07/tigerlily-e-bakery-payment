package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/injection"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	repos "github.com/ZAF07/tigerlily-e-bakery-payment/internal/repository/checkout"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ORM *gorm.DB

type Db struct {
	db *gorm.DB
}

func NewDB() *gorm.DB {
	connectDB()
	return ORM
}

func connectDB() {
	d := injection.GetGeneralConfig().PaymentDB

	logs := logger.NewLogger()
	// 🚨🚨 ENV PACKAGE IS NOT USED ANY LONGER 🚨🚨
	// db, err := gorm.Open("postgres", env.GetDBEnv())
	fmt.Println("DATABSE OLD STRING ", d.GetPostgresDBString())
	db, err := gorm.Open("postgres", d.GetPostgresDBString())
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
	config := injection.GetGeneralConfig().PaymentDB
	sourceName := config.GetPostgresDBString()
	fmt.Println("DATABASE NAME : ---> ", sourceName)

	db, err := sql.Open("postgres", sourceName)
	if err != nil {
		log.Fatalf("error connecting to database : %+v\n", err)
	}

	// Calling the DB() function on the *gorm.DB instance returns the underlying *sql.DB instance
	db.SetMaxOpenConns(config.MaxConn)
	db.SetConnMaxIdleTime(time.Duration(config.MaxIdleConn))
	d := repos.NewCheckoutRepo(db)
	return d
}
