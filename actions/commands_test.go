package actions

import (
	"geektrust/helpers"
	"geektrust/models"
	"testing"
)

var familyTree models.Family

func TestMaternalAuntCommandPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Yodhan", helpers.GenderFemale, false)
	if response != "Tritha" {
		t.Errorf("Expected Tritha, got %v", response)
	}

}

func TestMaternalAuntCommandNegativePersonNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Todhan", helpers.GenderFemale, false)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestMaternalAuntCommandNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Jaya", helpers.GenderFemale, false)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestPaternalAuntCommandPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Vritha", helpers.GenderFemale, true)
	if response != "Satya" {
		t.Errorf("Expected Satya, got %v", response)
	}
}

func TestPaternalAuntCommandNegativePersonNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Todhan", helpers.GenderFemale, true)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestPaternalAuntCommandNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Yodhan", helpers.GenderFemale, true)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestMaternalUncleCommandPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Yodhan", helpers.GenderMale, false)
	if response != "Vritha" {
		t.Errorf("Vritha Tritha, got %v", response)
	}
}

func TestMaternalUncleCommandNegativePersonNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Todhan", helpers.GenderMale, false)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestMaternalUncleCommandNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Dritha", helpers.GenderMale, false)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestPaternalUncleCommandPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Dritha", helpers.GenderMale, true)
	if response != "Ish Vich Aras" {
		t.Errorf("Expected Ish Vich Aras, got %v", response)
	}
}

func TestPaternalUncleCommandNegativePersonNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Todhan", helpers.GenderMale, true)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestPaternalUncleCommandNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := UncleAuntCommand(familyTree, "Yodhan", helpers.GenderMale, true)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestBrotherInLawPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := InLawCommand(familyTree, "Jaya", helpers.GenderMale)
	if response != "Vritha" {
		t.Errorf("Expected Vritha, got %v", response)
	}
}
func TestBrotherInLawPositiveMultiple(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := InLawCommand(familyTree, "Amba", helpers.GenderMale)
	if response != "Ish Vich Aras" {
		t.Errorf("Expected Ish Vich Aras, got %v", response)
	}
}
func TestBrotherInLawNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := InLawCommand(familyTree, "Yodhan", helpers.GenderMale)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestBrotherInLawNegativePersonNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := InLawCommand(familyTree, "Todhan", helpers.GenderMale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestSisterInLawPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := InLawCommand(familyTree, "Jaya", helpers.GenderFemale)
	if response != "Tritha" {
		t.Errorf("Expected Tritha, got %v", response)
	}
}

func TestSisterInLawNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := InLawCommand(familyTree, "Yodhan", helpers.GenderFemale)
	if response != "NONE" {
		t.Errorf("Expected NONE, got %v", response)
	}
}

func TestSisterInLawNegativePersonNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := InLawCommand(familyTree, "Todhan", helpers.GenderFemale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, got %v", response)
	}
}

func TestSonPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := ChildCommand(familyTree, "Jaya", helpers.GenderMale)
	if response != "Yodhan" {
		t.Errorf("Expected Yodhan, received %v", response)
	}
}

func TestSonNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := ChildCommand(familyTree, "Yodhan", helpers.GenderMale)
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestSonNegativeNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := ChildCommand(familyTree, "Todhan", helpers.GenderMale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}

func TestDaughterPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := ChildCommand(familyTree, "Amba", helpers.GenderFemale)
	if response != "Dritha Tritha" {
		t.Errorf("Expected Dritha Tritha, received %v", response)
	}
}

func TestDaughterNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := ChildCommand(familyTree, "Yodhan", helpers.GenderFemale)
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestDaughterNegativeNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := ChildCommand(familyTree, "Todhan", helpers.GenderFemale)
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}

func TestSiblingPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := SiblingsCommand(familyTree, "Vila")
	if response != "Chika" {
		t.Errorf("Expected Chika, received %v", response)
	}
}
func TestSiblingPositiveMultiple(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := SiblingsCommand(familyTree, "Chit")
	if response != "Ish Vich Aras Satya" {
		t.Errorf("Expected Ish Vich Aras Satya, received %v", response)
	}
}
func TestSiblingNegativeNone(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := SiblingsCommand(familyTree, "Yodhan")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}
func TestSiblingNegativeNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := SiblingsCommand(familyTree, "Villa")
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}

func TestAddChildToFamilyPositive(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := AddChildToFamily(familyTree, "Dritha", "Disha", "Female")
	if response != "CHILD_ADDITION_SUCCEEDED" {
		t.Errorf("Expected CHILD_ADDITION_SUCCEEDED, received %v", response)
	}
}

func TestAddChildToFamilyNegativeDuplicate(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := AddChildToFamily(familyTree, "Dritha", "Satya", "Female")
	if response != "CHILD_ADDITION_FAILED" {
		t.Errorf("Expected CHILD_ADDITION_FAILED, received %v", response)
	}
}

func TestAddChildToFamilyNegativeNotMother(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := AddChildToFamily(familyTree, "Aras", "Disha", "Female")
	if response != "CHILD_ADDITION_FAILED" {
		t.Errorf("Expected CHILD_ADDITION_FAILED, received %v", response)
	}
}

func TestAddChildToFamilyNegativeNotFound(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := AddChildToFamily(familyTree, "Aaras", "Disha", "Female")
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}
func TestRelationshipNotExist(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Prasang", "Brother")
	if response != "PERSON_NOT_FOUND" {
		t.Errorf("Expected PERSON_NOT_FOUND, received %v", response)
	}
}
func TestRelationshipSibling(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Siblings")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipChildDaughter(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Daughter")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipChildSon(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Son")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipInLaw(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Brother-In-Law")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}
func TestRelationshipInLaw2(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Sister-In-Law")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}
func TestRelationshipMaternalAunt(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Maternal-Aunt")
	if response != "Tritha" {
		t.Errorf("Expected Tritha, received %v", response)
	}
}

func TestRelationshipPaternalAunt(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Paternal-Aunt")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipMaternalUncle(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Maternal-Uncle")
	if response != "Vritha" {
		t.Errorf("Expected Vritha, received %v", response)
	}
}

func TestRelationshipPaternalUncle(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "Paternal-Uncle")
	if response != "NONE" {
		t.Errorf("Expected NONE, received %v", response)
	}
}

func TestRelationshipDefault(t *testing.T) {

	familyTree.PopulateFamilyTree()
	response := GetRelationshipOfFamily(familyTree, "Yodhan", "IPaternal-Uncle")
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

	familyTree.PopulateFamilyTree()
	result := RunCommands(input, familyTree)
	if !result {
		t.Errorf("Something is wrong")
	}
}
