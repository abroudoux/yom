package main

import (
	"fmt"
	"os"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/reader"
)

func main() {
	if len(os.Args) < 1 {
		logs.Error("You need to provide a file path as an argument.", nil)
		os.Exit(1)
	}

	filePath := os.Args[1]
	logs.Info(fmt.Sprintf("Reading file %s", filePath))

	reader.ReadFile(filePath)
}