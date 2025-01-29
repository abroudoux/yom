package utils

import (
	"fmt"
	"slices"
)

func IsPersonAlreadySaved(newPerson string, persons []string) bool {
	return slices.Contains(persons, newPerson)
}

func createDuos(persons []string) []string {
	var duos []string

	for i := 0; i < len(persons); i++ {
		for j := i + 1; j < len(persons); j++ {
			duos = append(duos, fmt.Sprintf("%s & %s", persons[i], persons[j]))
		}
	} 

	return duos
}

func CreateOptions(persons []string) []string {
	var options []string
	options = append(options, persons...)

	duos := createDuos(persons)
	options = append(options, duos...)

	return options
}