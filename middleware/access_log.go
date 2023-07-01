package middleware

import (
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type AccessLog struct{}

func NewAccessLog() *AccessLog {
	return &AccessLog{}
}

func (middle *AccessLog) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r.WithContext(r.Context()))

		logger, _ := zap.NewProduction()

		logger.Debug(
			"access-log",
			zap.String("method", r.Method),
			zap.String("path", r.RequestURI),
			zap.String("protocol", r.Proto),
			zap.String("addr", r.RemoteAddr),
			zap.String("user-agent", r.Header["User-Agent"][0]),
			zap.String("status-code", strconv.Itoa(rw.StatusCode)),
			zap.String("elapsed", time.Since(now).String()),
			zap.Int64("elapsed(ms)", time.Since(now).Milliseconds()),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
