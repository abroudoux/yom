package forms

import (
	"github.com/charmbracelet/huh"
)

func getPerson(message string) string {
	var person string
	huh.NewInput().Title(message).Prompt("? ").Value(&person).Run()
	return person
}

func GetPersons(persons *[]string) {
	newPerson := getPerson("Add a person")
	*persons = append(*persons, newPerson)
	addNewPerson := getConfirmation("Do you want to add a new person?")
	if addNewPerson {
		GetPersons(persons)
	}
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