package server

import (
	"net/http"
	middleware "socialForumBackend/internal/config"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r = r.SkipClean(true)
	r.Use(middleware.Recover())
	r.Use(middleware.Trace())
	return r
}

func router(h *Handler) *mux.Router {
	r := newRouter()
	r.HandleFunc("/v1/user/post", h.forumHandler.CreatePost).Methods(http.MethodPost)
	r.HandleFunc("/v1/user/post", h.forumHandler.DeletePost).Methods(http.MethodDelete)
	r.HandleFunc("/v1/user/all/post", h.forumHandler.ListAllPosts).Methods(http.MethodGet)
	r.HandleFunc("/v1/user/live/feed/{user_id}", h.forumHandler.LiveNewsFeedQuery).Methods(http.MethodGet)

	return r
}
