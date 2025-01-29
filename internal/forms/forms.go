package forms

import (
	"fmt"

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
	huh.NewConfirm().Title(message).Affirmative("Yes!").Negative("No.").Value(&confirm).Run()
	return confirm
}

func SelectPayer(persons []types.Person) types.Person {
	var payer types.Person

	options := make([]huh.Option[string], len(persons))
	for i, name := range persons {
		options[i] = huh.NewOption(name.Name, name.Name)
	}

	huh.NewSelect[string]().Title("Who has paid?").Options(options...).Value(&payer.Name).Run()
	return payer
}

func MakeDistribution(persons []types.Person, items []types.Item) map[string]string {
	return nil
}