package middleware

import (
	"os"

	"github.com/spf13/viper"
)

var configLoadedForTest bool

func Load() {
	if os.Getenv("ENVIRONMENT") == "test" {
		viper.SetConfigName("test")
	} else {
		viper.SetConfigName("application")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("./../../configs")
	viper.AddConfigPath("./../../../configs")
	viper.AddConfigPath("./../../../../configs")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	initAppConfig()
	initDatabaseConfig()

}

func LoadForTest() {
	os.Setenv("ENVIRONMENT", "test")
	if !configLoadedForTest {
		Load()
	}
	configLoadedForTest = true
}
