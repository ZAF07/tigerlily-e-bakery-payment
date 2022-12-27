package providers

import (
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/config"
	repo "github.com/ZAF07/tigerlily-e-bakery-payment/internal/repository/checkout"
)

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
