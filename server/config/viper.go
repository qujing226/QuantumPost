package config

import (
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func InitRemoteViper() {
	err := viper.AddRemoteProvider("etcd3", "localhost:12379",
		"/quantum_post")
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
}
