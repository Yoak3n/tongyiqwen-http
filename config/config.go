package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port   int    `yaml:"port"`
	Model  string `yaml:"model"`
	APIkey string `yaml:"api_key"`
}

var (
	conf *Config
	v    *viper.Viper
)

func init() {
	v = viper.New()
	conf = new(Config)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetDefault("port", 20104)
	v.SetDefault("model", "qwen-turbo")
	err := v.ReadInConfig()
	if err != nil {
		return
	}
	// 直接反序列化失败
	//err = v.Unmarshal(conf)
	//if err != nil {
	//	return
	//}
	v.WatchConfig()

	conf.Port = v.GetInt("port")
	conf.APIkey = v.GetString("api_key")
	conf.Model = v.GetString("model")

}
func GetConfig() *Config {
	return conf
}
