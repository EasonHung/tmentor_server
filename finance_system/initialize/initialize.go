package initialize

import (
	"log"
	"mentor_app/finance_system/config"
	"os"

	"github.com/spf13/viper"
)

var (
	GLOBAL_CONFIG config.ViperConfig
)

func init() {
	env := os.Getenv("mentor_env")
	viper.AddConfigPath("config/" + env)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

	// 反序列化config
	viper.Unmarshal(&GLOBAL_CONFIG)
}
