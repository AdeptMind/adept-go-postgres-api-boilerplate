package logger

import (
	log "github.com/sirupsen/logrus"
)

func ConfigureLogger(logLevel string) {
	switch logLevel {
	case log.TraceLevel.String():
		log.SetLevel(log.TraceLevel)
	case log.DebugLevel.String():
		log.SetLevel(log.DebugLevel)
	case log.WarnLevel.String():
		log.SetLevel(log.WarnLevel)
	case log.ErrorLevel.String():
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
