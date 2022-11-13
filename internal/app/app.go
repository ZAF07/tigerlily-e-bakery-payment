package app

import (
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/db"
)

func InitApplication() {
	appConfig := config.LoadConfig()
	initDB(appConfig)
}

func initDB(appConfig *config.ApplicationConfig) {
	appConfig.PaymentDB = db.InitPostgresDB()
}
