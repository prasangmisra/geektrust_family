package actions

import (
	"geektrust/helpers"
	"geektrust/models"
	"testing"
)

func TestMaternalAuntCommandPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Yodhan", helpers.GenderFemale, false)
	if response != "Tritha" {
		t.Errorf("Expected Tritha, got %v", response)
	}

}

func TestMaternalAuntCommandNegativePersonNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Todhan", helpers.GenderFemale, false)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestMaternalAuntCommandNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Jaya", helpers.GenderFemale, false)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestPaternalAuntCommandPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Vritha", helpers.GenderFemale, true)
	if response != "Satya" {
		t.Errorf("Expected Satya, got %v", response)
	}
}

func TestPaternalAuntCommandNegativePersonNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Todhan", helpers.GenderFemale, true)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestPaternalAuntCommandNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Yodhan", helpers.GenderFemale, true)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestMaternalUncleCommandPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Yodhan", helpers.GenderMale, false)
	if response != "Vritha" {
		t.Errorf("Vritha Tritha, got %v", response)
	}
}

func TestMaternalUncleCommandNegativePersonNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Todhan", helpers.GenderMale, false)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestMaternalUncleCommandNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Dritha", helpers.GenderMale, false)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestPaternalUncleCommandPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Dritha", helpers.GenderMale, true)
	if response != "Ish Vich Aras" {
		t.Errorf("Expected Ish Vich Aras, got %v", response)
	}
}

func TestPaternalUncleCommandNegativePersonNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Todhan", helpers.GenderMale, true)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestPaternalUncleCommandNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := UncleAuntCommand(family, "Yodhan", helpers.GenderMale, true)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestBrotherInLawPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := InLawCommand(family, "Jaya", helpers.GenderMale)
	if response != "Vritha" {
		t.Errorf("Expected Vritha, got %v", response)
	}
}
func TestBrotherInLawPositiveMultiple(t *testing.T) {
	var family models.Family
	family.Init()
	response := InLawCommand(family, "Amba", helpers.GenderMale)
	if response != "Ish Vich Aras" {
		t.Errorf("Expected Ish Vich Aras, got %v", response)
	}
}
func TestBrotherInLawNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := InLawCommand(family, "Yodhan", helpers.GenderMale)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestBrotherInLawNegativePersonNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := InLawCommand(family, "Todhan", helpers.GenderMale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestSisterInLawPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := InLawCommand(family, "Jaya", helpers.GenderFemale)
	if response != "Tritha" {
		t.Errorf("Expected Tritha, got %v", response)
	}
}

func TestSisterInLawNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := InLawCommand(family, "Yodhan", helpers.GenderFemale)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestSisterInLawNegativePersonNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := InLawCommand(family, "Todhan", helpers.GenderFemale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestSonPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := ChildCommand(family, "Jaya", helpers.GenderMale)
	if response != "Yodhan" {
		t.Errorf("Expected Yodhan, received %v", response)
	}
}

func TestSonNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := ChildCommand(family, "Yodhan", helpers.GenderMale)
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestSonNegativeNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := ChildCommand(family, "Todhan", helpers.GenderMale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}

func TestDaughterPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := ChildCommand(family, "Amba", helpers.GenderFemale)
	if response != "Dritha Tritha" {
		t.Errorf("Expected Dritha Tritha, received %v", response)
	}
}

func TestDaughterNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := ChildCommand(family, "Yodhan", helpers.GenderFemale)
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestDaughterNegativeNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := ChildCommand(family, "Todhan", helpers.GenderFemale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}

func TestSiblingPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := SiblingsCommand(family, "Vila")
	if response != "Chika" {
		t.Errorf("Expected Chika, received %v", response)
	}
}
func TestSiblingPositiveMultiple(t *testing.T) {
	var family models.Family
	family.Init()
	response := SiblingsCommand(family, "Chit")
	if response != "Ish Vich Aras Satya" {
		t.Errorf("Expected Ish Vich Aras Satya, received %v", response)
	}
}
func TestSiblingNegativeNone(t *testing.T) {
	var family models.Family
	family.Init()
	response := SiblingsCommand(family, "Yodhan")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}
func TestSiblingNegativeNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := SiblingsCommand(family, "Villa")
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}

func TestAddChildToFamilyPositive(t *testing.T) {
	var family models.Family
	family.Init()
	response := AddChildToFamily(family, "Dritha", "Disha", "Female")
	if response != "CHILD_ADDITION_SUCCEEDED" {
		t.Errorf("Expected CHILD_ADDITION_SUCCEEDED, received %v", response)
	}
}

func TestAddChildToFamilyNegativeDuplicate(t *testing.T) {
	var family models.Family
	family.Init()
	response := AddChildToFamily(family, "Dritha", "Satya", "Female")
	if response != "CHILD_ADDITION_FAILED" {
		t.Errorf("Expected CHILD_ADDITION_FAILED, received %v", response)
	}
}

func TestAddChildToFamilyNegativeNotMother(t *testing.T) {
	var family models.Family
	family.Init()
	response := AddChildToFamily(family, "Aras", "Disha", "Female")
	if response != "CHILD_ADDITION_FAILED" {
		t.Errorf("Expected CHILD_ADDITION_FAILED, received %v", response)
	}
}

func TestAddChildToFamilyNegativeNotFound(t *testing.T) {
	var family models.Family
	family.Init()
	response := AddChildToFamily(family, "Aaras", "Disha", "Female")
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}
func TestRelationshipNotExist(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Prasang", "Brother")
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}
func TestRelationshipSibling(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Siblings")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipChildDaughter(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Daughter")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipChildSon(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Son")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipInLaw(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Brother-In-Law")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}
func TestRelationshipInLaw2(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Sister-In-Law")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}
func TestRelationshipMaternalAunt(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Maternal-Aunt")
	if response != "Tritha" {
		t.Errorf("Expected Tritha, received %v", response)
	}
}

func TestRelationshipPaternalAunt(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Paternal-Aunt")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipMaternalUncle(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Maternal-Uncle")
	if response != "Vritha" {
		t.Errorf("Expected Vritha, received %v", response)
	}
}

func TestRelationshipPaternalUncle(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "Paternal-Uncle")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipDefault(t *testing.T) {
	var family models.Family
	family.Init()
	response := GetRelationshipOfFamily(family, "Yodhan", "IPaternal-Uncle")
	if response != "" {
		t.Errorf("Expected , received %v", response)
	}
}

func TestRunCommands(t *testing.T) {
	var input [][]string
	input1 := []string{"ADD_CHILD", "Anga", "Prasang", "Male"}
	input2 := []string{"GET_RELATIONSHIP", "Anga", "Son"}
	input3 := []string{"ABC", "DEF", "GHI"}
	input = append(input, input1, input2, input3)
	var family models.Family
	family.Init()
	result := RunCommands(input, family)
	if !result {
		t.Errorf("Something is wrong")
	}
}
