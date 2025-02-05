package parser

import (
	"regexp"
	"strings"

	"github.com/abroudoux/yom/internal/types"
)

type Item = types.Item

func ParseLines(lines []string) []Item {
	var parsedItems []Item
	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}

		if strings.HasPrefix(line, ">>>") {
			continue
		}

		re := regexp.MustCompile(`(.*)\s(\d+,\d+)\s€\s(\d+)$`)
		matches := re.FindStringSubmatch(line)

		if len(matches) == 4 {
			name := strings.TrimSpace(matches[1])
			price := matches[2]
			parsedItems = append(parsedItems, Item{Name: name, Price: price, Quantity: 1})
		}
	}

	return parsedItems
}
