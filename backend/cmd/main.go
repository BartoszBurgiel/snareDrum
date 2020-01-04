package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"snareDrum/backend/compiler"
	"snareDrum/backend/generator"
	"snareDrum/backend/interpreter"
	"snareDrum/backend/project"
	"strings"
)

func main() {

	fmt.Println("ARGS", os.Args)

	// Get args
	args := os.Args[1:]

	fmt.Println("args", args)
	// Which action shall be performed
	action := args[0]

	switch action {
	case "build":

		path := argExists(args, 1)

		stack := buildStackFromProject(path)

		// Compile to binary
		bin := compiler.Compile(stack)

		// Write off
		writeOff(path, bin)
		break
	case "debug":
		path := argExists(args, 1)
		stack := buildStackFromProject(path)
		stack.Execute()

		fmt.Println(stack.Debug())
		break
	case "exec":
		path := argExists(args, 1)

		// Get the type of the file
		fileType := getType(path)

		// check if executable or project
		if fileType == "sdexe" {
			// Get source code
			code, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
			}
			// Execute file
			output := compiler.Execute(code)
			fmt.Println(output)
			break
		} else {

			stack := buildStackFromProject(path)

			output, _ := stack.Execute()
			fmt.Println(output)
			break
		}

	case "generate":
		path := argExists(args, 1)

		// Check Lang
		if err := project.ValidateLang(path); err != nil {
			fmt.Println(err)
			break
		}

		// Get lang
		lang := project.GetLang(path)

		// Get text to translate
		text := argExists(args, 2)

		// Generate code
		code := generator.Generate(lang, text)

		err := ioutil.WriteFile("GEN.sd", []byte(code), 0644)
		if err != nil {
			fmt.Println(err)
		}
		break
	default:
		fmt.Printf("Unknown argument %s", action)
	}

	// Terminate program after all actions
	os.Exit(0)

}

/*
ROUTINES
*/
func notEnoughArguments() {
	fmt.Println("Insufficient number of arguments provided to")

	// Terminate Program
	os.Exit(0)
}

// check if given argument exists if true -> return the argument
func argExists(args []string, index int) string {
	if index > len(args) {
		notEnoughArguments()
	} else {
		return args[index]
	}
	return ""
}

// Get all data from the files and bild a stack
func buildStackFromProject(path string) interpreter.Stack {

	// If project is not valid
	if err := project.ValidateProject(path); err != nil {
		fmt.Println(err)

		os.Exit(0)
	}

	// Get programm
	lang := project.GetLang(path)
	code := project.GetSourceCode(path)

	// Build the stack
	stack := interpreter.Stack{}
	stack.New()
	stack.Build(string(code), lang)

	return stack
}

// write off the given contents into a file
func writeOff(path string, content []byte) {
	err := ioutil.WriteFile(project.GetName(path)+".sdexe", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// get type of a given file name
func getType(file string) string {
	trimmed := strings.SplitAfter(file, ".")

	if len(trimmed) > 1 {
		return trimmed[1]
	}

	return ""
}
