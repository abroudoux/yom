package reader

import (
	"bufio"
	"os"
	"strings"

	"github.com/abroudoux/yom/internal/logs"
)

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logs.Error("Error while opening file: %s", err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		logs.Fatal("Error while scanning file: %s", err)
		return nil, err
	}

	return lines, nil
}