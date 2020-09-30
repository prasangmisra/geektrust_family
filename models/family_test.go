package models

import (
	"geektrust/helpers"
	"testing"
)

var familyTree Family

func TestStruct(t *testing.T) {
	if familyTree.Person != nil {
		t.Errorf("This should have been nil")
	}
}
func TestAddChild(t *testing.T) {
	var queen Person
	var child Person
	queen.Gender = "Female"
	queen.Name = "Anga"
	child.Name = "Lika"
	child.Gender = "Male"
	familyTree.Person = &queen
	queen.Children = append(queen.Children, &child)
	_, resultBool := familyTree.AddChild("Anga", "Disha1", "Female")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(familyTree.Person.Children))
	}
}

func TestInit(t *testing.T) {
	familyTree.PopulateFamilyTree()
	if len(familyTree.Person.Children) != 5 {
		t.Errorf("Expected 5, got %v", len(familyTree.Person.Children))
	}
}

func TestAddChildAfterInit(t *testing.T) {
	familyTree.PopulateFamilyTree()
	_, resultBool := familyTree.AddChild("Krithi", "Disha", "Female")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(familyTree.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}

func TestAddChildNegativeDuplicate(t *testing.T) {
	familyTree.PopulateFamilyTree()
	_, resultBool := familyTree.AddChild("Krithi", "Krithi", "Female")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(familyTree.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}
func TestAddChildNegativeNoMother(t *testing.T) {
	familyTree.PopulateFamilyTree()
	_, resultBool := familyTree.AddChild("Disha", "Laksh", "Male")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(familyTree.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}
func TestAddChildNegativeMotherMale(t *testing.T) {
	familyTree.PopulateFamilyTree()
	_, resultBool := familyTree.AddChild("Chit", "Laksh", "Male")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(familyTree.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}
func TestFindPersonPositive(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool := familyTree.IfPersonInFamily("Atya")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}

func TestFindPersonPositiveRoot(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool := familyTree.IfPersonInFamily("Anga")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}
func TestMaternalAunt(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindParentUncleAunt("Yodhan", helpers.GenderFemale, false)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestMaternalAuntNegativeNotFound(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindParentUncleAunt("Todhan", helpers.GenderFemale, false)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestPaternalAunt(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindParentUncleAunt("Dritha", helpers.GenderFemale, true)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestPaternalAuntNegativeNone(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindParentUncleAunt("Yodhan", helpers.GenderFemale, true)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestMaternalUncle(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindParentUncleAunt("Yodhan", helpers.GenderMale, false)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestPaternalUncle(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindParentUncleAunt("Dritha", helpers.GenderMale, true)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestMaternalAuntNegative(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindParentUncleAunt("Laki", helpers.GenderFemale, false)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestBrotherInLawPositive(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindInLaw("Chit", helpers.GenderMale)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestBrotherInLawPositiveSpouce(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindInLaw("Amba", helpers.GenderMale)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestBrotherInLawNegativeNone(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindInLaw("Vila", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestBrotherInLawNegativeNotFound(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindInLaw("Wila", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestSisterInLawPositive(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindInLaw("Chit", helpers.GenderFemale)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestSisterInLawNegativeNone(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindInLaw("Vila", helpers.GenderFemale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestSisterInLawNegativeNotFound(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindInLaw("Wila", helpers.GenderFemale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestFindChildrenPositive(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindChildren("Chit", helpers.GenderMale)
	if !resultBool {
		t.Errorf("Expected true, got %v and children %v", resultBool, len(*result))
		for _, value := range *result {
			t.Errorf(value.Name)
		}
	}
}

func TestFindChildrenNegativeNoChild(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindChildren("Yodhan", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and children %v", resultBool, len(*result))
		for _, value := range *result {
			t.Errorf(value.Name)
		}
	}
}

func TestFindChildrenNegativeNotFound(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindChildren("Todhan", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and children %v", resultBool, len(*result))
		for _, value := range *result {
			t.Errorf(value.Name)
		}
	}
}
func TestFindSiblingsPositive(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindSiblings("Vyas")
	if !resultBool {
		t.Errorf("Expected true, for %v and sibling count %v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestFindSiblingsNegativeNoSiblings(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindSiblings("Yodhan")
	if resultBool {
		t.Errorf("Expected true, for %v and sibling count %v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestFindSiblingsNegativeNotFound(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool, _ := familyTree.FindSiblings("Todhan")
	if resultBool {
		t.Errorf("Expected true, for %v and sibling count %v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestFindParentPositive(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool := familyTree.FindParent("Yodhan")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}
func TestFindParentNegativeRoot(t *testing.T) {
	familyTree.PopulateFamilyTree()
	_, responseBool := familyTree.FindParent("Anga")
	if responseBool {
		t.Errorf("Expected false, received %v", responseBool)
	}
}
func TestFindParentNegative(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool := familyTree.FindParent("Shan")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}

func TestFindParentNegativeNotFound(t *testing.T) {
	familyTree.PopulateFamilyTree()
	result, resultBool := familyTree.FindParent("Todhan")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}
