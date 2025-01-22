package parser

import (
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
)

func ParseItems(items []string) []types.Item {
	var parsedItems []types.Item
	for _, item := range items {
		blocks := strings.Fields(item)
		if len(blocks) < 1 {
			logs.Warn("Ignoring invalid item: %s", item)
			continue
		}

		productName := strings.Join(blocks[:len(blocks) - 1], " ")
		productPrice := blocks[len(blocks) - 3]
		parsedItems = append(parsedItems, types.Item{Name: productName, Price: productPrice})
	}

	return parsedItems
}