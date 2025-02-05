package forms

import (
	"fmt"
	"strings"

	"github.com/abroudoux/yom/internal/logs"
	"github.com/abroudoux/yom/internal/types"
	"github.com/abroudoux/yom/internal/utils"
	"github.com/charmbracelet/huh"
)

type Person = types.Person
type Item = types.Item

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

func SelectPayer(persons []Person) Person {
	var personSelected Person

	options := make([]huh.Option[string], len(persons))
	for i, name := range persons {
		options[i] = huh.NewOption(name.Name, name.Name)
	}

	huh.NewSelect[string]().Title("Who has paid?").Options(options...).Value(&personSelected.Name).Run()
	return personSelected
}

func selectPerson(personsAndDuos *[]Person, title string) Person {
	var personSelected Person

	options := make([]huh.Option[string], len(*personsAndDuos))
	for i, person := range *personsAndDuos {
		options[i] = huh.NewOption(person.Name, person.Name)
	}

	huh.NewSelect[string]().Title(title).Options(options...).Value(&personSelected.Name).Run()

	return personSelected
}

func MakeDistribution(persons *[]Person, items []Item) error {
	personsAndDuos := utils.CreatePersonsAndDuos(persons)

	for _, item := range items {
		title := fmt.Sprintf("%s: %s€", item.Name, item.Price)
		personSelected := selectPerson(personsAndDuos, title)

		priceItem, err := utils.ConvertPriceStringToFlat(item.Price)
		if err != nil {
			return err
		}

		if strings.Contains(personSelected.Name, " & ") {
			names := strings.Split(personSelected.Name, " & ")
			for _, name := range names {
				utils.AddItemPriceToPerson(name, persons, priceItem / 2)
			}
		} else {
			utils.AddItemPriceToPerson(personSelected.Name, persons, priceItem)
		}
	}

	return nil
}
