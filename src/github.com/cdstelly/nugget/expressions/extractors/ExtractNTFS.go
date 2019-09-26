package extractors

import (
	"fmt"
	"github.com/cdstelly/nugget/NTypes"
	"strconv"
	"strings"
	"time"

	"github.com/cdstelly/nugget/expressions/transforms"
	"log"
	"net/rpc"
)

type ExtractNTFS struct {
	executed  bool
	dependsOn expressions.BaseAction
	filters   []NTypes.Filter

	NTFSImageMetadataLocation string
	Location                  string

	NTFSFiles    []NTypes.FileInfo
	NTFSDataRuns []NTypes.RealOffsetRun

	TskClient *rpc.Client

	beenUploaded bool
}

func (na *ExtractNTFS) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractNTFS) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractNTFS) SetDependency(action expressions.BaseAction) {
	na.dependsOn = action
}

func (na *ExtractNTFS) Execute() {
	//fmt.Println("Executing an NTFS extraction: ", na.NTFSImageMetadataLocation)
	//na.NTFSFiles = na.ExtractMetadataFromNTFS()
	na.NTFSFiles = na.ExtractMetadataFromNTFSwithTSK()
	na.executed = true
}

func (na *ExtractNTFS) ExtractMetadataFromNTFSwithTSK() []NTypes.FileInfo {
	if na.beenUploaded == false {
		na.UploadData()
	}
	bodyFileAsStr := getBodyFileFromTSK(na.Location)
	var files []NTypes.FileInfo
	for _, entry := range strings.Split(bodyFileAsStr, "\n") {
		if len(entry) > 10 {
			if strings.Contains(entry, "($FILE_NAME)") == false {
				files = append(files, na.convertBodyFileStringToFileInfo(entry))
			}
		}
	}
	return files
}

func (na *ExtractNTFS) convertBodyFileStringToFileInfo(input string) NTypes.FileInfo {
	/*
	   MD5
	   name
	   inode
	   mode_as_string
	   UID
	   GID
	   size
	   atime
	   mtime
	   ctime
	   crtime
	*/
	theSplitLine := strings.Split(input, "|")
	var myFile NTypes.FileInfo
	myFile.Filenames = append(myFile.Filenames, theSplitLine[1])
	myFile.Id = theSplitLine[2]
	//fmt.Println("the file id: ", myFile.Id)
	myFile.Flags = theSplitLine[3]

	mytmptwo, err := strconv.Atoi(theSplitLine[6])
	myFile.Filesize = uint64(mytmptwo)
	tmpTime, err := strconv.Atoi(theSplitLine[7])
	myFile.Accesstime = time.Unix(int64(tmpTime), 0)
	tmpTime, err = strconv.Atoi(theSplitLine[8])
	myFile.Modifytime = time.Unix(int64(tmpTime), 0)
	tmpTime, err = strconv.Atoi(theSplitLine[9])
	myFile.Createtime = time.Unix(int64(tmpTime), 0)
	tmpTime, err = strconv.Atoi(theSplitLine[10])
	myFile.Emodifytime = time.Unix(int64(tmpTime), 0)

	if err != nil {
		panic(err)
	}

	//fmt.Println("the filename: " + myFile.Filenames[0] + " and the size: " + strconv.Itoa(int(myFile.Filesize)))
	return myFile
}

func (na *ExtractNTFS) UploadData() {
	//load some data into tsk memory
	args := &NTypes.NugArg{[]byte(na.Location), ""}
	var reply string
	err := na.TskClient.Call("NugTSK.LoadData", args, &reply)
	if err != nil {
		log.Fatal("tsk load error:", err)
	}
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	na.beenUploaded = true
}

func getBodyFileFromTSK(fileLocation string) string {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//load some data into tsk memory
	args := &NTypes.NugArg{[]byte(fileLocation), ""}
	var reply string
	err = client.Call("NugTSK.GetBodyFile", args, &reply)
	if err != nil {
		log.Fatal("tsk getbodyfile error:", err)
	}
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	return reply
}

func (na *ExtractNTFS) GetResults() interface{} {
	if na.executed == false {
		na.Execute()
	}
	return na.NTFSFiles
}

func checkError(err error) {
	if err != nil {
		fmt.Println("[!] Nonfatal error: ", err.Error())
	}
}

func (na *ExtractNTFS) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}
