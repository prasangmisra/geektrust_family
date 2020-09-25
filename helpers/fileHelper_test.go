package helpers

import "testing"

func TestFileExistsPositive(t *testing.T) {
	result := FileExists("/Users/prasangmisra/go/src/geektrust/input.txt")
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestFileExistsNegativePackage(t *testing.T) {
	result := FileExists("/Users/prasangmisra/go/src/geektrust")
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
func TestFileExistsNegativeWrongFile(t *testing.T) {
	result := FileExists("/Users/prasangmisra/go/src/geektrust/input.tx")
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestReadFileFromPathPositive(t *testing.T) {
	result := ReadFileFromPath("/Users/prasangmisra/go/src/geektrust/input.txt")
	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %v", len(result))
	}
}

func TestReadFileFromPathNegativeNoFile(t *testing.T) {
	result := ReadFileFromPath("/Users/prasangmisra/go/src/geektrust/inputLong.txt")
	if len(result) != 0 {
		t.Errorf("Expected 3 lines, got %v", len(result))
	}
}
