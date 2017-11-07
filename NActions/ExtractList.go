package NActions

import (
	"../NTypes"

	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type ExtractList struct {
	executed  bool
	dependsOn BaseAction
	filters   []NTypes.Filter

	ListLocation string
	ListType     string
	ListContent  []string
}

func (na *ExtractList) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractList) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractList) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (na *ExtractList) loadFromHTTP(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, lin := range lines {
		switch na.ListType {
		case "md5":
			if isValidMD5(lin) {
				na.ListContent = append(na.ListContent, lin)
			}
		case "sha1":
			if isValidSHA1(lin) {
				na.ListContent = append(na.ListContent, lin)
			}
		case "sha256":
			if isValidSHA256(lin) {
				na.ListContent = append(na.ListContent, lin)
			}
		}
	}
}

func (na *ExtractList) Execute() {

	typeOfLocation := strings.Split(na.ListLocation, ":")[0]
	if typeOfLocation == "file" {

	} else if typeOfLocation == "url" {

	}
	theLines, err := readLines(na.ListLocation)
	if err != nil {
		fmt.Println("Error!", err)
	}

	for _, lin := range theLines {
		switch na.ListType {
		case "md5":
			if isValidMD5(lin) {
				na.ListContent = append(na.ListContent, lin)
			}
		case "sha1":
			if isValidSHA1(lin) {
				na.ListContent = append(na.ListContent, lin)
			}
		case "sha256":
			if isValidSHA256(lin) {
				na.ListContent = append(na.ListContent, lin)
			}
		}
	}
	na.executed = true
}

func isValidMD5(in string) bool {
	regex, err := regexp.Compile("[a-fA-F0-9]{32}")
	if err != nil {
		fmt.Println(err)
	}
	return regex.MatchString(in)
}

func isValidSHA1(in string) bool {
	regex, err := regexp.Compile("[a-fA-F0-9]{40}")
	if err != nil {
		fmt.Println(err)
	}
	return regex.MatchString(in)
}

func isValidSHA256(in string) bool {
	regex, err := regexp.Compile("[a-fA-F0-9]{64}")
	if err != nil {
		fmt.Println(err)
	}
	return regex.MatchString(in)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (na *ExtractList) GetResults() interface{} {
	if na.executed == false {
		na.Execute()
	}
	return na.ListContent
}

func (na *ExtractList) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	//na.executed = false
	na.filters = filters
}
