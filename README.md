This document explains how to run the solution in Go. 

# Link to the problem statement

https://www.geektrust.in/coding-problem/backend/family


# Assumption in the problem statement

It is said in the problem statement that the names will be unique.
In this case, if we try to add a child in the family whose name already exists, there will be a conflict.
To avoid this, I have added a check, which stops the code from adding such a child.

# Handling Dependencies

The source code has only 1 dependency which can be found in the go.mod file, i.e. Go v1.12


# Running The application

Use the command `go run main.go <path to the input file>`

# Building The application

Use the command `go build .` to see the executable in the same folder.
Then run `./geektrust <path to the input file>`

# Package & Directory structure of your Go application

The main package is `geektrust`. The project structure is described below:


```
bin/
    geektrust                           # command executable
src/
    geektrust                           # main package
        main.go                         # start program file
        input.txt                       # sample input file to run the code
        blankInput.txt                  # a blank text file used as an argument for testing
        go.mod                          # go modules file
        actions                         # package for business logic
            commands.go                 # file to run the individual commands of the input
            commands_test.go            # test file for the commands
            readFile.go                 # file to locate, evaluate the file and extract the contents
            readFile_test.go            # test file for the file access functions
        helpers                         # package for constants and helper controls
            constants.go                # collection of all the constants
            fileHelper_test.go          # test file for file access functions
            fileHelper.go               # functions to help access the input file
        models                          # package with the models used
            family_test.go              # test file for family tree based models
            family.go                   # model with structs and functions for family tree
            person_test.go              # test file for person based models
            person.go                   # model with structs and functions for a person
        Readme.md                       # the file you are currently reading
```



I have initialised the family tree not on program start but after the following conditions are met:
1. The input file exists
2. The input file is not empty
3. The input file has at least 1 valid command, i.e. ADD_CHILD or GET_RELATIONSHIP

I have done this on purpose to save unnecessary memory storage.
