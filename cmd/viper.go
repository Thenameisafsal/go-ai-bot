package cmd

import (
	"log"

	"github.com/spf13/viper"
)

func GetFromEnv(key string) string {
	viper.SetConfigFile("/Users/afsalahamada/Documents/go ai bot/cmd/.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	value := viper.Get(key).(string)
	return value
}
