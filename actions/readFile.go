package actions

import (
	"fmt"
	"geektrust/helpers"
	"geektrust/models"
	"strings"
)

//Init function starts the program
func Init(path string) bool {
	fileExistBool := helpers.FileExists(path)
	if fileExistBool {
		//This means file exists
		allCommands := helpers.ReadFileFromPath(path)
		if len(allCommands) > 0 {
			//There are some commands
			var validCommands [][]string
			for _, value := range allCommands {
				isValidCommand, validCommand := isValidCommand(value)
				if isValidCommand {
					validCommands = append(validCommands, validCommand)
				}
			}
			if len(validCommands) > 0 {
				//We have some valid commands
				//Since we have some commands to run, its a good time to initialize the family
				var family models.Family
				family.Init()
				return RunCommands(validCommands, family)
			}
			fmt.Println("No valid commands found")
			return false
		}
		fmt.Println("No commands found in the file")
		return false
	}
	fmt.Println("File is missing, curropted or inappropriate")
	return false
}

//isValidCommand function returns true if the string is a valid command
//There can be 2 types of commands, ADD_CHILD or GET_RELATIONSHIP
//ADD_CHILD has Mother's Name, Child's Name and gender
//GET_RELATIONSHIP has name and relationship
func isValidCommand(input string) (bool, []string) {
	inputSplit := strings.Split(input, " ")
	//Will return false if there is an empty string
	for _, value := range inputSplit {
		if value == "" {
			return false, []string{}
		}
	}
	if len(inputSplit) == 4 && inputSplit[0] == helpers.AddChild && (inputSplit[3] == helpers.GenderFemale || inputSplit[3] == helpers.GenderMale) {
		return true, inputSplit
	}
	if len(inputSplit) == 3 && inputSplit[0] == helpers.GetRelationship {
		return true, inputSplit
	}
	return false, []string{}
}
