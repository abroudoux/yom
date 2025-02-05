package utils

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/parser"
	"github.com/abroudoux/yom/internal/reader"
	. "github.com/abroudoux/yom/internal/types"
)

func GetItems(filePath string) []Item {
	logs.Info(fmt.Sprintf("Reading file %s", filePath))

    linesFile, err := reader.ReadFile(filePath)
	if err != nil {
		logs.Error("Error while reading file: %s", err)
		os.Exit(1)
	}

    items := parser.ParseLines(linesFile)
    return items
}

func CreateItemDict(items []Item) map[string]string {
	itemsDict := make(map[string]string, len(items))
	for _, item := range items {
		itemsDict[item.Name] = item.Price
	}

	return itemsDict
}

func IsNameAlreadySaved(newName string, names []string) bool {
	return slices.Contains(names, newName)
}

func CreatePersons(names []string) []Person {
	var persons []Person
	for _, name := range names {
		newPerson := Person{
			Name: name,
			Amount: 0.0,
			HasPaid: false,
		}
		persons = append(persons, newPerson)
	}

	return persons
}

func createDuosChoices(persons []Person) []Choice {
    var duosChoices []Choice

    for i := 0; i < len(persons); i++ {
        for j := i + 1; j < len(persons); j++ {
            firstName := persons[i].Name
            secondName := persons[j].Name
            duoName := firstName + " & " + secondName
            newDuoChoice := Choice{
                Name: duoName,
                Persons: []*Person{&persons[i], &persons[j]},
                Format: Duo,
            }
            duosChoices = append(duosChoices, newDuoChoice)
        }
    }

    return duosChoices
}

func createTrioChoices(persons *[]Person) Choice {
	trioName := fmt.Sprintf("%s, %s & %s", (*persons)[0].Name, (*persons)[1].Name, (*persons)[2].Name)
	trioChoice := Choice{
		Name:    trioName,
		Persons: []*Person{&(*persons)[0], &(*persons)[1], &(*persons)[2]},
		Format:  Trio,
	}
	return trioChoice
}

func CreateChoices(persons *[]Person) []Choice {
    var choices []Choice

    for i := range *persons {
        soloChoice := Choice{
            Name:    (*persons)[i].Name,
            Persons: []*Person{&(*persons)[i]},
            Format:  Solo,
        }
        choices = append(choices, soloChoice)
    }

    duosChoices := createDuosChoices(*persons)
    choices = append(choices, duosChoices...)

    if len(*persons) == 3 {
        trioChoice := createTrioChoices(persons)
        choices = append(choices, trioChoice)
    }

    return choices
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

func PrintResults(persons *[]Person) {
    fmt.Println("===== Total =====")
    var payer Person

    for _, person := range *persons {
        if person.HasPaid {
            payer = person
            break
        }
    }

    if payer.Name == "" {
        logs.WarnMsg("No payer found.")
        return
    }

    for _, person := range *persons {
        if person.HasPaid {
            logs.Info(fmt.Sprintf("%s paid %.2f€ in total.", person.Name, person.Amount))
        } else {
            amountOwed := person.Amount
            if amountOwed > 0 {
                logs.Info(fmt.Sprintf("%s owes %.2f€ to %s.", person.Name, amountOwed, payer.Name))
            } else {
                logs.Info(fmt.Sprintf("%s doesn't owe anything.", person.Name))
            }
        }
    }
}

