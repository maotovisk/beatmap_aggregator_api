package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type LogWriter struct {
	StdOut  io.Writer
	LogFile io.Writer
}

func (lw *LogWriter) Write(p []byte) (n int, err error) {
	n, err = lw.StdOut.Write(p)
	if err != nil {
		return n, err
	}
	n, err = lw.LogFile.Write(p)
	return n, err
}

func (lw *LogWriter) Printf(format string, args ...any) {
	fmt.Fprintf(lw.StdOut, format, args...)
	fmt.Fprintf(lw.LogFile, format, args...)
}

// statusRecorder wraps the ResponseWriter to capture the status code
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code before calling the wrapped WriteHeader
func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

func (sr *statusRecorder) Write(b []byte) (int, error) {
	if sr.statusCode == 0 {
		sr.statusCode = http.StatusOK
	}

	return sr.ResponseWriter.Write(b)
}

// Logger Middleware
func LogHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w}

		h.ServeHTTP(rec, r)
		duration := time.Since(start)

		statusCode := rec.statusCode

		log.Printf("%s %s %s from %s - %d %s in %s\n",
			r.Method,
			r.URL.Path,
			r.Proto,
			r.RemoteAddr,
			statusCode,
			http.StatusText(statusCode),
			duration.String())
	})
}
