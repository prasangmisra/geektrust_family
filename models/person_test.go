package models

import (
	"geektrust/helpers"
	"testing"
)

func TestFindChildrenWithName(t *testing.T) {
	var children1, children2, children3 Person
	var mother, father Person
	children1.Name = "A"
	children2.Name = "B"
	children3.Name = "C"
	mother.Children = append(mother.Children, &children1, &children2, &children3)
	father.Spouce = &mother
	father.Gender = helpers.GenderMale
	mother.Gender = helpers.GenderFemale
	mother.Spouce = &father
	result := father.FindChildrenWithName("A")
	if result.Name != "A" {
		t.Errorf("Expected A, received %v", result.Name)
	}
}

func TestFindChildrenWithNameNegative(t *testing.T) {
	var children1, children2, children3 Person
	var mother, father Person
	children1.Name = "A"
	children2.Name = "B"
	children3.Name = "C"
	mother.Children = append(mother.Children, &children1, &children2, &children3)
	father.Spouce = &mother
	father.Gender = helpers.GenderMale
	mother.Gender = helpers.GenderFemale
	mother.Spouce = &father
	result := father.FindChildrenWithName("D")
	if result != nil {
		t.Errorf("Expected nil, received %v", result)
	}
}
