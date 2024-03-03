package middleware

import "github.com/afex/hystrix-go/hystrix"

var hystrixStreamHandler *hystrix.StreamHandler

func HystrixStreamHandler() *hystrix.StreamHandler {
	if hystrixStreamHandler == nil {
		panic("please initialize hystrix stream")
	}

	return hystrixStreamHandler
}

func InitHystrixStream() {
	hystrixStreamHandler = hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
}

func StopHystrixStream() {
	if hystrixStreamHandler == nil {
		return
	}

	hystrixStreamHandler.Stop()
}
