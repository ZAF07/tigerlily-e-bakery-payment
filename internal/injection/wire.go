//go:build wireinject
// +build wireinject

package injection

import (
	"database/sql"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/injection/providers"
	"github.com/google/wire"
)

func GetGeneralConfig() *config.GeneralConfig {
	wire.Build(providers.GeneralConfigProvider)
	return &config.GeneralConfig{}
}

func GetPaymentDBInstance() *sql.DB {
	wire.Build(providers.PaymentDBInstanceProvider)
	return config.LoadConfig().PaymentDB
}

func GetStripeServiceConfig() *config.StripeService {
	wire.Build(providers.StripeConfigProvider)
	return &config.LoadConfig().GeneralConfig.StripeService
}
