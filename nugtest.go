package main

import (
	"./NActions"
	"./NTypes"
	"./nug"
	"bufio"
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// going to need to rethink how to store results.
// if we want to support something like 'x = select pdfs'.. then say 'x add md5, sha1', then say 'print x'

var (
	interactiveMode bool
	pathToInput     string
	registers       map[string]interface{}

	nodeMap map[antlr.ParseTree]interface{}

	typeRegistry map[string]reflect.Type
)

func init() {
	flag.StringVar(&pathToInput, "input", "input.nug", "Path to input")
	flag.BoolVar(&interactiveMode, "interactive", false, "Interactive mode")
	flag.Parse()
	registers = make(map[string]interface{})

	//nodemap can be used to share data across constructs, just store it here whenever and retrieve based on ctx
	nodeMap = make(map[antlr.ParseTree]interface{})

	//hold a string->nugType values
	typeRegistry = make(map[string]reflect.Type)
	setupTypeRegstry()
}

//todo: finish typeregistry so we can clean up typing code
func setupTypeRegstry() {
	typeRegistry["md5"] = reflect.TypeOf(NTypes.MD5{})
	typeRegistry["sha1"] = reflect.TypeOf(NTypes.SHA1{})
	typeRegistry["sha256"] = reflect.TypeOf(NTypes.SHA256{})
	typeRegistry["datetime"] = reflect.TypeOf(NTypes.Datetime{})
	typeRegistry["file"] = reflect.TypeOf(NTypes.FileInfo{})
}

//getValue and setValue will take care of passing around results from individual rules
func setValue(ctx antlr.ParseTree, value interface{}) {
	nodeMap[ctx] = value
}

func getValue(n antlr.ParseTree) interface{} {
	return nodeMap[n]
}

type TreeShapeListener struct {
	*parser.BaseNuggetListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (s *TreeShapeListener) ExitNugget_action(ctx *parser.Nugget_actionContext) {
	var myFilters []NTypes.Filter
	var action_verb string

	// handle filters
	if listOFilters, ok := getValue(ctx.Action_word()).([]NTypes.Filter); ok {
		myFilters = listOFilters
		action_verb = "filter"
		//handle extracts
	} else if _, ok := getValue(ctx.Action_word()).(NTypes.Extract); ok {
		action_verb = "extract"
		//fmt.Println("it's an extract")
		//handle sorts
	} else if _, ok := getValue(ctx.Action_word()).(NTypes.Sort); ok {
		action_verb = "sort"
		//handle unions
	} else if _, ok := getValue(ctx.Action_word()).(NTypes.Union); ok {
		action_verb = "union"
		//handle everything else
	} else {
		if av, ok := getValue(ctx.Action_word()).(string); ok {
			action_verb = av
		} else {
			fmt.Println("error - wasn't able to determine action type")
		}
	}

	var theAction NActions.BaseAction
	switch action_verb {
	case "filter":
		//don't need to do anything here - will assign filters to actions in just a second
		theAction = &NActions.FilterAction{}
	case "sort":
		thisSortField := getValue(ctx.Action_word()).(NTypes.Sort)
		theAction = &NActions.SortAction{SortField: thisSortField.Field}
	case "extract":
		extractType := getValue(ctx.Action_word()).(NTypes.Extract)
		if extractType.AsType == "pcap" {
			theAction = &NActions.ExtractPCAP{}
		} else if extractType.AsType == "ntfs" {
			theAction = &NActions.ExtractNTFS{}
		} else if extractType.AsType == "md5hashes" {
			theAction = &NActions.ExtractList{ListType: "md5", ListLocation: extractType.PathToExtract}
		} else if extractType.AsType == "sha1hashes" {
			theAction = &NActions.ExtractList{ListType: "sha1", ListLocation: extractType.PathToExtract}
		} else if extractType.AsType == "sha256hashes" {
			theAction = &NActions.ExtractList{ListType: "sha256", ListLocation: extractType.PathToExtract}
		} else if extractType.AsType == "memory" {
			theAction = &NActions.ExtractMemory{}
		} else {
			fmt.Println("Error parsing given type: ", extractType.AsType)
		}
	case "sha1":
		theAction = &NActions.SHA1Action{}
	case "sha256":
		theAction = &NActions.SHA256Action{}
	case "md5":
		theAction = &NActions.MD5Action{}
	case "diskinfo":
		theAction = &NActions.DiskInfoAction{}
	case "union":
		//todo: figure out how to pass file info forward as well? -- actually, maybe not needed - we're doing exactly what was told to be done
		unionType := getValue(ctx.Action_word()).(NTypes.Union)
		var theListFromVar []string
		if val, ok := registers[unionType.AgainstVarName].(NActions.BaseAction); ok {
			fmt.Println(val.GetResults())
			theListFromVar = val.GetResults().([]string)
		} else {
			fmt.Println("Error! Was not able to find var: ", unionType.AgainstVarName)
		}
		theAction = &NActions.UnionAction{VariableList: theListFromVar}
	case "pslist":
		theAction = &NActions.ProcessListAction{}
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
		//todo: minimize this code by using a registry, ie map types to strings as I do in setupTypeRegstry
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
		case "md5": //TODO: investigate impact of 'natural types' such as string - see if we should wrap in an NType
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

func (s *TreeShapeListener) ExitAssign(ctx *parser.AssignContext) {
	varIdentifier := ctx.ID(0).GetText()

	actions := ctx.AllNugget_action()
	//setup actions if necessary
	var builtActions []NActions.BaseAction
	for _, action := range actions {
		rawAction := getValue(action)
		//if it's an extract action, we need to look behind and get some more info (like filepath and type)
		//fmt.Println("action is ", reflect.TypeOf(rawAction )," : ", rawAction )

		if extractAction, ok := rawAction.(*NActions.ExtractNTFS); ok {
			//todo: get real values not dummy ones
			extractAction.NTFSImageDataLocation = "G:\\school\\image\\jo.ntfs"
			extractAction.NTFSImageMetadataLocation = "G:\\school\\jo.extract"
		}
		if extractAction, ok := rawAction.(*NActions.ExtractPCAP); ok {
			//todo: get real values not dummy ones
			//extractAction.PCAPLocation = "G:\\school\\sample.pcap"
			//extractAction.PCAPLocation = "G:\\school\\m57\\data\\net-2009-11-19-09_54.pcap"

			extractAction.PCAPLocation = "G:\\school\\m57\\data\\net-2009-12-09-11_59.pcap"
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
			//fmt.Println("action at index: ", index, "is ", reflect.TypeOf(builtAction)," : ", builtAction, " and depends on: ", builtActions[index+1])
			var depAction NActions.BaseAction
			depAction = builtActions[index+1]
			builtAction.(NActions.BaseAction).SetDependency(depAction)
		} else {
			//fmt.Println("action at index: ", index,"is ", reflect.TypeOf(builtAction)," : ", builtAction, " and has no dependency. Setting dep to the var")
			if len(ctx.AllID()) > 1 {
				depVar := ctx.ID(1).GetText()

				// is it an existing var?
				if _, ok := registers[depVar]; ok {
					//if it's an action..
					if dep, ok := registers[depVar].(NActions.BaseAction); ok {
						//fmt.Println("the dependency for this action will be variable: ", nVar)
						builtAction.(NActions.BaseAction).SetDependency(dep)
					}
				} else {
					fmt.Println("Error: Var '", depVar, "' not recognized.")
				}
			}
		}
		//fmt.Println("setting the var ", varIdentifier, " to ", builtActions[0])
		registers[varIdentifier] = builtActions[0]
	}
}

func (s *TreeShapeListener) ExitFilter(ctx *parser.FilterContext) {
	var allFiltersForAction []NTypes.Filter
	for i, _ := range ctx.AllFilter_term() {
		myf := getValue(ctx.Filter_term(i))
		if dep, ok := myf.(NTypes.Filter); ok {
			allFiltersForAction = append(allFiltersForAction, dep)
		}
	}
	setValue(ctx, allFiltersForAction)
}

func (s *TreeShapeListener) ExitNugget_type(ctx *parser.Nugget_typeContext) {
	setValue(ctx, ctx.GetText())
}

func (s *TreeShapeListener) ExitAsType(ctx *parser.AsTypeContext) {
	//fmt.Println("setting exit as type to: ", getValue(ctx.Nugget_type()))
	setValue(ctx, getValue(ctx.Nugget_type()))
}

func (s *TreeShapeListener) ExitAction_word(ctx *parser.Action_wordContext) {
	//handle extractions
	if ctx.AsType() != nil {
		myT := getValue(ctx.AsType())
		//fmt.Println("got: ", myT)
		if myT == "pcap" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\sample.pcap", AsType: "pcap"})
		} else if myT == "ntfs" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\jo.ntfs", AsType: "ntfs"})
		} else if myT == "listof-md5" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\md5hashes.txt", AsType: "md5hashes"})
		} else if myT == "listof-sha1" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\sha1hashes.txt", AsType: "sha1hashes"})
		} else if myT == "listof-sha256" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\sha256hashes.txt", AsType: "sha256hashes"})
		} else if myT == "memory" {
			setValue(ctx, NTypes.Extract{PathToExtract: "G:\\school\\jo-2009-12-03.mddramimage.zip", AsType: "memory"})
		} else {
			fmt.Println("error on type extraction")
		}
		//handle filters
	} else if ctx.Filter() != nil {
		setValue(ctx, getValue(ctx.Filter()))
		//handle sorts
	} else if ctx.ByField() != nil {
		//fmt.Println("sort by: ", getValue(ctx.ByField()))
		if val, ok := getValue(ctx.ByField()).(string); ok {
			//fmt.Println("sort string: ", val)
			setValue(ctx, NTypes.Sort{Field: val})
		}
		//handle union
	} else if ctx.ID() != nil {
		setValue(ctx, NTypes.Union{Results: []string{""}, AgainstVarName: ctx.ID().GetText()})
		//handle grep
	} else {
		setValue(ctx, ctx.GetText())
	}
}

