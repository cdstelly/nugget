package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"./NTypes"
	"./nug2"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var (
	pathToInput string
	registers map[string]interface{}
)

func init() {
	flag.StringVar(&pathToInput, "input", "input2.nug", "Path to input")
	flag.Parse()
	registers = make(map[string]interface{})
}

type TreeShapeListener struct {
	*parser.BaseNugget2Listener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//		fmt.Println("Entered a rule, result: " + ctx.GetText())
}

func (this *TreeShapeListener) EnterDefine(ctx *parser.DefineContext) {
	isList := ctx.LISTOP() != nil
	identifier := ctx.ID().GetText()
	nugget_type := ctx.Nugget_type().GetText()

	fmt.Println("found a define: ", ctx.ID(), " ", ctx.Nugget_type().GetText(), " is a list?: ", isList)

	if _, exists := registers[identifier]; exists {
		fmt.Println("the variable ", identifier, " already exists")
	} else {
		switch nugget_type {
		case "ntfs":
			if isList {
				registers[identifier] = []NTypes.TNTFS{}
			} else {
				registers[identifier] = NTypes.TNTFS{}
			}
		case "file":
			if isList {
				registers[identifier] = []NTypes.FileInfo{}
			} else {
				registers[identifier] = NTypes.FileInfo{}
			}
		case "sha1":
			if isList {
				registers[identifier] = []NTypes.SHA1{}
			} else {
				registers[identifier] = NTypes.SHA1{}
			}
		case "md5"://TODO: investigate impact of 'natural types' such as string - see if we should wrap in an NType
			if isList {
				registers[identifier] = []string{}
			} else {
				registers[identifier] = ""
			}
		case "string":
			if isList {
				registers[identifier] = []NTypes.NString{}
			} else {
				registers[identifier] = NTypes.NString{}
			}
		case "packet":
			if isList {
				registers[identifier] = []NTypes.NPacket{}
			} else {
				registers[identifier] = NTypes.NPacket{}
			}
		default:
			fmt.Println("Was not able to build type: ", nugget_type)
		}
	}
}

/*
func (this *TreeShapeListener) EnterAssign(ctx *parser.AssignContext) {
	if v, ok := registers[ctx.StrLit().GetText()]; ok {
		fmt.Println(v)
	}
	if v, ok := sources[ctx.StrLit().GetText()]; ok {
		fmt.Println(v.Name)
	}
}

func (this *TreeShapeListener) EnterInitextract(ctx *parser.InitextractContext) {
	fmt.Println("extrating filesystem..")
	newsource := source{Name: ctx.Target().GetText()}
	sources[ctx.StrLit().GetText()] = newsource
}

func (this *TreeShapeListener) EnterAssign(ctx *parser.AssignContext) {
	//fmt.Println("assigning to variable: ", ctx.StrLit(0).GetText(), " the value : ", ctx.StrLit(1).GetText())
	registers[ctx.StrLit(0).GetText()] = ctx.StrLit(1).GetText()
}
*/

func main() {
	file, err := os.Open(pathToInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := antlr.NewInputStream(scanner.Text())
		lexer := parser.NewNugget2Lexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewNugget2Parser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true
		tree := p.Prog()
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
