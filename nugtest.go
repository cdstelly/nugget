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
	//"github.com/golang-collections/collections/stack"
	"reflect"
)

var (
	pathToInput string
	registers map[string]interface{}

	nodeMap map[antlr.ParseTree]interface{}
)

func init() {
	flag.StringVar(&pathToInput, "input", "input2.nug", "Path to input")
	flag.Parse()
	registers = make(map[string]interface{})

	//nodemap can be used to share data across constructs, just store it here whenever and retrieve based on ctx
	nodeMap = make(map[antlr.ParseTree]interface{})
}

func setValue(ctx antlr.ParseTree, value interface{}) {
	nodeMap[ctx] = value
}

func getValue(n antlr.ParseTree) interface{} {
	return nodeMap[n]
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

func (s *TreeShapeListener) ExitNugget_action(ctx *parser.Nugget_actionContext) {
	//grab filters if the node below us was a filter
	var myFilters []NTypes.Filter
	var action_verb string
	//test is a filter?
	if listOFilters,ok := getValue(ctx.Action_word()).([]NTypes.Filter); ok {
		fmt.Println("ok - we have a list of filters")
		myFilters = listOFilters
		action_verb = "filter"
	} else {
		fmt.Println(reflect.TypeOf(ctx.Action_word()))
		if av,ok := getValue(ctx.Action_word()).(string); ok {
			action_verb = av
			fmt.Println("action verb: ", av)
		} else {
			fmt.Println("uh oh - wasn't able to determine action type")
		}
	}

	var theAction NActions.BaseAction
	switch action_verb {
	case "filter":
		//don't need to do anything here - will assign filters to actions in just a second
	case "extract":
		//todo: keyword 'files' is expected here, but don't worry about it for now
		theAction = &NActions.ExtractNTFS{}
	case "sha1":
		theAction = &NActions.SHA1Action{}
	case "md5":
		theAction = &NActions.MD5Action{}

	default:
		fmt.Println("action was not found: ", action_verb) //parser should prevent us from getting here..
	}

	if action_verb != "filter" {
		theAction.SetFilters(myFilters)
		setValue(ctx, theAction)
	}
	fmt.Println("set filters and value")
}

func (this *TreeShapeListener) EnterDefine(ctx *parser.DefineContext) {
	isList := ctx.LISTOP() != nil
	identifier := ctx.ID().GetText()
	nugget_type := ctx.Nugget_type().GetText()

	//fmt.Println("found a define: ", ctx.ID(), " ", ctx.Nugget_type().GetText(), " is a list?: ", isList)

	if _, exists := registers[identifier]; exists {
		fmt.Println("the variable ", identifier, " already exists!")
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

func (s *TreeShapeListener) ExitAssign(ctx *parser.AssignContext) {
	varIdentifier := ctx.ID(0).GetText()

	//if no actions, then we do a simple calculation and assign it to a register, something like: myimage = "file.dd" as ntfs
	if len(ctx.AllNugget_action()) == 0 {
		//if it's an astype string
		if ctx.AsType() != nil {
			extractTarget := ctx.STRING().GetText()
			extractType := ctx.AsType().GetStop().GetText()
			fmt.Println("a direct assignment has extract info: ", extractTarget, " ", extractType)
			registers[varIdentifier] = NTypes.TNTFS{PathToExtract: extractTarget,AsType:extractType}
		}
	} else {
		actions := ctx.AllNugget_action()

		//setup actions if necessary
		var builtActions []NActions.BaseAction
		for _,action := range actions {
			rawAction := getValue(action)
			//if it's an extract action, we need to look behind and get some more info (like filepath and type)
			if extractAction, ok := rawAction.(NActions.ExtractNTFS); ok {
				//todo: get real values not dummy ones
				extractAction.NTFSImageDataLocation = "G:\\school\\image\\jo.ntfs"
				extractAction.NTFSImageMetadataLocation = "C:\\Users\\drew\\School\\backups_before_upgrade\\jo.extract"
			}
			if act, ok := rawAction.(NActions.BaseAction); ok {
				builtActions = append(builtActions, act)
			}
		}

		//reverse the order of the actions
		for i := len(builtActions)/2 - 1; i >= 0; i-- {
			opp := len(builtActions) - 1 - i
			builtActions[i], builtActions[opp] = builtActions[opp], builtActions[i]
		}

		//we have raw actions, now build the chain of dependencies for each
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
				} else { //was not recognized, shouldn't reach here
					fmt.Println("Error: pattern not recognized.")
				}
			}
			fmt.Println("setting the var ", varIdentifier, " to ", builtActions[0])
			registers[varIdentifier] = builtActions[0]
		}
	}
}

func (s *TreeShapeListener) ExitFilter(ctx *parser.FilterContext) {
	var allFiltersForAction []NTypes.Filter
	for i,_ := range ctx.AllFilter_term() {
		myf := getValue(ctx.Filter_term(i))
		if dep, ok := myf.(NTypes.Filter); ok {
			//fmt.Println("OH MY GOD I THINK I HAVE THIS SYSTEM FIGURED OUT ", dep)
			allFiltersForAction = append(allFiltersForAction , dep)
		}
	}
	fmt.Println(allFiltersForAction)
	setValue(ctx, allFiltersForAction)
}

func (s *TreeShapeListener) ExitAction_word(ctx *parser.Action_wordContext) {
	if ctx.Filter() != nil {
		setValue(ctx, getValue(ctx.Filter()))
	} else {
		setValue(ctx, ctx.GetText())
	}
}

func (s *TreeShapeListener) ExitFilter_term(ctx *parser.Filter_termContext) {
	setValue(ctx, NTypes.Filter{Field: ctx.ID().GetText(), Op:ctx.COMPOP().GetText(), Value:ctx.STRING().GetText()})
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
			//we have a datatype md5
			fmt.Println("have md5 action: ", a)
			a.Execute()
		}
	} else {
		fmt.Println("Error: var '", identifier, "' not found")
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