func (s *TreeShapeListener) ExitFilter_term(ctx *parser.Filter_termContext) {
	setValue(ctx, NTypes.Filter{Field: ctx.ID().GetText(), Op: ctx.COMPOP().GetText(), Value: ctx.STRING().GetText()})
}

func (s *TreeShapeListener) ExitByField(ctx *parser.ByFieldContext) {
	setValue(ctx, ctx.ID().GetText())
}

func (s *TreeShapeListener) ExitSingleton_var(ctx *parser.Singleton_varContext) {
	theVar := ctx.ID().GetText()
	if v, ok := registers[theVar]; ok {
		//fmt.Println(theVar, "[", reflect.TypeOf(v),"]:", v)
		if ba, ok := v.(NActions.BaseAction); ok {
			fmt.Println("Results for var ", theVar, ": ", ba.GetResults())
			//ba.GetResults()
			//fmt.Println("cutting off results for now..")
		} else {
			fmt.Println("couldn't execute var : ", theVar, "because it is not of baseAction type")
		}
	} else {
		fmt.Println("Variable not recongized:" + theVar)
		os.Exit(1)
	}
}

func (s *TreeShapeListener) ExitOperation_on_singleton(ctx *parser.Operation_on_singletonContext) {
	var operation string
	if op, ok := getValue(ctx.Singleton_op()).(string); ok {
		operation = op
	}
	theVar := ctx.ID().GetText()
	var subfield string
	if strings.Contains(theVar, ".") {
		root := strings.Split(theVar, ".")[0]
		subfield = strings.Split(theVar, ".")[1]
		theVar = root
	}

	if val, ok := registers[theVar].(NActions.BaseAction); ok {
		switch operation {
		case "type":
			fmt.Println(reflect.TypeOf(val))
		case "print":
			fmt.Println(val)
		case "typex":
			fmt.Println(reflect.TypeOf(val.GetResults()))
		case "printx":
			myResults := val.GetResults()
			fmt.Println(subfield)
			fmt.Println(reflect.TypeOf(myResults))

			t := reflect.TypeOf(myResults)
			for i := 0; i < t.NumField(); i++ {
				ft := t.Field(i).Type
				if ft.Kind() == reflect.Ptr {
					ft = ft.Elem()
				}
				//				fmt.Println(ft.Field(0).Name)
				fmt.Sprintf("percent v: %v\n", ft)
			}

			st := reflect.ValueOf(myResults)
			typeOfTE := st.Type()
			for i := 0; i < st.NumField(); i++ {
				f := st.Field(i)
				fmt.Printf("%d: %s %s\n", i,
					typeOfTE.Field(i).Name, f.Type())
			}

		case "raw":
			if files, ok := val.GetResults().([]NTypes.FileInfo); ok {
				for _, fi := range files {
					fmt.Println(fi)
				}
			} else {
				fmt.Println(val.GetResults())
			}
		case "size":
			fmt.Println("len not implemented yet")
		default:
			fmt.Println("operation not recognized..")
		}
	} else {
		fmt.Println("Variable not recognized: ", theVar)
	}
}

