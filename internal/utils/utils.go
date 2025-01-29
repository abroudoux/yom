package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
)

func IsNameAlreadySaved(newName string, names []string) bool {
	return slices.Contains(names, newName)
}

func createDuos(persons []types.Person) []types.Person {
	var duos []types.Person

	for i := 0; i < len(persons); i++ {
		for j := i + 1; j < len(persons); j++ {
			firstName := persons[i].Name
			secondName := persons[j].Name
			duoName := firstName + " & " + secondName
			newDuo := types.Person{
				Name: duoName,
				Amount: 0.0,
			}
			duos = append(duos, newDuo)
		}
	} 

	return duos
}

func CreatePersonsAndDuos(persons *[]types.Person) *[]types.Person {
	duos := createDuos(*persons)
	personsAndDuos := append(*persons, duos...)
	return &personsAndDuos
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

func PrintResults(persons *[]types.Person, payer types.Person, ) {
	for _, person := range *persons {
		if person.Name == payer.Name {
			break
		}
		logs.Info(fmt.Sprintf("%s owes %v€ to %s", person.Name, person.Amount, payer.Name))
	}
}

func ConvertPriceStringToFlat(price string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(price, ",", "."), 64)
} 