package forms

import (
	"fmt"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
	"github.com/abroudoux/yom/internal/utils"
	"github.com/charmbracelet/huh"
)

type Person = types.Person
type Item = types.Item
type Choice = types.Choice

func getName(message string) string {
	var name string
	huh.NewInput().Title(message).Prompt("? ").Value(&name).Run()
	return name
}

func GetNames(names *[]string) {
	newName := getName("Add a new name.")

	nameAlreadySaved := utils.IsNameAlreadySaved(newName, *names)
	if nameAlreadySaved {
		logs.WarnMsg("This name is already registered, please enter another one.")
		GetNames(names)
		return
	}

	*names = append(*names, newName)
	logs.Info(fmt.Sprintf("%s has been added.", newName))
	addNewName := getConfirmation("Do you want to add a new person?")
	if addNewName {
		GetNames(names)
		return
	}
}

func getConfirmation(message string) bool {
	var confirm bool
	huh.NewConfirm().Title(message).Negative("No").Affirmative("Yes!").Value(&confirm).Run()
	return confirm
}

func SelectPayer(persons *[]Person) {
	var selectedName string

	options := make([]huh.Option[string], len(*persons))
	for i, person := range *persons {
		options[i] = huh.NewOption(person.Name, person.Name)
	}

	huh.NewSelect[string]().Title("Who has paid?").Options(options...).Value(&selectedName).Run()

	for i := range *persons {
		if (*persons)[i].Name == selectedName {
			(*persons)[i].HasPaid = true
			logs.Info(fmt.Sprintf("%s has paid.", (*persons)[i].Name))
			break
		}
	}
}

func selectOption(choices []Choice, title string) Choice {
	var optionSelected Choice

	options := make([]huh.Option[string], len(choices))
	for i, choice := range choices {
		options[i] = huh.NewOption(choice.Name, choice.Name)
	}

	huh.NewSelect[string]().Title(title).Options(options...).Value(&optionSelected.Name).Run()

	return optionSelected
}

func MakeDistribution(persons *[]Person, items []Item) error {
    choices := utils.CreateChoices(persons)
    if len(choices) == 0 {
        return fmt.Errorf("no choices available")
    }

    for _, item := range items {
        title := fmt.Sprintf("%s: %s€", item.Name, item.Price)
        optionSelected := selectOption(choices, title)

        priceItem, err := utils.ConvertPriceStringToFlat(item.Price)
        if err != nil {
            return err
        }

        var selectedChoice Choice
        for _, choice := range choices {
            if choice.Name == optionSelected.Name {
                selectedChoice = choice
                break
            }
        }

        if len(selectedChoice.Persons) == 0 {
            return fmt.Errorf("no persons associated with the selected choice")
        }

        splitPrice := priceItem / float64(len(selectedChoice.Persons))
        for _, person := range selectedChoice.Persons {
            person.Amount += splitPrice
        }
    }

    return nil
}

