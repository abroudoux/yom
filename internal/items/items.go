package items

import (
	"fmt"
	"os"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/parser"
	"github.com/abroudoux/yom/internal/reader"
	"github.com/abroudoux/yom/internal/types"
)

type Item = types.Item

func GetItems(filePath string) []Item {
	logs.Info(fmt.Sprintf("Reading file %s", filePath))

    linesFile, err := reader.ReadFile(filePath)
	if err != nil {
		logs.Error("Error while reading file: %s", err)
		os.Exit(1)
	}

    items := parser.ParseLines(linesFile)
    return items
}