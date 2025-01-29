package utils

import "slices"

func IsPersonAlreadySaved(newPerson string, persons []string) bool {
	return slices.Contains(persons, newPerson)
}