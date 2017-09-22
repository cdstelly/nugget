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

	"reflect"
	"net/rpc"
	"log"
)

var (
	pathToInput string
	registers map[string]interface{}

	nodeMap map[antlr.ParseTree]interface{}

	typeRegistry map[string]reflect.Type

)

func init() {
	flag.StringVar(&pathToInput, "input", "input2.nug", "Path to input")
	flag.Parse()
	registers = make(map[string]interface{})

	//nodemap can be used to share data across constructs, just store it here whenever and retrieve based on ctx
	nodeMap = make(map[antlr.ParseTree]interface{})

	//hold a string->nugType values
	typeRegistry = make(map[string]reflect.Type)
	setupTypeRegstry()
}

func setupTypeRegstry() {
	typeRegistry["md5"] = reflect.TypeOf(NTypes.MD5{})
	typeRegistry["sha1"] = reflect.TypeOf(NTypes.SHA1{})
	typeRegistry["sha256"] = reflect.TypeOf(NTypes.SHA256{})
	typeRegistry["datetime"] = reflect.TypeOf(NTypes.Datetime{})
	typeRegistry["file"] = reflect.TypeOf(NTypes.FileInfo{})
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

func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {

}

func (s *TreeShapeListener) ExitNugget_action(ctx *parser.Nugget_actionContext) {
	var myFilters []NTypes.Filter
	var action_verb string

	//we have some filters
	if listOFilters, ok := getValue(ctx.Action_word()).([]NTypes.Filter); ok {
		myFilters = listOFilters
		action_verb = "filter"
	} else if _, ok := getValue(ctx.Action_word()).(NTypes.Extract); ok {
		action_verb = "extract"
	} else {
		if av,ok := getValue(ctx.Action_word()).(string); ok {
			action_verb = av
		} else {
			fmt.Println("uh oh - wasn't able to determine action type")
		}
	}

	var theAction NActions.BaseAction
	switch action_verb {
	case "filter":
		//don't need to do anything here - will assign filters to actions in just a second
		fmt.Println("Here we go boys!")
		theAction = &NActions.FilterAction{}
	case "extract":

		extractType := getValue(ctx.Action_word()).(NTypes.Extract)
		if extractType.AsType == "pcap" {
			theAction = &NActions.ExtractPCAP{}
		} else if extractType.AsType == "ntfs" {
			theAction = &NActions.ExtractNTFS{}
		} else {
			fmt.Println("Error parsing given type: ", extractType.AsType)
		}
	case "sha1":
		theAction = &NActions.SHA1Action{}
	case "sha256":
		theAction = &NActions.SHA256Action{}
	case "md5":
		theAction = &NActions.MD5Action{}
	default:
		fmt.Println("action was not found: ", action_verb) //parser should prevent us from getting here..
	}

	if action_verb != "filter" {
		theAction.SetFilters(myFilters)
		setValue(ctx, theAction)
	} else {
		theAction.SetFilters(myFilters)
		setValue(ctx, theAction)
	}
}

func (s *TreeShapeListener) ExitDefine(ctx *parser.DefineContext) {
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
				registers[identifier] = []NTypes.Extract{}
			} else {
				registers[identifier] = NTypes.Extract{}
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
		case "pcap":
			if isList {
				registers[identifier] = []NTypes.Extract{}
			} else {
				registers[identifier] = NTypes.Extract{}
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

func (s *TreeShapeListener) ExitDefine_tuple(ctx *parser.Define_tupleContext) {
	// isList := ctx.LISTOP() != nil  //todo: implement lists of tuples... shudder....
	identifier := ctx.ID().GetText()

	var theTuples []interface{}

	for _, t := range ctx.AllNugget_type() {
		//fmt.Println(t.GetText())
		v := reflect.New(typeRegistry[t.GetText()])
		theTuples = append(theTuples, v)
	}

	registers[identifier] = theTuples
}

func (s *TreeShapeListener) EnterAssign(ctx *parser.AssignContext) {
	//currentVariable = ctx.ID(0).GetText()
	//fmt.Println("cv:",currentVariable)
}

// howbout a new approach to actions

// every exit action, we build and append to a list
// this list is unique for every line of input
// this means that list[4] relies on list[3] to execute and provide results for list[4]


func (s *TreeShapeListener) ExitAssign(ctx *parser.AssignContext) {
	varIdentifier := ctx.ID(0).GetText()

	//if no actions, then we do a simple calculation and assign it to a register, something like: myimage = "file.dd" as ntfs
	/* removed when moving 'astype' with 'extract' action
	if ctx.AsType() != nil {
		extractTarget := ctx.STRING().GetText()
		extractType := ctx.AsType().GetStop().GetText()
		//fmt.Println("a direct assignment has extract info: ", extractTarget, " ", extractType)
		registers[varIdentifier] = NTypes.Extract{PathToExtract: extractTarget,AsType:extractType}
	} else {
		*/


	actions := ctx.AllNugget_action()
	//setup actions if necessary
	var builtActions []NActions.BaseAction
	for _,action := range actions {
		rawAction := getValue(action)
		//if it's an extract action, we need to look behind and get some more info (like filepath and type)
		fmt.Println("action is ", reflect.TypeOf(rawAction )," : ", rawAction )

		if extractAction, ok := rawAction.(*NActions.ExtractNTFS); ok {
			//todo: get real values not dummy ones
			extractAction.NTFSImageDataLocation = "G:\\school\\image\\jo.ntfs"
			extractAction.NTFSImageMetadataLocation = "G:\\school\\jo.extract"
		}
		if extractAction, ok := rawAction.(*NActions.ExtractPCAP); ok {
			//todo: get real values not dummy ones
			extractAction.PCAPLocation = "G:\\school\\sample.pcap"
		}
		if act, ok := rawAction.(NActions.BaseAction); ok {
			builtActions = append(builtActions, act)
		}
	}

	//reverse the order of the actions [so that item 1 depends on item 2, etc.]
	for i := len(builtActions)/2 - 1; i >= 0; i-- {
		opp := len(builtActions) - 1 - i
		builtActions[i], builtActions[opp] = builtActions[opp], builtActions[i]
	}

	//we have raw actions, now build the chain of dependencies for each
	for index, builtAction := range builtActions {
		if index+1 < len(builtActions) {
			fmt.Println("action at index: ", index, "is ", reflect.TypeOf(builtAction)," : ", builtAction, " and depends on: ", builtActions[index+1])
			var depAction NActions.BaseAction
			depAction = builtActions[index+1]
			builtAction.(NActions.BaseAction).SetDependency(depAction)
		} else {
			fmt.Println("action at index: ", index,"is ", reflect.TypeOf(builtAction)," : ", builtAction, " and has no dependency. Setting dep to the var")
			if len(ctx.AllID()) > 1 {
				depVar := ctx.ID(1).GetText()

				// is it an existing var?
				if _, ok := registers[depVar]; ok {
					//if it's an action..
					if dep, ok := registers[depVar].(NActions.BaseAction); ok {
						//fmt.Println("the dependency for this action will be variable: ", nVar)
						builtAction.(NActions.BaseAction).SetDependency(dep)
					}
					// else what if it's an extraction??
				} else {
					fmt.Println("Error: Var '", depVar, "' not recognized.")
				}
			}
		}
		fmt.Println("setting the var ", varIdentifier, " to ", builtActions[0])
		registers[varIdentifier] = builtActions[0]
	}
}

func (s *TreeShapeListener) ExitFilter(ctx *parser.FilterContext) {
	var allFiltersForAction []NTypes.Filter
	for i,_ := range ctx.AllFilter_term() {
		myf := getValue(ctx.Filter_term(i))
		if dep, ok := myf.(NTypes.Filter); ok {
			allFiltersForAction = append(allFiltersForAction , dep)
		}
	}
	setValue(ctx, allFiltersForAction)
}

func (s *TreeShapeListener) ExitNugget_type(ctx *parser.Nugget_typeContext) {
	setValue(ctx, ctx.GetText())
}

func (s *TreeShapeListener) ExitAsType(ctx *parser.AsTypeContext) {
	fmt.Println("setting exit as type to: ", getValue(ctx.Nugget_type()))
	setValue(ctx, getValue(ctx.Nugget_type()))
}

//investigate from here -
// maybe we generate an extract type here and pass it up the chain, just like we do for flter
func (s *TreeShapeListener) ExitAction_word(ctx *parser.Action_wordContext) {

	if ctx.AsType() != nil {
		myT := getValue(ctx.AsType())
		//fmt.Println("got: ", myT)
		if myT == "pcap" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\sample.pcap", AsType: "pcap"} )
		} else if myT == "ntfs" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\jo.ntfs", AsType: "ntfs"} )
		} else {
			fmt.Println("error on type extraction")
		}
	} else if ctx.Filter() != nil {
		setValue(ctx, getValue(ctx.Filter()))
	} else {
		setValue(ctx, ctx.GetText())
	}
}

func (s *TreeShapeListener) ExitFilter_term(ctx *parser.Filter_termContext) {
	setValue(ctx, NTypes.Filter{Field: ctx.ID().GetText(), Op:ctx.COMPOP().GetText(), Value:ctx.STRING().GetText()})
}

func (s *TreeShapeListener) ExitSingleton_var(ctx *parser.Singleton_varContext) {
	theVar := ctx.ID().GetText()
	if v, ok := registers[theVar]; ok {
		fmt.Println(theVar, "[", reflect.TypeOf(v),"]:", v)
		if ba,ok := v.(NActions.BaseAction); ok {
			fmt.Println("Results for var:",theVar, ": ", ba.GetResults())
			//ba.GetResults()
			//fmt.Println("cutting off results for now..")
		} else {
			fmt.Println("couldn't execute var : ", theVar	, "because it is not of baseAction type")
		}
	} else {
		fmt.Println("Variable not recongized:" + theVar)
		os.Exit(1)
	}
}

func (s *TreeShapeListener) ExitOperation_on_singleton(ctx *parser.Operation_on_singletonContext) {
	var operation string
	if op, ok := getValue(ctx.Singleton_op()).(string);ok {
		operation = op
	}

	theVar := ctx.ID().GetText()
	if val, ok := registers[theVar].(NActions.BaseAction); ok {
		switch operation {
		case "type":
			fmt.Println(reflect.TypeOf(val))
		case "print":
			fmt.Println(val)
		case "typex":
			fmt.Println(reflect.TypeOf(val.GetResults()))
		case "printx":
			fmt.Println(val.GetResults())
		case "size":
			fmt.Println("len not implemented yet")
		default:
			fmt.Println("operation not recognized..")
		}
	} else {
		fmt.Println("var not recognized:", theVar)
	}
}

func (s *TreeShapeListener) ExitSingleton_op(ctx *parser.Singleton_opContext) {
	setValue(ctx, ctx.GetText())
}

type NugArg struct {
	TheData []byte
}

type NugData int

func testRemoteMD5() {
	client, err := rpc.DialHTTP("tcp", "192.168.1.198:2000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &NugArg{[]byte("test")}
	var reply string
	err = client.Call("NugData.DoMD5", args, &reply)
	if err != nil {
		log.Fatal("md5 error:", err)
	}
	fmt.Printf("md5: %s=%s", string(args.TheData), reply)
	os.Exit(0)
}

func testRemoteTSK() {
	client, err := rpc.DialHTTP("tcp", "192.168.1.198:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//load some data into tsk memory
	args := &NugArg{[]byte("test")}
	var reply string
	err = client.Call("NugTSK.LoadData", args, &reply)
	if err != nil {
		log.Fatal("tsk load error:", err)
	}
	fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)

	//execute data len test
	arg2 := &NugArg{[]byte("")}
	err = client.Call("NugTSK.GetDataLen", arg2, &reply)
	if err != nil {
		log.Fatal("tsk len error:", err)
	}
	fmt.Printf("tsk len =%s\n", reply)

	os.Exit(0)
}

func main() {
	//testRemoteMD5()
	//testRemoteTSK()

	file, err := os.Open(pathToInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" { continue }
		fmt.Printf("nugget> %s\n", scanner.Text())
		input := antlr.NewInputStream(scanner.Text())
		lexer := parser.NewNugget2Lexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewNugget2Parser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true
		tree := p.Prog()
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
		fmt.Println()
	}
}
