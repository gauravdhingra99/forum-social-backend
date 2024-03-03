package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type contextKey string

const (
	RequestID                = "X-Transaction-Id"
	ownerIDCtxKey contextKey = "OwnerID"
)

func Trace() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if r.URL.Path == "/ping" || strings.HasPrefix(r.URL.Path, "/doc") {
				next.ServeHTTP(w, r)
				return
			}

			requestID := r.Header.Get(RequestID)
			if requestID == "" {
				requestID = uuid.NewV4().String()
			}

			ctx = context.WithValue(ctx, ownerIDCtxKey, requestID)

			recorder := newResponseRecorder(w, r)
			next.ServeHTTP(recorder, r.WithContext(ctx))
		})
	}
}

type responseRecorder struct {
	r *http.Request
	w http.ResponseWriter

	startTime time.Time
	path      string
}

func newResponseRecorder(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	pathTemplate, _ := mux.CurrentRoute(r).GetPathTemplate()
	if pathTemplate == "" {
		pathTemplate = r.URL.String()
	}

	return &responseRecorder{
		r: r,
		w: w,

		path:      pathTemplate,
		startTime: time.Now(),
	}
}

func (r *responseRecorder) Header() http.Header {
	return r.w.Header()
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	return r.w.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.w.WriteHeader(statusCode)
}
