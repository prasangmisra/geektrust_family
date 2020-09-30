package actions

import (
	"fmt"

	"geektrust/helpers"
	"geektrust/models"
)

//RunCommands function will run all the valid commands from the file
func RunCommands(commands [][]string, familyTree models.Family) bool {
	for _, command := range commands {
		switch command[0] {
		case helpers.AddChild:
			fmt.Println(AddChildToFamily(familyTree, command[1], command[2], command[3]))
			break
		case helpers.GetRelationship:
			fmt.Println(GetRelationshipOfFamily(familyTree, command[1], command[2]))
			break
		default:
			fmt.Println("")
		}
	}
	return true
}

//AddChildToFamily is the business logic implementation of the command
func AddChildToFamily(familyTree models.Family, motherOfNewChild string, newChildName string, newChildGender string) string {
	responseString, _ := familyTree.AddChild(motherOfNewChild, newChildName, newChildGender)
	return responseString
}

//GetRelationshipOfFamily function is the business logic of the implementation of this command
func GetRelationshipOfFamily(familyTree models.Family, familyMemberName string, familyMemberRelationship string) string {
	//First of all check if member exists in family
	_, existBool := familyTree.IfPersonInFamily(familyMemberName)
	if !existBool {
		return helpers.ResponseNotFound
	}
	switch familyMemberRelationship {
	case helpers.RelationshipSiblings:
		return SiblingsCommand(familyTree, familyMemberName)
	case helpers.RelationshipDaughter:
		return ChildCommand(familyTree, familyMemberName, helpers.GenderFemale)
	case helpers.RelationshipSon:
		return ChildCommand(familyTree, familyMemberName, helpers.GenderMale)
	case helpers.RelationshipBrotherInLaw:
		return InLawCommand(familyTree, familyMemberName, helpers.GenderMale)
	case helpers.RelationshipSisterInLaw:
		return InLawCommand(familyTree, familyMemberName, helpers.GenderFemale)
	case helpers.RelationshipMaternalAunt:
		return UncleAuntCommand(familyTree, familyMemberName, helpers.GenderFemale, false)
	case helpers.RelationshipPaternalAunt:
		return UncleAuntCommand(familyTree, familyMemberName, helpers.GenderFemale, true)
	case helpers.RelationshipMaternalUncle:
		return UncleAuntCommand(familyTree, familyMemberName, helpers.GenderMale, false)
	case helpers.RelationshipPaternalUncle:
		return UncleAuntCommand(familyTree, familyMemberName, helpers.GenderMale, true)
	default:
		return ""
	}
}

//UncleAuntCommand function is the command to find the Maternal/Paternal Uncle/Aunt if present
func UncleAuntCommand(familyTree models.Family, familyMemberName string, uncleAuntGender string, isPaternal bool) string {
	uncleAunts, uncleAuntBool, uncleAuntsResponseString := familyTree.FindParentUncleAunt(familyMemberName, uncleAuntGender, isPaternal)
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

//InLawCommand function is the command to find the brother/sister in law of the family member
func InLawCommand(familyTree models.Family, familyMemberName string, inLawGender string) string {
	inLaw, inLawBool, inLawResponseString := familyTree.FindInLaw(familyMemberName, inLawGender)
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

//ChildCommand function returns the names of the children of a family member if present
//childGender is Male for son and Female for daughter
func ChildCommand(familyTree models.Family, familyMemberName string, childGender string) string {
	allChildren, hasChildren, hasChidrenResponse := familyTree.FindChildren(familyMemberName, childGender)
	if !hasChildren {
		return hasChidrenResponse
	}
	listOfChidren := ""
	for key, child := range *allChildren {
		if key == 0 {
			listOfChidren += child.Name
		} else {
			listOfChidren += " " + child.Name
		}
	}
	return listOfChidren
}

//SiblingsCommand function returns the list of all the siblings of a family member
func SiblingsCommand(familyTree models.Family, familyMemberName string) string {
	allSiblings, hasSiblings, hasSiblingResponse := familyTree.FindSiblings(familyMemberName)
	if !hasSiblings {
		return hasSiblingResponse
	}
	listOfSiblings := ""
	for key, sibling := range *allSiblings {
		if key == 0 {
			listOfSiblings += sibling.Name
		} else {
			listOfSiblings += " " + sibling.Name
		}
	}
	return listOfSiblings
}
