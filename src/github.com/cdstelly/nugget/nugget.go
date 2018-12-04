package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/cdstelly/nugget/NTypes"
	"github.com/cdstelly/nugget/expressions/extractors"
	"github.com/cdstelly/nugget/expressions/transforms"
	"github.com/cdstelly/nugget/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"strings"
)

var (
	interactiveMode bool
	pathToInput     string
	runtimeServer   string
	registers       map[string]interface{}

	nodeMap map[antlr.ParseTree]interface{}

	typeRegistry     map[string]reflect.Type
	interruptCounter int

	tskClient *rpc.Client
	volClient *rpc.Client
)

func init() {
	flag.StringVar(&pathToInput, "input", "", "Path to input")
	flag.BoolVar(&interactiveMode, "interactive", false, "Interactive mode")
	flag.StringVar(&runtimeServer, "runtimeIP", "127.0.0.1", "Network address of runtime server")
	flag.Parse()

	if !flagCheck() {
		flag.PrintDefaults()
		os.Exit(0)
	}
	registers = make(map[string]interface{})

	//nodemap can be used to share data across constructs, just store it here whenever and retrieve based on ctx
	nodeMap = make(map[antlr.ParseTree]interface{})

	//hold a string->nugType values
	typeRegistry = make(map[string]reflect.Type)
	setupTypeRegstry()

	CatchTerm()
	interruptCounter = 0
}

