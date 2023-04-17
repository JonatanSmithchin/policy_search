package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	Mysql "police_search/Databases"
)

type Config struct {
	DB Mysql.DBConfig
}

var Conf = &Config{}

func InitConfig(configPath string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("unmarshal config failed: %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			log.Fatalf("unmarshal config failed: %v", err)
		}
	})
}
