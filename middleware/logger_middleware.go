package middleware

import (
	"net/http"
	"time"

	"github.com/Axrous/mnc/helper"
	"github.com/Axrous/mnc/model/web"
)

type LoggerMiddleware struct {
	Handler http.Handler
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (l *loggingResponseWriter) WriteHeader(status int)  {
	l.statusCode = status
	l.ResponseWriter.WriteHeader(status)
}

func NewLoggerMiddleware(handler http.Handler) http.Handler {
	return &LoggerMiddleware{
		Handler: handler,
	}
}

func (middleware *LoggerMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	recorder := &loggingResponseWriter{
		ResponseWriter: writer,
		statusCode:     200,
	}
	startTime := time.Now()
	middleware.Handler.ServeHTTP(recorder, request)

	endTime := time.Since(startTime)
	logRequest := web.LogRequest{
		StartTime:  startTime,
		EndTime:    endTime,
		StatusCode: recorder.statusCode,
		ClientIp:   request.RemoteAddr,
		Method:     request.Method,
		Path:       request.RequestURI,
		UserAgent:  request.UserAgent(),
	}

	helper.GenerateLogMiddleware(logRequest)
}