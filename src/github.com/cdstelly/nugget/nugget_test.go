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
  { "SHA 1", `filehashes = "file.dd" | extract  as ntfs | sha1
filehashes`} ,

 { "MD5 hash", `filehashes = "file.dd" | extract  as ntfs | md5
		filehashes`} }


func TestSyntax(t *testing.T) {
	log.Print("Testing syntax only")

	for _, test := range tests {

		tree, nerr := GetTreeForInput(test.Expression)
		if nerr != nil {
			t.Error(
				"For", test.Description,
				"failed to process: ", test.Expression,
				"returning: ", tree.GetText())
		}
	}
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
}
