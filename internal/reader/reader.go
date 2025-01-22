package reader

import (
	"errors"
	"fmt"
	"os"

	"github.com/abroudoux/yom/internal/logs"
)

func ReadFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		logs.Fatal("Error while reading file: %s", err)
	}
	fmt.Print(string(data))


	return []string{}, errors.New("Not implemented yet")
}