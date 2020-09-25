package actions

import (
	"fmt"
	"geektrust/helpers"
	"geektrust/models"
)

//RunCommands function will run all the valid commands from the file
func RunCommands(input [][]string, family models.Family) bool {
	for _, value := range input {
		switch value[0] {
		case helpers.AddChild:
			fmt.Println(AddChildToFamily(family, value[1], value[2], value[3]))
			break
		case helpers.GetRelationship:
			fmt.Println(GetRelationshipOfFamily(family, value[1], value[2]))
			break
		default:
			fmt.Println("")
		}
	}
	return true
}

//AddChildToFamily is the business logic implementation of the command
func AddChildToFamily(family models.Family, mother string, child string, gender string) string {
	responseString, _ := family.AddChild(mother, child, gender)
	return responseString
}

//GetRelationshipOfFamily function is the business logic of the implementation of this command
func GetRelationshipOfFamily(family models.Family, name string, relationship string) string {
	//First of all check if member exists in family
	_, existBool := family.IfPersonInFamily(name)
	if !existBool {
		return helpers.ResponseNotFound
	}
	switch relationship {
	case helpers.RelationshipSiblings:
		return SiblingsCommand(family, name)
	case helpers.RelationshipDaughter:
		return ChildCommand(family, name, helpers.GenderFemale)
	case helpers.RelationshipSon:
		return ChildCommand(family, name, helpers.GenderMale)
	case helpers.RelationshipBrotherInLaw:
		return InLawCommand(family, name, helpers.GenderMale)
	case helpers.RelationshipSisterInLaw:
		return InLawCommand(family, name, helpers.GenderFemale)
	case helpers.RelationshipMaternalAunt:
		return UncleAuntCommand(family, name, helpers.GenderFemale, false)
	case helpers.RelationshipPaternalAunt:
		return UncleAuntCommand(family, name, helpers.GenderFemale, true)
	case helpers.RelationshipMaternalUncle:
		return UncleAuntCommand(family, name, helpers.GenderMale, false)
	case helpers.RelationshipPaternalUncle:
		return UncleAuntCommand(family, name, helpers.GenderMale, true)
	default:
		return ""
	}
}

//UncleAuntCommand function
func UncleAuntCommand(family models.Family, name string, gender string, isPaternal bool) string {
	uncleAunts, uncleAuntBool, uncleAuntsResponseString := family.FindParentUncleAunt(name, gender, isPaternal)
	if !uncleAuntBool {
		return uncleAuntsResponseString
	}
	responseString := ""
	for key, value := range *uncleAunts {
		if key == 0 {
			responseString += value.Name
		} else {
			responseString += " " + value.Name
		}

	}
	return responseString
}

//InLawCommand function
func InLawCommand(family models.Family, name string, gender string) string {
	inLaw, inLawBool, inLawResponseString := family.FindInLaw(name, gender)
	if !inLawBool {
		return inLawResponseString
	}
	responseString := ""
	for key, value := range *inLaw {
		if key == 0 {
			responseString += value.Name
		} else {
			responseString += " " + value.Name
		}
	}
	return responseString
}

//ChildCommand function
func ChildCommand(family models.Family, name string, gender string) string {
	children, childrenBool, childResponseString := family.FindChildren(name, gender)
	if !childrenBool {
		return childResponseString
	}
	responseString := ""
	for key, value := range *children {
		if key == 0 {
			responseString += value.Name
		} else {
			responseString += " " + value.Name
		}
	}
	return responseString
}

//SiblingsCommand function
func SiblingsCommand(family models.Family, name string) string {
	siblings, siblingBool, siblingResponse := family.FindSiblings(name)
	if !siblingBool {
		return siblingResponse
	}
	responseString := ""
	for key, value := range *siblings {
		if key == 0 {
			responseString += value.Name
		} else {
			responseString += " " + value.Name
		}
	}
	return responseString
}
