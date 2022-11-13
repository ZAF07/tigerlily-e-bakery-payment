package providers

import (
	"database/sql"

	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/config"
)

func GeneralConfigProvider() *config.GeneralConfig {
	return &config.LoadConfig().GeneralConfig
}

func PaymentDBInstanceProvider() *sql.DB {
	return config.LoadConfig().PaymentDB
}

func StripeConfigProvider() *config.StripeService {
	return &config.LoadConfig().GeneralConfig.StripeService
}
