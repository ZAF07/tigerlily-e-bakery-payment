package app

import (
	"github.com/Tiger-Coders/tigerlily-payment/internal/config"
	"github.com/Tiger-Coders/tigerlily-payment/internal/db"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/constants"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/logger"
)

func InitApplication() {
	appConfig := config.LoadConfig()
	initLogger(appConfig)
	initDB(appConfig)
}

func initDB(appConfig *config.ApplicationConfig) {

	switch appConfig.GeneralConfig.DBType {
	case constants.POSTGRES:
		appConfig.PaymentDB = db.InitPostgresDB()
	default:
		db.NewDB()
	}
}

func initLogger(appConfig *config.ApplicationConfig) {
	loggerType := appConfig.GeneralConfig.Logger

	switch loggerType {
	case constants.Default:
		appConfig.DefaultLogger = loadDefaultLogger()
	}
	return
}

func loadDefaultLogger() *logger.Logger {
	return logger.NewLogger()
}
