package main

import (
	"fmt"
	"os"

	"github.com/abroudoux/yom/internal/forms"
	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/parser"
	"github.com/abroudoux/yom/internal/reader"
)

func main() {
	if len(os.Args) < 1 {
		logs.Error("You need to provide a file path as an argument.", nil)
		os.Exit(1)
	}

	filePath := os.Args[1]
	logs.Info(fmt.Sprintf("Reading file %s", filePath))

	items, err := reader.ReadFile(filePath)
	if err != nil {
		logs.Error("Error while reading file: %s", err)
		os.Exit(1)
	}

	parsedItems := parser.ParseItems(items)
	for _, item := range parsedItems {
		logs.Info(fmt.Sprintf("Item: %s, Price: %s", item.Name, item.Price))
	}

	persons := []string{}
	forms.GetPersons(&persons)
	for _, v := range persons {
		println(v)
	}

	if len(persons) < 2 {
		logs.ErrorMsg("You need at least 2 persons to run this program.")
		os.Exit(1)
	}

	payer := forms.SelectPayer(persons)
	logs.Info(fmt.Sprintf("Person %s has paid.", payer))
}