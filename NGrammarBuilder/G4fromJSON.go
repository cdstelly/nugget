package main

import (
	"io/ioutil"
	"os"
	"log"
	"strings"
	"encoding/json"
	"fmt"
)

//this is unweildy and needs to improve
//goal is to standardize the way we add new tools or functions to the nugget specification
//this will be really neat since we're using a grammar based appraoch to this DSL
// idea will be to fill out a json file, run this file, and have it appropriatly populate the G4 (antlr grammer) file
//reqs: ANTLR4 up and running..

type NuggetJSONEntry struct {
	Name string `json:"name"`
	CodePath string `json:"codepath"`    //relative to the base nugget folder

}

var (
	G4Path string
	OutPath string
	JSONDir string
	NJSONEntries []NuggetJSONEntry
)

func init() {
	G4Path = "../Nugget2.g4"
	OutPath = "../nug2/"
	JSONDir = "NGrammarBuilder"
}

func main() {
	files, err := ioutil.ReadDir(JSONDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(),"json") {
			NJSONEntries = append(NJSONEntries,readJSONFile(JSONDir + string(os.PathSeparator) + f.Name()) )
		}
	}
	for _,nje := range NJSONEntries {
		fmt.Println(nje)
	}
}

func readJSONFile(fp string) NuggetJSONEntry {
	raw, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var c NuggetJSONEntry
	json.Unmarshal(raw, &c)
	return c
}
