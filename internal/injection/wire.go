//go:build wireinject
// +build wireinject

package injection

import (
	// "database/sql"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/injection/providers"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	repo "github.com/ZAF07/tigerlily-e-bakery-payment/internal/repository/checkout"
	"github.com/google/wire"
	// "github.com/jinzhu/gorm"
)

func GetGeneralConfig() *config.GeneralConfig {
	wire.Build(providers.GeneralConfigProvider)
	return &config.GeneralConfig{}
}

func GetApplicationConfig() *config.ApplicationConfig {
	wire.Build(providers.ApplicationConfigProvider)
	return &config.ApplicationConfig{}
}

func GetLogger() *logger.Logger {
	wire.Build(providers.LoggerProvider)
	return config.LoadConfig().DefaultLogger
}

// func GetPaymentDBInstance() *sql.DB {
// 	wire.Build(providers.PaymentDBInstanceProvider)
// 	return config.LoadConfig().PaymentDB
// }

func GetPaymentDBInstance() repo.CheckoutDBInterface {
	wire.Build(providers.PaymentDBInstanceProviderInterface)
	return config.LoadConfig().PaymentDB
}

func GetStripeServiceConfig() *config.StripeService {
	wire.Build(providers.StripeConfigProvider)
	return &config.LoadConfig().GeneralConfig.StripeService
}
