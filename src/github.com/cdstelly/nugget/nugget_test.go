package main

import (
	"testing"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

func TestParseMD5Input(t *testing.T) {
	log.Print("Testing MD5 input parsing")
	input :=
		`
		filehashes = "file.dd" | extract  as ntfs | md5
		filehashes
		`
	_, nerr := GetTreeForInput(input)
	if nerr != nil {
		t.Fail()
	}
}


func TestExecuteMD5Input (t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	log.Print("Testing MD5 input execution")
	input :=
		`
		filehashes = "file.dd" | extract  as ntfs | md5
		filehashes
		`
	tree, nerr := GetTreeForInput(input)
	if nerr != nil {
		t.Fail()
	}
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	fmt.Println()
}