func (s *TreeShapeListener) ExitSingleton_op(ctx *parser.Singleton_opContext) {
	setValue(ctx, ctx.GetText())
}

func (s *TreeShapeListener) ExitByteOffsetSize(ctx *parser.ByteOffsetSizeContext) {
	byteOffset, err := strconv.Atoi(ctx.INT(0).GetText())
	clusterSize, err := strconv.Atoi(ctx.INT(1).GetText())
	if err != nil {
		fmt.Println("Error parsing byte offset: ", err)
	}
	setValue(ctx, NTypes.OffsetInfo{byteOffset, clusterSize})
}

func TabOrNewline(r rune) bool {
	return r == '\t' || r == '\n'
}

func main() {
	const placeOfInterest = '\t'

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")

	fmt.Printf("%x ", placeOfInterest)

	fmt.Printf("\n")

	if interactiveMode {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("nugget> ")

		var myByte []byte
		myByte = make([]byte, 1)

		text, err := reader.Read(myByte)

		if err != nil {
			fmt.Println("Input error: ", err)
		}

		fmt.Println(text)

	} else {
		file, err := os.Open(pathToInput)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if scanner.Text() == "" {
				continue
			}
			fmt.Printf("nugget> %s\n", scanner.Text())
			input := antlr.NewInputStream(scanner.Text())
			lexer := parser.NewNuggetLexer(input)
			tokenstream := antlr.NewCommonTokenStream(lexer, 0)
			tokenparser := parser.NewNuggetParser(tokenstream)
			tokenparser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			tokenparser.BuildParseTrees = true
			tree := tokenparser.Prog()
			antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
			fmt.Println()
		}
	}
}
