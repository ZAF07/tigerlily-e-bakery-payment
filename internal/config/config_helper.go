package config

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func loadConfig() *ApplicationConfig {
	return appConfigLoader()
}

var appConfigLoader = loadConfigurations()

func loadConfigurations() func() *ApplicationConfig {
	config := &ApplicationConfig{}

	var once sync.Once

	return func() *ApplicationConfig {
		once.Do(func() {
			var configFilePath string

			flag.StringVar(&configFilePath, "config", "config.yml", "Path to configuration file...")
			flag.Parse()

			config = parseAndWatchConfigFile(configFilePath)
		})
		return config
	}
}

func parseAndWatchConfigFile(filepath string) *ApplicationConfig {

	AppConfig := &ApplicationConfig{}
	generalConfig := &GeneralConfig{}

	v := viper.New()
	v.SetConfigFile(filepath)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("ERROR READING CONFIG --> ", err)
	}
	unmarshalConfig(generalConfig, v)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("~~~ Config file '%+v' has been modified ~~~", e.Name)
		unmarshalConfig(generalConfig, v)
	})
	AppConfig.GeneralConfig = *generalConfig

	return AppConfig
}

func unmarshalConfig(config *GeneralConfig, v *viper.Viper) {
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("[CONFIG] Error unmarshaling app config : %+v\n", err)
	}
}
