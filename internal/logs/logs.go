package logs

import (
	"github.com/charmbracelet/log"
)

func Error(msg string, err error) {
	log.Error(msg, err)
}

func Fatal(msg string, err error) {
	log.Fatal(msg, err)
}

func Info(msg string) {
	log.Info(msg)
}