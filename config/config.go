package config

import (
	"os"

	"github.com/spf13/viper"
)

type (
	JaegerConfig struct {
		CollectorEndpoint string
	}
)

func getViper() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yml")
	v.AddConfigPath("./config")
	return v
}

var (
	defaultViper *viper.Viper
)

func init() {
	v := getViper()
	v.SetConfigName("default")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	defaultViper = getViper()
	defaultViper.SetConfigName(os.Getenv("config_name"))
	// 将default中的配置全部以默认配置写入
	for k, v := range v.AllSettings() {
		defaultViper.SetDefault(k, v)
	}
	// 读取当前运行环境对应的配置
	err = defaultViper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// GetString viper get string
func GetString(key string) string {
	return defaultViper.GetString(key)
}

func GetJaegerConfig() JaegerConfig {
	prefix := "jaeger."

	jaegerConfig := JaegerConfig{
		CollectorEndpoint: defaultViper.GetString(prefix + "endpoint"),
	}
	return jaegerConfig
}
