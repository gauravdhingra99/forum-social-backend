package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	config "socialForumBackend/internal/config"
	"socialForumBackend/internal/handler"
	"socialForumBackend/internal/store"
	"socialForumBackend/service"
	"strconv"
	"syscall"
)

type Handler struct {
	forumHandler *handler.ForumHandler
}

func Start() error {
	postStore := store.NewStore()
	postService := service.NewPostService(postStore)
	h := &Handler{
		forumHandler: handler.NewForumHandler(postService),
	}
	startServer(router(h))
	return nil
}

func startServer(handler http.Handler) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(config.App.Port),
		Handler:      handler,
		ReadTimeout:  config.App.HTTPReadTimeout,
		WriteTimeout: config.App.HTTPWriteTimeout,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("failed to start the server : " + err.Error())
		}
	}()

	config.InitHystrixStream()
	defer config.StopHystrixStream()
	go func() {
		if err := http.ListenAndServe(config.App.HystrixStreamAddress, config.HystrixStreamHandler()); err != nil {
			panic("failed to start the hystrix server : " + err.Error())
		}
	}()

	<-stop
	err := server.Shutdown(context.Background())
	if err != nil {
		return
	}
}
