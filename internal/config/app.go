package middleware

import (
	"time"
)

type AppConfig struct {
	Port                 int
	HTTPReadTimeout      time.Duration
	HTTPWriteTimeout     time.Duration
	HystrixStreamAddress string
}

var App AppConfig

func initAppConfig() {
	App = AppConfig{
		Port:                 mustGetInt("APP_PORT"),
		HTTPReadTimeout:      mustGetDurationMs("SERVER_HTTP_READ_TIMEOUT_MS"),
		HTTPWriteTimeout:     mustGetDurationMs("SERVER_HTTP_WRITE_TIMEOUT_MS"),
		HystrixStreamAddress: mustGetString("HYSTRIX_STREAM_ADDRESS"),
	}
}
