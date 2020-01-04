package project

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"snareDrum/backend/interpreter"
	"strings"
)

// ValidateProject if the structure and language meta data are valid
func ValidateProject(path string) error {

	// check if header.json is existant
	if _, err := os.Stat(path + "/header.json"); os.IsNotExist(err) {
		return fmt.Errorf("Language meta data could not be found in this directory (%s/header.json)", path)
	}

	// check if main.sd is existant
	if _, err := os.Stat(path + "/main.sd"); os.IsNotExist(err) {
		return fmt.Errorf("Source code could not be found in this directory (%s/main.sd)", path)
	}

	if err := ValidateLang(path); err != nil {
		return err
	}

	return nil
}

// ValidateLang and check if header.json is valid
func ValidateLang(path string) error {
	// check if header.json is valid
	tempLang := interpreter.Lang{}

	addHeaderName(&path)

	// get lang file's contents
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(jsonData, &tempLang)
	if err != nil {
		fmt.Println(err)
		return err

	}

	// Check if all fields of tempLang are defined
	// I know it looks bad but there's no other way to check it :c
	if tempLang.Cell.Add == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "Cell.Add", path)
	} else if tempLang.Cell.Sub == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "Cell.Sub", path)
	} else if tempLang.IO.In == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "IO.In", path)
	} else if tempLang.IO.Out == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "IO.Out", path)
	} else if tempLang.Pointer.Up == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "Pointer.Up", path)
	} else if tempLang.Pointer.Down == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "Pointer.Down", path)
	} else if tempLang.Loop.Start == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "Loop.Start", path)
	} else if tempLang.Loop.End == "" {
		return fmt.Errorf("Token for %s wasn't defined correctly in %s/header.json", "Loop.End", path)
	}
	return nil
}

// GetSourceCode from the project directory
func GetSourceCode(path string) []byte {
	data, _ := ioutil.ReadFile(path + "/main.sd")
	return data
}

// GetLang uage meta data from the project directory
func GetLang(path string) interpreter.Lang {

	// check if header.json is valid
	tempLang := interpreter.Lang{}

	addHeaderName(&path)

	// get lang file's contents
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(jsonData, &tempLang)
	if err != nil {
		fmt.Println(err)
	}

	return tempLang
}

// GetName of the project
func GetName(path string) string {
	// trim path
	trimmed := strings.SplitAfter(path, "/")
	return trimmed[len(trimmed)-1]
}

// If header.json is not at the end of the path -> add
func addHeaderName(path *string) {
	// If path misses file specification -> add
	if trimmed := strings.SplitAfter(*path, "/"); trimmed[len(trimmed)-1] != "header.json" {
		a := "/header.json"
		*path += a
	}
}
