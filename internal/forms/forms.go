package forms

import (
	"slices"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/charmbracelet/huh"
)

func getPerson(message string) string {
	var person string
	huh.NewInput().Title(message).Prompt("? ").Value(&person).Run()
	return person
}

func GetPersons(persons *[]string) {
	newPerson := getPerson("Add a person")

	personAlreadySaved := isPersonAlreadySaved(newPerson, *persons)
	if personAlreadySaved {
		logs.Warn("This user is already registered, please enter another one: ", "")
		GetPersons(persons)
		return
	}

	*persons = append(*persons, newPerson)
	addNewPerson := getConfirmation("Do you want to add a new person?")
	if addNewPerson {
		GetPersons(persons)
		return
	}
}

func isPersonAlreadySaved(newPerson string, persons []string) bool {
	return slices.Contains(persons, newPerson)
}

func getConfirmation(message string) bool {
	var confirm bool
	huh.NewConfirm().Title(message).Affirmative("Yes!").Negative("No.").Value(&confirm).Run()
	return confirm
}

func SelectPayer(persons []string) string {
	var payer string
	huh.NewSelect[string]().Title("Who has paid?").Options(huh.NewOption("Arthur", "Arthur")).Value(&payer).Run()
	return payer
}