package db

import (
	"database/sql"
	"log"

	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/injection"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/env"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
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

	logs := logger.NewLogger()
	db, err := gorm.Open("postgres", env.GetDBEnv())
	if err != nil {
		logs.ErrorLogger.Printf("Couldn't connect to Database %+v", err)
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	logs.InfoLogger.Println("Successfully connected to Database")

	ORM = db
}

// LATEST IMPLEMENTATION
func InitPostgresDB() {
	config := injection.GetGeneralConfig().PaymentDB
	sourceName := config.GetPostgresDBString()

	db, err := sql.Open("postgres", sourceName)
	if err != nil {
		log.Fatalf("error connecting to database : %+v\n", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("ping error database : %+v\n", err)
	}
	db.SetMaxOpenConns(config.MaxConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
}
