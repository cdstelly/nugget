package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"./NTypes"
	"./NActions"
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

func (s *TreeShapeListener) EnterAssign(ctx *parser.AssignContext) {
	//does the last execution on the right match the type on the left? do i even bother checking or let runtime go handle it?
	//1. look at all actions
	//2. does the rightmost action return what we want? if so, let's queue everything up
	//		a. from left to right, does the
	actions := ctx.AllNugget_action()
	varIdentifier := ctx.ID(0).GetText()
	//reverse the order of the actions
	for i := len(actions)/2-1; i >= 0; i-- {
		opp := len(actions)-1-i
		actions[i], actions[opp] = actions[opp], actions[i]
	}

	var builtActions []NActions.BaseAction
	for _,action := range actions {
		action_verb := action.GetStart().GetText()
		var theAction NActions.BaseAction
		switch action_verb {
		case "extract":
			//if it's an extract, it's coming from a variable or a raw load.
			//if it's a variable, get the info from it. otherwise, get it from the input.
			if len(ctx.AllID()) > 1 {
				// it's a variable
				fmt.Println("error: haven't implemented loading from a variable yet")
			} else {
				extractTarget := ctx.STRING().GetText()
				extractType := ctx.AsType().GetStop().GetText()
				fmt.Println("the extract info: ", extractTarget, " ", extractType)
			}
			fmt.Println("warning: using hardcoded target info for testing")
			extractTarget := "jo.extract"
			theAction = &NActions.ExtractNTFS{NTFSImageLocation:extractTarget}
		case "sha1":
			theAction = &NActions.SHA1Action{}
		case "md5":
			theAction = &NActions.MD5Action{}
		default:
			fmt.Println("action was not found: ", action_verb)  //parser should prevent us from getting here..
		}
		builtActions = append(builtActions, theAction)
	}

	for index, builtAction := range builtActions {
		if index+1 < len(builtActions) {
			//fmt.Println("action at index: ", index, "is ", builtAction, " and depends on: ", builtActions[index+1])
			var depAction NActions.BaseAction
			depAction = builtActions[index+1]
			builtAction.(NActions.BaseAction).SetDependency(depAction)
		} else {
			//fmt.Println("action at index: ", index, " is ", builtAction, " and has no dependency")
		}
		registers[varIdentifier] = builtActions[0]
	}
}

func (s *TreeShapeListener) EnterNugget_action(ctx *parser.Nugget_actionContext) {

}

func (this *TreeShapeListener) EnterSingleton_var(ctx *parser.Singleton_varContext) {
	identifier := ctx.ID().GetText()
	if v, ok := registers[identifier]; ok {
		fmt.Println(v)
		if a, ok := registers[identifier].(NActions.ExtractNTFS); ok {
			//we have a datatype ExztractNTFS
			fmt.Println("have extract ntfs actoin: ", a)
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
