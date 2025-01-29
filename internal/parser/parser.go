package parser

import (
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
)

func ParseLines(lines []string) []types.Item {
	var parsedItems []types.Item
	for _, line := range lines {
		blocks := strings.Fields(line)
		if len(blocks) < 1 {
			logs.Warn("Ignoring invalid line: %s", line)
			continue
		}

		productName := strings.Join(blocks[:len(blocks) - 1], " ")
		productPrice := blocks[len(blocks) - 3]
		parsedItems = append(parsedItems, types.Item{Name: productName, Price: productPrice, Quantity: 1})
	}

	return parsedItems
}