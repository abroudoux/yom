package logs

import "github.com/charmbracelet/log"

func Error(msg string) {
	log.Error(msg)
}

func Info(msg string) {
	log.Info(msg)
}