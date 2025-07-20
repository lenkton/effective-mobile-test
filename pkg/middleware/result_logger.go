package middleware

import (
	"log"
	"net/http"
	"time"
)

type ResultLogger struct {
	next http.Handler
}

func NewResultLogger(next http.Handler) ResultLogger {
	return ResultLogger{next: next}
}

// ServeHTTP implements http.Handler.
func (resultLogger ResultLogger) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	t0 := time.Now()
	spy := spyWriter{ResponseWriter: writer}
	resultLogger.next.ServeHTTP(&spy, request)
	t1 := time.Now()
	log.Printf("[%d] [%v]: %s\n", spy.code, t1.Sub(t0), request.RequestURI)
}

type spyWriter struct {
	http.ResponseWriter
	code int
}

func (spy *spyWriter) WriteHeader(code int) {
	spy.code = code
	spy.ResponseWriter.WriteHeader(code)
}
