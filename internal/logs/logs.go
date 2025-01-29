package logs

import (
	"github.com/charmbracelet/log"
)

func Error(msg string, err error) {
	log.Error(msg, err)
}

func ErrorMsg(msg string) {
	log.Error(msg)
}

func Fatal(msg string, err error) {
	log.Fatal(msg, err)
}

func FatalMsg(msg string) {
	log.Fatal(msg)
}

func Info(msg string) {
	log.Info(msg)
}

func Warn(msg string, val string) {
	log.Warn(msg, val)
}

func WarnMsg(msg string) {
	log.Warn(msg)
}