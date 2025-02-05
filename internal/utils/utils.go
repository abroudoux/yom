package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
)

type Person = types.Person
type Item = types.Item

func IsNameAlreadySaved(newName string, names []string) bool {
	return slices.Contains(names, newName)
}

func createDuos(persons []Person) []Person {
	var duos []Person

	for i := 0; i < len(persons); i++ {
		for j := i + 1; j < len(persons); j++ {
			firstName := persons[i].Name
			secondName := persons[j].Name
			duoName := firstName + " & " + secondName
			newDuo := Person{
				Name: duoName,
				Amount: 0.0,
			}
			duos = append(duos, newDuo)
		}
	} 

	return duos
}

func CreatePersonsAndDuos(persons *[]Person) *[]Person {
	duos := createDuos(*persons)
	personsAndDuos := append(*persons, duos...)
	return &personsAndDuos
}

func CreateAllCombinations(persons *[]Person) *[]Person {
	return persons
}

func CreatePersons(names []string) []Person {
	var persons []Person
	for _, name := range names {
		newPerson := Person{
			Name: name,
			Amount: 0.0,
		}
		persons = append(persons, newPerson)
	}

	return persons
}

func CreateItemDict(items []Item) map[string]string {
	itemsDict := make(map[string]string, len(items))
	for _, item := range items {
		itemsDict[item.Name] = item.Price
	}

	return itemsDict
}

func PrintResults(persons *[]Person, payer Person, ) {
	fmt.Println("===== Total =====")
	for _, person := range *persons {
		if person.Name == payer.Name {
			logs.Info(fmt.Sprintf("%s spend %v€.", person.Name, person.Amount))
			continue
		}
		logs.Info(fmt.Sprintf("%s owes %v€ to %s.", person.Name, person.Amount, payer.Name))
	}
}

func ConvertPriceStringToFlat(price string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(price, ",", "."), 64)
}

func AddItemPriceToPerson(personSelected string, persons *[]Person, priceItem float64) {
	for i := range *persons {
		if (*persons)[i].Name == personSelected {
			(*persons)[i].Amount += priceItem
			break
		}
	}
}
