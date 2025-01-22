package main

import (
	"fmt"
	"os"

	logs "github.com/abroudoux/yom/internal/logs"
)

func main() {
	if len(os.Args) < 1 {
		logs.Error("You need to provide a file path as an argument.")
		os.Exit(1)
	}

	filePath := os.Args[1]

	msg := fmt.Sprintf("Reading file %s", filePath)
	logs.Info(msg)
}