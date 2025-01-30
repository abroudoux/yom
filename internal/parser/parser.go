package parser

import (
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
)

type Item = types.Item

func ParseLines(lines []string) []Item {
	var parsedItems []Item
	for _, line := range lines {
		if strings.HasPrefix(line, ">>>") {
			continue
		}

		blocks := strings.Fields(line)
		if len(blocks) < 1 {
			logs.WarnMsg("Ignoring invalid line: %s")
			continue
		}

		productName := strings.Join(blocks[:len(blocks) - 1], " ")
		productPrice := blocks[len(blocks) - 3]
		parsedItems = append(parsedItems, Item{Name: productName, Price: productPrice, Quantity: 1})
	}

	return parsedItems
}