package forms

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
	"github.com/abroudoux/yom/internal/utils"
	"github.com/charmbracelet/huh"
)

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

func SelectPayer(persons []types.Person) types.Person {
	return selectPerson(persons, "Who has paid?")
}

func selectPerson(persons []types.Person, msg string) types.Person {
	var personSelected types.Person

	options := make([]huh.Option[string], len(persons))
	for i, name := range persons {
		options[i] = huh.NewOption(name.Name, name.Name)
	}

	huh.NewSelect[string]().Title(msg).Options(options...).Value(&personSelected.Name).Run()
	return personSelected
}

func attributeItem(persons []types.Person, item types.Item) types.Person {
	msg := fmt.Sprintf("%s (%s)", item.Name, item.Price)
	return selectPerson(persons, msg)
}

func MakeDistribution(persons *[]types.Person, items []types.Item) (error) {
	for _, item := range items {
		selectedPerson := attributeItem(*persons, item)
		priceItem, err := strconv.ParseFloat(strings.ReplaceAll(item.Price, ",", "."), 64)
		if err != nil {
			return err
		}

		for i := range *persons {
            if (*persons)[i].Name == selectedPerson.Name {
                (*persons)[i].Amount += priceItem
                break
            }
        }
	}

	return nil
}