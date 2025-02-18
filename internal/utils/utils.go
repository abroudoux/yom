package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	. "github.com/abroudoux/yom/internal/types"
)

func IsNameAlreadySaved(newName string, names []string) bool {
	return slices.Contains(names, newName)
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

