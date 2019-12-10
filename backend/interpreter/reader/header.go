package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	interpreter "snareDrum/backend/interpreter/core"
)

/*
Read the header file which defines the language tokens
and syntax
*/

// ReadHeader reads given json data and converts it into lang struct
func ReadHeader(jsonData []byte) interpreter.Lang {
	out := &interpreter.Lang{}

	err := json.Unmarshal(jsonData, out)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return *out
}

// ReadHeaderFile reads header.js file and returns Lang struct
func ReadHeaderFile(dir string) interpreter.Lang {
	// Open file
	jsonData, err := ioutil.ReadFile(dir + "header.js")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return ReadHeader(jsonData)
}
