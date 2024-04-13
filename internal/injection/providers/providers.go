package providers

import (
	"github.com/Tiger-Coders/tigerlily-payment/internal/config"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/logger"
	repo "github.com/Tiger-Coders/tigerlily-payment/internal/repository/checkout"
)

func ApplicationConfigProvider() *config.ApplicationConfig {
	return config.LoadConfig()
}

func LoggerProvider() *logger.Logger {
	return config.LoadConfig().DefaultLogger
}

func GeneralConfigProvider() *config.GeneralConfig {
	return &config.LoadConfig().GeneralConfig
}

// func PaymentDBInstanceProvider() *sql.DB {
// 	return config.LoadConfig().PaymentDB
// }

func PaymentDBInstanceProviderInterface() repo.CheckoutDBInterface {
	return config.LoadConfig().PaymentDB
}

func StripeConfigProvider() *config.StripeService {
	return &config.LoadConfig().GeneralConfig.StripeService
}
