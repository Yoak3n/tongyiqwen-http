package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port            int    `yaml:"port"`
	AccessKey       string `yaml:"access_key"`
	AccessSecretKey string `yaml:"access_secret_key"`
	AppID           string `yaml:"app_id"`
	AgentKey        string `yaml:"agent_key"`
	Token           string `yaml:"token"`
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
	conf.AgentKey = v.GetString("agent_key")
	conf.AppID = v.GetString("app_id")
	conf.AccessKey = v.GetString("access_key")
	conf.AccessSecretKey = v.GetString("access_secret_key")
	conf.Port = v.GetInt("port")
	conf.Token = v.GetString("token")

}
func GetConfig() *Config {
	return conf
}

func RefreshToken(token string) {
	v.Set("token", token)
	err := v.SafeWriteConfig()
	if err != nil {
		return
	}
}
