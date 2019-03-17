package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HTTPLogger is logging values
func HTTPLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		LogInfo(fmt.Sprintf(
			"%s\t%s\t\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		))
	})
}

const (
	// LogLevelError will output only errors
	LogLevelError = iota
	// LogLevelWarning will output warnings and all above
	LogLevelWarning
	// LogLevelInfo will output infos and all above
	LogLevelInfo
	// LogLevelDebug will output all information
	LogLevelDebug
)

// LogFatal is a wrapper for log.Fatal
func LogFatal(err error) {
	log.Fatal(err)
}

// LogError will output information on info verbosity level
func LogError(message error) {
	if Config.Debug.LogLevel >= LogLevelError {
		log.Println(message)
	}
}

// LogWarning will output information on info verbosity level
func LogWarning(message string) {
	if Config.Debug.LogLevel >= LogLevelWarning {
		log.Println(message)
	}
}

// LogInfo will output information on info verbosity level
func LogInfo(message string) {
	if Config.Debug.LogLevel >= LogLevelInfo {
		log.Println(message)
	}
}

// LogDebug will output information on info verbosity level
func LogDebug(message string) {
	if Config.Debug.LogLevel >= LogLevelDebug {
		fmt.Println(message)
	}
}
