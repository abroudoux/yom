package forms

import (
	"fmt"
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
	var personSelected types.Person

	options := make([]huh.Option[string], len(persons))
	for i, name := range persons {
		options[i] = huh.NewOption(name.Name, name.Name)
	}

	huh.NewSelect[string]().Title("Who has paid?").Options(options...).Value(&personSelected.Name).Run()
	return personSelected
}

func selectPerson(personsAndDuos *[]types.Person, title string) types.Person {
	var personSelected types.Person

	options := make([]huh.Option[string], len(*personsAndDuos))
	for i, person := range *personsAndDuos {
		options[i] = huh.NewOption(person.Name, person.Name)
	}

	huh.NewSelect[string]().Title(title).Options(options...).Value(&personSelected.Name).Run()

	return personSelected
}

func addItemPriceToPerson(personSelected string, persons *[]types.Person, priceItem float64) {
	for i := range *persons {
		if (*persons)[i].Name == personSelected {
			(*persons)[i].Amount += priceItem
			break
		}
	}
}

func MakeDistribution(persons *[]types.Person, items []types.Item) error {
	personsAndDuos := utils.CreatePersonsAndDuos(persons)

	for _, item := range items {
		title := fmt.Sprintf("%s: %s", item.Name, item.Price)
		personSelected := selectPerson(personsAndDuos, title)

		priceItem, err := utils.ConvertPriceStringToFlat(item.Price)
		if err != nil {
			return err
		}

		if strings.Contains(personSelected.Name, " & ") {
			names := strings.Split(personSelected.Name, " & ")
			for _, name := range names {
				addItemPriceToPerson(name, persons, priceItem / 2)
			}
		} else {
			addItemPriceToPerson(personSelected.Name, persons, priceItem)
		}
	}

	return nil
}
