package actions

import "testing"

func TestValidCommandGetRelationShipPositive(t *testing.T) {
	inputString := "GET_RELATIONSHIP Satvy Sister-In-Law"
	resultBool, _ := isValidCommand(inputString)
	if !resultBool {
		t.Errorf("Expected true, received %v", resultBool)
	}

}

func TestValidCommandGetRelationShipNegativeEmptyString(t *testing.T) {
	inputString := "GET_RELATIONSHIP Satvy "
	resultBool, result := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected true, received %v", len(result))
	}
}

func TestValidCommandGetRelationShipNegativeLessParams(t *testing.T) {
	inputString := "GET_RELATIONSHIP Satvy"
	resultBool, result := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected true, received %v", len(result))
	}
}

func TestValidCommandGetRelationShipNegativeMoreParams(t *testing.T) {
	inputString := "GET_RELATIONSHIP Satvy Sister In Law"
	resultBool, result := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected true, received %v", len(result))
	}
}

func TestValidCommandGetRelationShipNegativeWrongCommandName(t *testing.T) {
	inputString := "GET_RELATIONSHIPS Satvy Sister-In-Law"
	resultBool, result := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected true, received %v", len(result))
	}
}

func TestValidCommandAddPositive(t *testing.T) {
	inputString := "ADD_CHILD Satvy Disha Female"
	resultBool, _ := isValidCommand(inputString)
	if !resultBool {
		t.Errorf("Expected true, received %v", resultBool)
	}

}

func TestValidCommandAddNegativeLessParams(t *testing.T) {
	inputString := "ADD_CHILD Satvy Disha "
	resultBool, _ := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected false, received %v", resultBool)
	}

}

func TestValidCommandAddNegativeMoreParams(t *testing.T) {
	inputString := "ADD_CHILD Satvy Disha Female Female"
	resultBool, _ := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected false, received %v", resultBool)
	}
}

func TestValidCommandAddNegativeWrongGender(t *testing.T) {
	inputString := "ADD_CHILD Satvy Disha Females"
	resultBool, _ := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected false, received %v", resultBool)
	}
}

func TestValidCommandAddNegativeWrongCommandName(t *testing.T) {
	inputString := "ADD_CHILDREN Satvy Disha Females"
	resultBool, _ := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected false, received %v", resultBool)
	}
}

func TestValidCommandNegativeWrongCommandName(t *testing.T) {
	inputString := "ADD_GET Satvy Disha Females"
	resultBool, _ := isValidCommand(inputString)
	if resultBool {
		t.Errorf("Expected false, received %v", resultBool)
	}
}

func TestInitNoFile(t *testing.T) {
	result := Init("input.txt")
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
func TestInitSuccess(t *testing.T) {
	result := Init("/Users/prasangmisra/go/src/geektrust/input.txt")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}

func TestInitFailNoValidCommands(t *testing.T) {
	result := Init("/Users/prasangmisra/go/src/geektrust/go.mod")
	if result {
		t.Errorf("Expected true, received %v", result)
	}
}

func TestInitFailNoCommands(t *testing.T) {
	result := Init("/Users/prasangmisra/go/src/geektrust/blankInput.txt")
	if result {
		t.Errorf("Expected true, received %v", result)
	}
}
