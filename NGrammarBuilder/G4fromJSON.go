package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"log"
	"strings"
	"encoding/json"
	"fmt"
)

//this is quite unweildy and needs to improve
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
	AntlrPath string
)

func init() {
	G4Path = "Nugget.g4"
	OutPath = "nug"
	JSONDir = "NGrammarBuilder"
	AntlrPath = "C:\\Users\\drew\\.Gogland1.0\\config\\plugins\\intellij-plugin-v4\\lib\\antlr-4.7-intellij-1.8.4-version-complete.jar"
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
	// build the text we're going to insert into the .g4 file
	var stringToEnter string
	for _,nje := range NJSONEntries {
		stringToEnter = stringToEnter + "'" + nje.Name + "'" + "|" + "\r\n\t"
	}
	//stringToEnter = strings.TrimSuffix(stringToEnter, "|")
	stringToEnter += "'NUGGETGENERATORPLACEHOLDER'"
	//fmt.Println("ste",stringToEnter)
	// insert the text into the g4 file
	rewriteG4(stringToEnter)
	// execute an antlr build routine
	executeAntlrBuild()
}

func executeAntlrBuild() {
	cmd := exec.Command("java", "-jar", AntlrPath, "-o", OutPath, "-Dlanguage=Go", G4Path)
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	} else {
		fmt.Println("Built Nugget")
	}
}

func rewriteG4(replaceWith string) {
	from := "'NUGGETGENERATORPLACEHOLDER'"
	read, err := ioutil.ReadFile(G4Path)
	if err != nil {
		log.Fatal(err)
	}
	newContents := strings.Replace(string(read), from, replaceWith, -1)
	err = ioutil.WriteFile(G4Path, []byte(newContents), 0)
	if err != nil {
		log.Fatal(err)
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
