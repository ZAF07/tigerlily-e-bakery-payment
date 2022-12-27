package config

import (
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	repo "github.com/ZAF07/tigerlily-e-bakery-payment/internal/repository/checkout"
)

// 💡 TODO: DB field should be an interface with all possible DB methods so that we can implement diff DB instances upon start-up

type ApplicationConfig struct {
	GeneralConfig GeneralConfig
	// PaymentDB     *gorm.DB
	// PaymentDB     *sql.DB
	PaymentDB     repo.CheckoutDBInterface
	DefaultLogger *logger.Logger
}

type GeneralConfig struct {
	Port          string        `mapstructure:"port" json:"port"`
	Environment   string        `mapstructure:"environment" json:"environment"`
	Logger        string        `mapstructure:"logger" json:"logger"`
	DBType        string        `mapstructure:"db_type" json:"db_type"`
	ServerConfig  ServerConfig  `mapstructure:"server_config" json:"server_config"`
	PaymentDB     PaymentDB     `mapstructure:"payment_db" json:"payment_db"`
	StripeService StripeService `mapstructure:"stripe_service" json:"stripe_service"`
}

type PaymentDB struct {
	SSL         string `mapstructure:"ssl" json:"ssl"`
	Port        string `mapstructure:"port" json:"port"`
	Host        string `mapstructure:"host" json:"host"`
	User        string `mapstructure:"user" json:"user"`
	Name        string `mapstructure:"name" json:"name"`
	Password    string `mapstructure:"password" json:"password"`
	MaxConn     int    `mapstructure:"max_conns" json:"max_conns"`
	MaxIdleConn int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
}

type ServerConfig struct {
	ReadTimeout  int      `mapstructure:"read_timeout" json:"read_timeout"`
	WriteTimeout int      `mapstructure:"write_timeout" json:"write_timeout"`
	AllowOrigins []string `mapstructure:"allow_origins" json:"allow_origins"`
	AllowMethods []string `mapstructure:"allow_methods" json:"allow_methods"`
	AllowHeaders []string `mapstructure:"allow_headers" json:"allow_headers"`
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

func (d *PaymentDB) GetPostgresDBString() string {
	if d.Password != "" {
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", d.Host, d.User, d.Password, d.Name, d.Port, d.SSL)
	}
	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s", d.Host, d.User, d.Name, d.Port, d.SSL)
}
