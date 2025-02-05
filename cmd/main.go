package main

import (
	"os"

	"github.com/abroudoux/yom/internal/forms"
	"github.com/abroudoux/yom/internal/items"
	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/utils"
)

func main() {
	if len(os.Args) < 1 {
		logs.Error("You need to provide a file path as an argument.", nil)
		os.Exit(1)
	}

	filePath := os.Args[1]
	items := items.GetItems(filePath)

	persons := forms.GetPersons()
	forms.SelectPayer(&persons)

	choices := utils.CreateChoices(&persons)
    if len(choices) == 0 {
        logs.ErrorMsg("No choice available")
		os.Exit(1)
    }

	err := forms.MakeDistribution(choices, items)
	if err != nil {
		logs.Error("Error while the distribution of items", err)
	}

	utils.PrintResults(&persons)
}