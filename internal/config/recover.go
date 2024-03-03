package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Recover() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
