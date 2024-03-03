package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

func mustGetString(key string) string {
	mustHave(key)
	return viper.GetString(key)
}

func mustGetDurationMs(key string) time.Duration {
	return time.Millisecond * time.Duration(mustGetInt(key))
}

func mustGetDurationMinute(key string) time.Duration {
	return time.Minute * time.Duration(mustGetInt(key))
}

func mustGetInt(key string) int {
	mustHave(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid Integer value", key))
	}

	return v
}

func mustGetFloat(key string) float64 {
	v, err := strconv.ParseFloat(viper.GetString(key), 64)
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid float value", key))
	}

	return v
}

func mustHave(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}
