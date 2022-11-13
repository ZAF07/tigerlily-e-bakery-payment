package config

import "fmt"

type ApplicationConfig struct {
	GeneralConfig GeneralConfig
}

type GeneralConfig struct {
	Port          string
	Environment   string
	Server        ServerConfig  `mapstructure:"server_config" json:"server_config"`
	PaymentDB     PaymentDB     `mapstructure:"payment_db" json:"payment_db"`
	StripeService StripeService `mapstructure:"stripe_service" json:"stripe_service"`
}

type PaymentDB struct {
	SSLDisable string `mapstructure:"ssl" json:"ssl"`
	Port       string `mapstructure:"port" json:"port"`
	Host       string `mapstructure:"host" json:"host"`
	User       string `mapstructure:"user" json:"user"`
	Password   string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	ReadTimeout  int      `mapstructure:"read_timeout" json:"read_timeout"`
	WriteTimeout int      `mapstructure:"write_timeout" json:"write_timeout"`
	AllowOrigins []string `mapstructure:"allow_origins" json:"allow_origins"`
	AllowMethods []string `mapstructure:"allow_methods" json:"allow_methods"`
}

type StripeService struct {
	Domain     string `mapstructure:"domain" json:"domain"`
	Authkey    string `mapstructure:"key" json:"key"`
	Currency   string `mapstructure:"currency" json:"currency"`
	SuccessURL string `mapstructure:"success_url" json:"success_url"`
	CancelURL  string `mapstructure:"cancel_url" json:"cancel_url"`
}

func LoadConfig() *ApplicationConfig {
	return loadConfig()
}

func (c *ApplicationConfig) GetApplicationPort() string {
	port := fmt.Sprintf(":%s", c.GeneralConfig.Port)
	return port
}
