package main

import (
	"fmt"
	"os"

	"github.com/abroudoux/yom/internal/forms"
	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/parser"
	"github.com/abroudoux/yom/internal/reader"
	"github.com/abroudoux/yom/internal/utils"
)

func main() {
	if len(os.Args) < 1 {
		logs.Error("You need to provide a file path as an argument.", nil)
		os.Exit(1)
	}

	filePath := os.Args[1]
	logs.Info(fmt.Sprintf("Reading file %s", filePath))

	linesFile, err := reader.ReadFile(filePath)
	if err != nil {
		logs.Error("Error while reading file: %s", err)
		os.Exit(1)
	}

	items := parser.ParseLines(linesFile)

	names := []string{}
	forms.GetNames(&names)
	if len(names) < 2 {
		logs.ErrorMsg("You need at least 2 persons to run this program.")
		os.Exit(1)
	}

	persons := utils.CreatePersons(names)
	forms.SelectPayer(&persons)

	err = forms.MakeDistribution(&persons, items)
	if err != nil {
		logs.Error("Error while the distribution of items", err)
	}

	utils.PrintResults(&persons)
}