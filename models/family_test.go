package models

import (
	"geektrust/helpers"
	"testing"
)

var family Family

func TestStruct(t *testing.T) {
	var family Family
	if family.Person != nil {
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
	family.Person = &queen
	queen.Children = append(queen.Children, &child)
	_, resultBool := family.AddChild("Anga", "Disha1", "Female")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(family.Person.Children))
	}
}

func TestInit(t *testing.T) {
	family.Init()
	if len(family.Person.Children) != 5 {
		t.Errorf("Expected 5, got %v", len(family.Person.Children))
	}
}

func TestAddChildAfterInit(t *testing.T) {
	family.Init()
	_, resultBool := family.AddChild("Krithi", "Disha", "Female")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(family.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}

func TestAddChildNegativeDuplicate(t *testing.T) {
	family.Init()
	_, resultBool := family.AddChild("Krithi", "Krithi", "Female")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(family.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}
func TestAddChildNegativeNoMother(t *testing.T) {
	family.Init()
	_, resultBool := family.AddChild("Disha", "Laksh", "Male")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(family.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}
func TestAddChildNegativeMotherMale(t *testing.T) {
	family.Init()
	_, resultBool := family.AddChild("Chit", "Laksh", "Male")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, len(family.Person.FindChildrenWithName("Satya").FindChildrenWithName("Vyas").FindChildrenWithName("Krithi").Children))
	}
}
func TestFindPersonPositive(t *testing.T) {
	family.Init()
	result, resultBool := family.IfPersonInFamily("Atya")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}

func TestFindPersonPositiveRoot(t *testing.T) {
	family.Init()
	result, resultBool := family.IfPersonInFamily("Anga")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}
func TestMaternalAunt(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindParentUncleAunt("Yodhan", helpers.GenderFemale, false)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestMaternalAuntNegativeNotFound(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindParentUncleAunt("Todhan", helpers.GenderFemale, false)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestPaternalAunt(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindParentUncleAunt("Dritha", helpers.GenderFemale, true)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestPaternalAuntNegativeNone(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindParentUncleAunt("Yodhan", helpers.GenderFemale, true)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestMaternalUncle(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindParentUncleAunt("Yodhan", helpers.GenderMale, false)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestPaternalUncle(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindParentUncleAunt("Dritha", helpers.GenderMale, true)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestMaternalAuntNegative(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindParentUncleAunt("Laki", helpers.GenderFemale, false)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestBrotherInLawPositive(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindInLaw("Chit", helpers.GenderMale)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestBrotherInLawPositiveSpouce(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindInLaw("Amba", helpers.GenderMale)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestBrotherInLawNegativeNone(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindInLaw("Vila", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestBrotherInLawNegativeNotFound(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindInLaw("Wila", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestSisterInLawPositive(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindInLaw("Chit", helpers.GenderFemale)
	if !resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestSisterInLawNegativeNone(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindInLaw("Vila", helpers.GenderFemale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestSisterInLawNegativeNotFound(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindInLaw("Wila", helpers.GenderFemale)
	if resultBool {
		t.Errorf("Expected true, got %v and number:%v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestFindChildrenPositive(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindChildren("Chit", helpers.GenderMale)
	if !resultBool {
		t.Errorf("Expected true, got %v and children %v", resultBool, len(*result))
		for _, value := range *result {
			t.Errorf(value.Name)
		}
	}
}

func TestFindChildrenNegativeNoChild(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindChildren("Yodhan", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and children %v", resultBool, len(*result))
		for _, value := range *result {
			t.Errorf(value.Name)
		}
	}
}

func TestFindChildrenNegativeNotFound(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindChildren("Todhan", helpers.GenderMale)
	if resultBool {
		t.Errorf("Expected true, got %v and children %v", resultBool, len(*result))
		for _, value := range *result {
			t.Errorf(value.Name)
		}
	}
}
func TestFindSiblingsPositive(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindSiblings("Vyas")
	if !resultBool {
		t.Errorf("Expected true, for %v and sibling count %v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestFindSiblingsNegativeNoSiblings(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindSiblings("Yodhan")
	if resultBool {
		t.Errorf("Expected true, for %v and sibling count %v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}

func TestFindSiblingsNegativeNotFound(t *testing.T) {
	family.Init()
	result, resultBool, _ := family.FindSiblings("Todhan")
	if resultBool {
		t.Errorf("Expected true, for %v and sibling count %v", resultBool, len(*result))
		for _, value := range *result {
			t.Error(value.Name)
		}
	}
}
func TestFindParentPositive(t *testing.T) {
	family.Init()
	result, resultBool := family.FindParent("Yodhan")
	if !resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}
func TestFindParentNegativeRoot(t *testing.T) {
	family.Init()
	_, responseBool := family.FindParent("Anga")
	if responseBool {
		t.Errorf("Expected false, received %v", responseBool)
	}
}
func TestFindParentNegative(t *testing.T) {
	family.Init()
	result, resultBool := family.FindParent("Shan")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}

func TestFindParentNegativeNotFound(t *testing.T) {
	family.Init()
	result, resultBool := family.FindParent("Todhan")
	if resultBool {
		t.Errorf("Expected true, got %v with name %v", resultBool, result.Name)
	}
}
