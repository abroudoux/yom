package forms

import (
	"fmt"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/utils"
	"github.com/charmbracelet/huh"
)

func getPerson(message string) string {
	var person string
	huh.NewInput().Title(message).Prompt("? ").Value(&person).Run()
	return person
}

func GetPersons(persons *[]string) {
	newPerson := getPerson("Add a person")

	personAlreadySaved := utils.IsPersonAlreadySaved(newPerson, *persons)
	if personAlreadySaved {
		logs.WarnMsg("This user is already registered, please enter another one.")
		GetPersons(persons)
		return
	}

	*persons = append(*persons, newPerson)
	logs.Info(fmt.Sprintf("%s has been added.", newPerson))
	addNewPerson := getConfirmation("Do you want to add a new person?")
	if addNewPerson {
		GetPersons(persons)
		return
	}
}

func getConfirmation(message string) bool {
	var confirm bool
	huh.NewConfirm().Title(message).Affirmative("Yes!").Negative("No.").Value(&confirm).Run()
	return confirm
}

func SelectPayer(persons []string) string {
	var payer string

	options := make([]huh.Option[string], len(persons))
	for i, person := range persons {
		options[i] = huh.NewOption(person, person)
	}

	huh.NewSelect[string]().Title("Who has paid?").Options(options...).Value(&payer).Run()
	return payer
}