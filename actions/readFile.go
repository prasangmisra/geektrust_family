package actions

import (
	"geektrust/helpers"
	"geektrust/models"
	"strings"
)

//Init function starts the program
func Init(path string) bool {
	//This means file exists
	allCommands := helpers.ReadFileFromPath(path)
	if len(allCommands) > 0 {
		//There are some commands
		var validCommands [][]string
		for _, commandLine := range allCommands {
			isValidCommand, command := IsValidCommand(commandLine)
			if isValidCommand {
				validCommands = append(validCommands, command)
			}
		}
		if len(validCommands) > 0 {
			//We have some valid commands
			//Since we have some commands to run, its a good time to initialize the family
			var familyTree models.Family
			familyTree.PopulateFamilyTree()
			return RunCommands(validCommands, familyTree)
		}
		//fmt.Println("No valid commands found")
		return false
	}
	//fmt.Println("No commands found in the file")
	return false

}

//IsValidCommand function returns true if the string is a valid command
//There can be 2 types of commands, ADD_CHILD or GET_RELATIONSHIP
//ADD_CHILD has Mother's Name, Child's Name and gender
//GET_RELATIONSHIP has name and relationship
func IsValidCommand(command string) (bool, []string) {
	commandSplit := strings.Split(command, " ")
	//Will return false if there is an empty string
	for _, value := range commandSplit {
		if value == "" {
			return false, []string{}
		}
	}
	if len(commandSplit) == 4 && commandSplit[0] == helpers.AddChild && (commandSplit[3] == helpers.GenderFemale || commandSplit[3] == helpers.GenderMale) {
		return true, commandSplit
	}
	if len(commandSplit) == 3 && commandSplit[0] == helpers.GetRelationship {
		return true, commandSplit
	}
	return false, []string{}
}
