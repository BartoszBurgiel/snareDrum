package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"snareDrum/backend/compiler"
	"snareDrum/backend/generator"
	"snareDrum/backend/interpreter"
	"snareDrum/backend/project"
	"snareDrum/backend/ui"
	"strings"
	"time"
)

func main() {

	// Get args
	args := os.Args[1:]

	// Which action shall be performed
	action := args[0]

	switch action {
	case "build":
		path := argExists(args, 1)

		// Build stack
		stack := buildStackFromProject(path)

		// Compile and write off
		progress := 0
		go ui.ProgressBar(&progress, len(stack.Register.Methods), "Compiling")
		compiler.CompileAndWrite(stack, path, &progress)

		// Sleep and give let the progress bar finnish
		time.Sleep(1000)

		break
	case "debug":
		path := argExists(args, 1)
		stack := buildStackFromProject(path)
		stack.Execute()

		fmt.Println("\n", stack.Debug())
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
			fmt.Println("\n", output.String())
			output.Reset()
			break
		} else {

			stack := buildStackFromProject(path)
			output, _ := stack.Execute()
			fmt.Println("\n", output.String())
			output.Reset()
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

		// Set off generator and the goroutine and track progress
		progress := 0
		go ui.ProgressBar(&progress, len(text), "Generating")
		code := generator.Generate(lang, text, &progress)

		// Write to file
		err := ioutil.WriteFile("GEN.sd", code.Bytes(), 0644)
		if err != nil {
			fmt.Println(err)
		}
		break
	case "translate-file":
		path := argExists(args, 1)

		// Check Lang
		if err := project.ValidateLang(path); err != nil {
			fmt.Println(err)
			break
		}

		// Get lang
		lang := project.GetLang(path)

		// Get text to translate
		pathToText := argExists(args, 2)

		// Get content
		content, err := ioutil.ReadFile(pathToText)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Generate code
		progress := 0
		go ui.ProgressBar(&progress, len(content), "Translating")
		generator.GenerateFile(lang, content, &progress)
		break

	default:
		fmt.Printf("Unknown argument '%s'", action)
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

	progress := 0
	go ui.ProgressBar(&progress, len(code), "Building")
	stack.Build(lang, string(code), &progress)

	// Sleep and give the progress bar time to complete
	time.Sleep(1000)

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
