package logger

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
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
	if viper.GetInt("app.log_level") >= LogLevelError {
		log.Println(message)
	}
}

// LogWarning will output information on info verbosity level
func LogWarning(message string) {
	if viper.GetInt("app.log_level") >= LogLevelWarning {
		log.Println(message)
	}
}

// LogInfo will output information on info verbosity level
func LogInfo(message string) {
	if viper.GetInt("app.log_level") >= LogLevelInfo {
		log.Println(message)
	}
}

// LogDebug will output information on info verbosity level
func LogDebug(message string) {
	if viper.GetInt("app.log_level") >= LogLevelDebug {
		fmt.Println(message)
	}
}

