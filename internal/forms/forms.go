package forms

import "github.com/charmbracelet/huh"

func GetPayer() string {
	var payer string
	huh.NewInput().Title("Who has paid?").Prompt("? ").Value(&payer).Run()
	return payer
}

// func GetPersons() []string {}