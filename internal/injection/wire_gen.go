// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/injection/providers"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/repository/checkout"
)

// Injectors from wire.go:

func GetGeneralConfig() *config.GeneralConfig {
	generalConfig := providers.GeneralConfigProvider()
	return generalConfig
}

func GetApplicationConfig() *config.ApplicationConfig {
	applicationConfig := providers.ApplicationConfigProvider()
	return applicationConfig
}

func GetLogger() *logger.Logger {
	loggerLogger := providers.LoggerProvider()
	return loggerLogger
}

func GetPaymentDBInstance() checkout.CheckoutDBInterface {
	checkoutDBInterface := providers.PaymentDBInstanceProviderInterface()
	return checkoutDBInterface
}

func GetStripeServiceConfig() *config.StripeService {
	stripeService := providers.StripeConfigProvider()
	return stripeService
}
