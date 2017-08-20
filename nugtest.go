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

	//log "github.com/sirupsen/logrus"
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

func (s *TreeShapeListener) EnterNugget_action(ctx *parser.Nugget_actionContext) {
	if ctx.Filter() != nil {
		fmt.Println(">>>", ctx.Filter().)
	}
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

//todo: see if i can break this up
func (s *TreeShapeListener) EnterAssign(ctx *parser.AssignContext) {
	varIdentifier := ctx.ID(0).GetText()

	//if no actions, then we do a simple calculation and assign it to a register
	if len(ctx.AllNugget_action()) == 0 {
		//if it's an extract string

		if ctx.AsType() != nil {
			extractTarget := ctx.STRING().GetText()
			extractType := ctx.AsType().GetStop().GetText()
			fmt.Println("a direct assignment has extract info: ", extractTarget, " ", extractType)
			registers[varIdentifier] = NTypes.TNTFS{PathToExtract: extractTarget,AsType:extractType}
		}
	} else {
		actions := ctx.AllNugget_action()




		//reverse the order of the actions
		for i := len(actions)/2 - 1; i >= 0; i-- {
			opp := len(actions) - 1 - i
			actions[i], actions[opp] = actions[opp], actions[i]
		}

		var builtActions []NActions.BaseAction
		for _, action := range actions {
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
				extractTarget := "C:\\Users\\drew\\School\\backups_before_upgrade\\jo.extract"
				extractData := "G:\\school\\image\\jo.ntfs"
				theAction = &NActions.ExtractNTFS{NTFSImageMetadataLocation:extractTarget, NTFSImageDataLocation:extractData}
			case "sha1":
				theAction = &NActions.SHA1Action{}
			case "md5":
				theAction = &NActions.MD5Action{}
			default:
				fmt.Println("action was not found: ", action_verb) //parser should prevent us from getting here..
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
				fmt.Println("action at index: ", index, " is ", builtAction, " and has no dependency. Setting dep to the var")
				if len(ctx.AllID()) > 1 {
					depVar := ctx.ID(1).GetText()
					// is it an existing var?
					if nVar, ok := registers[depVar]; ok {
						//if it's an action..
						if dep, ok := registers[depVar].(NActions.BaseAction); ok {
							//we have a datatype baseAction
							fmt.Println("the dependency for this action will be variable: ", nVar)
							builtAction.(NActions.BaseAction).SetDependency(dep)
						}
					} else {
						fmt.Println("Error: Var '", depVar, "' not recognized.")
					}
				} else if ctx.AsType() != nil { //is it an asType?

				} else { //was not recognized
					fmt.Println("Error: pattern not recognized.")
				}
			}
			registers[varIdentifier] = builtActions[0]
		}
	}
}

//resume here
// ok - problems arise. first, we need to figure out how to percolate up results from listeners. ie, filter was triggered, but how do i get the results of that up to teh nugget action?
//  online people have multiple listeners.. weird. why?? 
func getFilterForContext(ctx *parser.AssignContext) {
	if true {
		fmt.Println("no filter found")
	} else {
		fmt.Println("filter found")
	}
}

func (this *TreeShapeListener) EnterSingleton_var(ctx *parser.Singleton_varContext) {
	identifier := ctx.ID().GetText()
	if v, ok := registers[identifier]; ok {
//		fmt.Println("ident", ctx.ID().GetText(), " is: ", v)
		fmt.Printf("Variable %s has type: %T\n", identifier, v)
		if a, ok := registers[identifier].(*NActions.ExtractNTFS); ok {
			//we have a datatype ExztractNTFS
			fmt.Println("have extract ntfs action: ", a)
			a.Execute()
		}
		if a, ok := registers[identifier].(*NActions.SHA1Action); ok {
			//we have a datatype SHA1Extraction
			fmt.Println("have sha1 action: ", a)
			a.Execute()
		}
		if a, ok := registers[identifier].(*NActions.MD5Action); ok {
			//we have a datatype SHA1Extraction
			fmt.Println("have md5 action: ", a)
			a.Execute()
		}
	} else {
		fmt.Println("Error: var'", identifier, "' not found")
	}
}

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
