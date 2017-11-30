package main

import (
	"testing"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

type nuggetTest struct {
	Description string
	Expression string
}

var tests = []nuggetTest {
  { "SHA 1: hash", `filehashes = "file.dd" | extract  as ntfs | sha1
filehashes`} ,

 { "MD5: hash", `filehashes = "file.dd" | extract  as ntfs | md5
		filehashes`},

		{"Network: HTTP with subfield", `printx rawhttp
httpa = rawhttp | extract as http
print httpa.Host`}	, {"Memory: Process list", `mymemdump = "thememory.mddramimage" | extract as memory
mypslist = mymemdump | pslist
typex mypslist
printx mypslist
printx mypslist.Processes`}	, {"Union: Known hashes", `
KnownKittyHashes = "knownHashes.txt" | extract as listof-md5
Matched = MyHashes | union KnownKittyHashes
Matched
printx Matched `}, {"Filter and Sort", `myhashesfild = "file.dd" | extract  as ntfs | filter ctime > "01/01/09" | sort by ctime
myhashesfild`}}


func TestSyntax(t *testing.T) {
	log.Print("Testing syntax only..")

	for _, test := range tests {

		tree, nerr := GetTreeForInput(test.Expression)
		if nerr != nil {
			t.Error(
				"For", test.Description,
				"failed to process: ", test.Expression,
				"returning: ", tree.GetText())
		}
	}
	log.Print("Syntax test complete.")
}


func TestExecute (t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	log.Println("Testing syntax AND execution of input")
	for _, test := range tests {
		tree, nerr := GetTreeForInput(test.Expression)
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
		fmt.Println()
		if nerr != nil {
			t.Error(
				"For", test.Description,
				"failed to process: ", test.Expression,
				"returning: ", tree.GetText())
		}
	}
	log.Println("Exeuction testing complete")
}
