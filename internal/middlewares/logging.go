package middlewares

import (
	"net/http"
	"time"

	"github.com/mvigor/metricsd/internal/utils"
)

func WithLogging(h http.Handler) http.Handler {
	logger := utils.GetLogger()
	sugar := logger.Sugar()
	defer sugar.Sync()

	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &ResponseData{
			Status: 0,
			Size:   0,
		}
		lw := LoggingResponseWriter{
			ResponseWriter: w,
			ResponseData:   responseData,
		}
		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		sugar.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"status", responseData.Status,
			"duration", duration,
			"size", responseData.Size,
		)
	}
	return http.HandlerFunc(logFn)
}