func flagCheck() bool {
	if pathToInput == "" {
		if !interactiveMode {
			return false
		}
	}
	return true
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
	var actionVerb string

	// handle filters
	if listOFilters, ok := getValue(ctx.Action_word()).([]NTypes.Filter); ok {
		myFilters = listOFilters
		actionVerb = "filter"
		//handle extracts
	} else if _, ok := getValue(ctx.Action_word()).(NTypes.Extract); ok {
		actionVerb = "extract"
		//fmt.Println("it's an extract")
		//handle sorts
	} else if _, ok := getValue(ctx.Action_word()).(NTypes.Sort); ok {
		actionVerb = "sort"
		//handle unions
	} else if _, ok := getValue(ctx.Action_word()).(NTypes.Union); ok {
		actionVerb = "union"
		//handle everything else
	} else {
		if av, ok := getValue(ctx.Action_word()).(string); ok {
			actionVerb = av
		} else {
			fmt.Println("error - wasn't able to determine action type")
		}
	}

	var theAction expressions.BaseAction
	switch actionVerb {
	case "filter":
		//don't need to do anything here - will assign filters to actions in just a second
		theAction = &expressions.FilterAction{}
	case "sort":
		thisSortField := getValue(ctx.Action_word()).(NTypes.Sort)
		theAction = &expressions.SortAction{SortField: thisSortField.Field}
	case "extract":
		extractType := getValue(ctx.Action_word()).(NTypes.Extract)
		if extractType.AsType == "pcap" {
			theAction = &extractors.ExtractPCAP{}
		} else if extractType.AsType == "ntfs" {
			theAction = &extractors.ExtractNTFS{Location: extractType.PathToExtract, TskClient: tskClient}
		} else if extractType.AsType == "md5hashes" {
			theAction = &extractors.ExtractList{ListType: "md5", ListLocation: extractType.PathToExtract}
		} else if extractType.AsType == "sha1hashes" {
			theAction = &extractors.ExtractList{ListType: "sha1", ListLocation: extractType.PathToExtract}
		} else if extractType.AsType == "sha256hashes" {
			theAction = &extractors.ExtractList{ListType: "sha256", ListLocation: extractType.PathToExtract}
		} else if extractType.AsType == "memory" {
			theAction = &extractors.ExtractMemory{Location: extractType.PathToExtract, VolClient: volClient}
		} else if extractType.AsType == "http" {
			theAction = &expressions.HTTPAction{}
		} else {
			fmt.Println("Error parsing given type: ", extractType.AsType)
		}
	case "sha1":
		theAction = &expressions.SHA1Action{}
	case "sha256":
		theAction = &expressions.SHA256Action{}
	case "md5":
		theAction = &expressions.MD5Action{}
	case "diskinfo":
		theAction = &expressions.DiskInfoAction{}
	case "union":
		//todo: figure out how to pass file info forward as well? -- actually, maybe not needed - we're doing exactly what was told to be done
		unionType := getValue(ctx.Action_word()).(NTypes.Union)
		var theListFromVar []string
		if val, ok := registers[unionType.AgainstVarName].(expressions.BaseAction); ok {
			fmt.Println(val.GetResults())
			theListFromVar = val.GetResults().([]string)
		} else {
			fmt.Println("Error! Was not able to find var: ", unionType.AgainstVarName)
		}
		theAction = &expressions.UnionAction{VariableList: theListFromVar}
	case "pslist":
		theAction = &expressions.ProcessListAction{}
	default:
		fmt.Println("action was not found")
	}

	if actionVerb != "filter" {
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
	var builtActions []expressions.BaseAction
	for _, action := range actions {
		rawAction := getValue(action)
		//if it's an extract action, we need to look behind and get some more info (like filepath and type)
		//fmt.Println("action is ", reflect.TypeOf(rawAction )," : ", rawAction )

		if extractAction, ok := rawAction.(*extractors.ExtractNTFS); ok {
			url := ctx.STRING().GetText()
			extractAction.Location = strings.Trim(url, `"`)
		}
		if extractAction, ok := rawAction.(*extractors.ExtractMemory); ok {
			url := ctx.STRING().GetText()
			extractAction.Location = strings.Trim(url, `"`)
		}
		if extractAction, ok := rawAction.(*extractors.ExtractPCAP); ok {
			url := ctx.STRING().GetText()
			extractAction.PCAPLocation = strings.Trim(url, `"`)
		}
		if act, ok := rawAction.(expressions.BaseAction); ok {
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
			var depAction expressions.BaseAction
			depAction = builtActions[index+1]
			builtAction.(expressions.BaseAction).SetDependency(depAction)
		} else {
			//fmt.Println("action at index: ", index,"is ", reflect.TypeOf(builtAction)," : ", builtAction, " and has no dependency. Setting dep to the var")
			if len(ctx.AllID()) > 1 {
				depVar := ctx.ID(1).GetText()

				// is it an existing var?
				if _, ok := registers[depVar]; ok {
					//if it's an action..
					if dep, ok := registers[depVar].(expressions.BaseAction); ok {
						//fmt.Println("the dependency for this action will be variable: ", nVar)
						builtAction.(expressions.BaseAction).SetDependency(dep)
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
		f := getValue(ctx.Filter_term(i))
		if dep, ok := f.(NTypes.Filter); ok {
			allFiltersForAction = append(allFiltersForAction, dep)
		}
	}
	setValue(ctx, allFiltersForAction)
}

func (s *TreeShapeListener) ExitNugget_type(ctx *parser.Nugget_typeContext) {
	setValue(ctx, ctx.GetText())
}

func (s *TreeShapeListener) ExitAsType(ctx *parser.AsTypeContext) {
	setValue(ctx, getValue(ctx.Nugget_type()))
}

func (s *TreeShapeListener) ExitAction_word(ctx *parser.Action_wordContext) {
	//handle extractions
	if ctx.AsType() != nil {
		givenType := getValue(ctx.AsType())
		if givenType == "pcap" {
			setValue(ctx, NTypes.Extract{AsType: "pcap"})
		} else if givenType == "ntfs" {
			setValue(ctx, NTypes.Extract{PathToExtract: "jo.ntfs", AsType: "ntfs"})
		} else if givenType == "listof-md5" {
			setValue(ctx, NTypes.Extract{PathToExtract: "md5hashes.txt", AsType: "md5hashes"})
		} else if givenType == "listof-sha1" {
			setValue(ctx, NTypes.Extract{PathToExtract: "sha1hashes.txt", AsType: "sha1hashes"})
		} else if givenType == "listof-sha256" {
			setValue(ctx, NTypes.Extract{PathToExtract: "sha256hashes.txt", AsType: "sha256hashes"})
		} else if givenType == "memory" {
			setValue(ctx, NTypes.Extract{PathToExtract: "jo-2009-12-03.mddramimage.zip", AsType: "memory"})
		} else if givenType == "http" {
			setValue(ctx, NTypes.Extract{AsType: "http"})
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
			//fmt.Println("sort string: ", val)-
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
		if ba, ok := v.(expressions.BaseAction); ok {
			//fmt.Println("Results for var ", theVar, ": ", ba.GetResults())
			ba.GetResults()
			//fmt.Println("cutting off results for now..")
		} else {
			fmt.Println("couldn't execute var : ", theVar, "because it is not of baseAction type")
		}
	} else {
		fmt.Println("Variable not recongized:" + theVar)
		// os.Exit(1)
	}
}

func ContainsSubfield(input string) bool {
	return strings.Contains(input, ".")
}

func GetSubfield(input string) string {
	return strings.Split(input, ".")[1]
}

//Use reflection to obtain the data stored in the given subfield
//There are two basic cases to handle - whether or not the given action yields a set of results or a single result
func GetResultsOfSubfield(rootAction expressions.BaseAction, subfield string) []string {
	var ResultsOfSubfield []string
	resultsOfActionExecution := rootAction.GetResults()
	//fmt.Println("the subfield: " + subfield)
	var fieldList []string
	st := reflect.ValueOf(resultsOfActionExecution)
	reflectType := reflect.TypeOf(resultsOfActionExecution)
	switch reflectType.Kind() {
	case reflect.Slice:
		//todo there has to be some voodoo that streamlines this..look into reflecting on slices of interfaces
		//iterate through everything, 'convert' it to basetype, reflect on the subfield, print the subfield
		for i := 0; i < st.Len(); i++ {
			//st[i].interface will now satisfy the basetype
			instanceFromList := st.Index(i).Interface()
			if t, ok := instanceFromList.(NTypes.BaseType); ok {
				value := reflect.ValueOf(t)
				typeOfValue := value.Type()
				fieldFound := false
				var finalField reflect.Value
				for i := 0; i < value.NumField(); i++ {
					if subfield == typeOfValue.Field(i).Name {
						fieldFound = true
						finalField = value.Field(i)
					}
				}
				if fieldFound {
					//fmt.Println("The subfield: " + subfield + " has value: \n" , finalField.String())
					//fmt.Println(">>", finalField.String())

					//fmt.Println("<  ", finalField, finalField.Kind())
					var subfieldAsString string

					switch finalField.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						//fmt.Println(">>>", strconv.FormatInt(finalField.Int(),10))
						subfieldAsString = strconv.FormatInt(finalField.Int(), 10)
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						//	fmt.Println(">>>", strconv.FormatUint(finalField.Uint(),10))
						subfieldAsString = strconv.FormatUint(finalField.Uint(), 10)
					case reflect.String:
						//fmt.Println(">>>", finalField.String())
						subfieldAsString = finalField.String()
					}
					ResultsOfSubfield = append(ResultsOfSubfield, subfieldAsString)
				} else {
					fmt.Printf("Error: subfield '%s' does not exist for type: '%s'. \n", subfield, typeOfValue.String())
				}
			}
		}
	default:
		typeOfTE := st.Type()
		fieldFound := false
		var finalField reflect.Value
		//Using reflection, iterate through all subfields of the type of the input. Compare to what we're given to printout.
		for i := 0; i < st.NumField(); i++ {
			fieldList = append(fieldList, typeOfTE.Field(i).Name)
			if subfield == typeOfTE.Field(i).Name {
				fieldFound = true
				finalField = st.Field(i)
			}
		}
		if fieldFound {
			//fmt.Println("The subfield: " + subfield + " has value: \n" + finalField.String())
			ResultsOfSubfield = append(ResultsOfSubfield, finalField.String())
		} else {
			fmt.Printf("Error: subfield '%s' does not exist for type: '%s'. \nPossibilites: %s", subfield, typeOfTE.String(), fieldList)
		}
	}
	return ResultsOfSubfield
}

func fieldRootsMatch(fields []string) bool {
	if len(fields) <= 1 {
		return true
	}

	firstRootTerm := fields[0]
	if ContainsSubfield(firstRootTerm) {
		firstRootTerm = strings.Split(firstRootTerm, ".")[0]
	}

	for _, f := range fields {
		if ContainsSubfield(f) {
			root := strings.Split(f, ".")[0]
			if root != firstRootTerm {
				return false
			}
		}
	}
	return true
}

func GetResultsForActionWithSpecificFields(Action expressions.BaseAction, fields []string) [][]string {
	var MappedResults [][]string
	MappedResults = make([][]string, len(fields)) // MappedResults is []([]string)

	//check that each field has the same root value
	if fieldRootsMatch(fields) == true {
		//continue processing
		for idx, f := range fields {
			if ContainsSubfield(f) {
				MappedResults[idx] = GetResultsOfSubfield(Action, GetSubfield(f))
			} else {
				rawStringOfResults := fmt.Sprintf("%v", Action.GetResults())
				sliceOfResults := strings.Split(rawStringOfResults, " ")
				MappedResults[idx] = sliceOfResults
			}
			//fmt.Println("len: ", len(MappedResults[f]), " type: ", reflect.TypeOf(MappedResults[f]))
		}
	} else {
		fmt.Println("Error - field roots must match")
		return nil
	}
	return MappedResults
}

//for each term,add to a list
//print each entry with each field...
//todo: graceful error handling when getresults returns a connection failed
func (s *TreeShapeListener) ExitOperation_on_singleton(ctx *parser.Operation_on_singletonContext) {
	var fieldsToPrint []string
	var operation string
	var BaseVariable string

	if op, ok := getValue(ctx.Singleton_op()).(string); ok {
		operation = op
	}
	for _, id := range ctx.AllID() {
		givenVar := id.GetText()
		fieldsToPrint = append(fieldsToPrint, givenVar)
	}

	//fmt.Println("Going to do:", operation, "on ", fieldsToPrint)

	BaseVariable = strings.Split(ctx.ID(0).GetText(), ".")[0]
	var ActionForEvaluation expressions.BaseAction
	if val, ok := registers[BaseVariable].(expressions.BaseAction); ok {
		ActionForEvaluation = val
	} else {
		fmt.Println("Error: Variable not recognized: ", BaseVariable)
		return
	}

	switch operation {
	case "type":
		fmt.Println(reflect.TypeOf(ActionForEvaluation.GetResults()))
	case "print":
		resultSliceOfSlices := GetResultsForActionWithSpecificFields(ActionForEvaluation, fieldsToPrint)
		//fmt.Println(resultSliceOfSlices)

		maxIndex := len(resultSliceOfSlices[0])
		for indexCounter := 0; indexCounter < maxIndex; indexCounter++ {
			row := ""
			for fieldCounter := 0; fieldCounter < len(resultSliceOfSlices); fieldCounter++ {
				row = row + resultSliceOfSlices[fieldCounter][indexCounter] + " "
			}
			row = strings.TrimRight(row, " ")
			fmt.Println(row)
		}
	case "raw":
		if files, ok := ActionForEvaluation.GetResults().([]NTypes.FileInfo); ok {
			for _, fi := range files {
				fmt.Println(fi)
			}
		} else {
			fmt.Println(ActionForEvaluation.GetResults())
		}
	default:
		fmt.Println("operation not recognized..")
		return
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

func GetTreeForInput(input string) (parser.IProgContext, error) {
	antlrInputStream := antlr.NewInputStream(input)
	lexer := parser.NewNuggetLexer(antlrInputStream)
	tokenstream := antlr.NewCommonTokenStream(lexer, 0)
	tokenparser := parser.NewNuggetParser(tokenstream)
	tokenparser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tokenparser.BuildParseTrees = true
	tree := tokenparser.Prog()

	if tree.GetText() == "" {
		return nil, errors.New("tree is empty")
	}

	return tree, nil
}

func main() {
	fmt.Println("Welcome to nugget version 0.1a")
	flagCheck()

	SetupRuntimeConnections()

	if interactiveMode {

		SetupIPCServer()

		for {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("nugget> ")

			//todo: figure out the best way to implement the tab complete
			text, err := reader.ReadString('\n')

			if text == "exit" {
				os.Exit(0)
			}

			if err != nil {
				fmt.Println("Input error: ", err)
			} else {
				resetInterruptCounter()
			}
			tree, nerr := GetTreeForInput(text)
			if nerr == nil {
				antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
				fmt.Println()
			}
		}

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
			tree, nerr := GetTreeForInput(scanner.Text())
			if nerr == nil {
				antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
				fmt.Println()
			}
		}
	}
}

func CatchTerm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println(sig.String() + " ctrl+c again to exit")
			interruptCounter += 1
			if interruptCounter >= 2 {
				fmt.Println("[!] User exit")
				os.Exit(1)
			}
		}
	}()
}

func resetInterruptCounter() {
	interruptCounter = 0
}

func SetupRuntimeConnections() {
	log.Println("[-] Setting up runtime connections .. ")
	//setup TSK
	var runtimeConnErr error
	tskClient, runtimeConnErr = rpc.DialHTTP("tcp", runtimeServer+":2001")
	if runtimeConnErr != nil {
		log.Fatal("dialing:", runtimeConnErr)
	} else {
		log.Println("[-] Connection to TSK established")
	}

	//setup VOL
	volClient, runtimeConnErr = rpc.DialHTTP("tcp", "127.0.0.1:2001")
	if runtimeConnErr != nil {
		log.Fatal("[!] Dialing Volatility Client Failed\n", runtimeConnErr)
	} else {
		log.Println("[-] Connection to Volatility established")
	}

	//setup PCAP
	log.Println("[-] All runtime connections successfully established.")
}

func SetupIPCServer() {
	lis, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	NTypes.RegisterServiceNet_FileInfoServer(server, &queryServer{})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("[!] Failed to serve: %v", err)
		}
	}()
}

type queryServer struct{}

func (s *queryServer) Get_FileInfo(ctx context.Context, in *NTypes.Net_Query) (*NTypes.Net_FileInfo, error) {

	return &NTypes.Net_FileInfo{Id: "test", Filenames:[]string{"test1.txt"}}, nil
}

func (s *queryServer) Stream_FileInfo(query *NTypes.Net_Query, stream NTypes.ServiceNet_FileInfo_Stream_FileInfoServer) error {

	var fiList []NTypes.Net_FileInfo
	fi1 := NTypes.Net_FileInfo{Id:"test1", Filenames:[]string{"test1.txt"}}
	fi2 := NTypes.Net_FileInfo{Id:"test2", Filenames:[]string{"test2.txt"}}
	fiList = append(fiList, fi1)
	fiList = append(fiList, fi2)
	for _, fi := range fiList {
		if err := stream.Send(&fi); err != nil {
			return err
		}
	}
	return nil

}
