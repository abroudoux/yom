package utils

import (
	"fmt"
	"slices"

	"github.com/abroudoux/yom/internal/types"
)

func IsNameAlreadySaved(newName string, names []string) bool {
	return slices.Contains(names, newName)
}

func createDuos(names []string) []string {
	var duos []string

	for i := 0; i < len(names); i++ {
		for j := i + 1; j < len(names); j++ {
			firstName := names[i]
			secondName := names[j]
			duos = append(duos, fmt.Sprintf("%s & %s", firstName, secondName))
		}
	} 

	return duos
}

func CreateOptions(persons []string) []string {
	var options []string
	options = append(options, persons...)

	//duos := createDuos(persons)
	//options = append(options, duos...)

	return options
}

func CreatePersons(names []string) []types.Person {
	var persons []types.Person
	for _, name := range names {
		newPerson := types.Person{
			Name: name,
			Amount: 0.0,
		}
		persons = append(persons, newPerson)
	}

	return persons
}

func CreateItemDict(items []types.Item) map[string]string {
	itemsDict := make(map[string]string, len(items))
	for _, item := range items {
		itemsDict[item.Name] = item.Price
	}

	return itemsDict
}