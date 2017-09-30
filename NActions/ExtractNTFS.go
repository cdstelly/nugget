package NActions

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
	"../NTypes"
	"fmt"

	"regexp"
	"net/rpc"
	"log"
)

type ExtractNTFS struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	NTFSImageMetadataLocation string
	NTFSImageDataLocation string

	NTFSFiles []NTypes.FileInfo
	NTFSDataRuns []NTypes.RealOffsetRun

	beenUploaded bool
}

func (na *ExtractNTFS) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractNTFS) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractNTFS) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (na *ExtractNTFS) Execute() {
	fmt.Println("Executing an NTFS extraction: ", na.NTFSImageMetadataLocation)
	//na.NTFSFiles = na.ExtractMetadataFromNTFS()
	na.NTFSFiles = na.ExtractMetadataFromNTFSwithTSK()
	na.executed = true
}

func (na *ExtractNTFS) ExtractMetadataFromNTFSwithTSK() []NTypes.FileInfo {

	if na.beenUploaded == false {
		na.UploadData()
	}
	bodyFileAsStr := getBodyFileFromTSK()
	var files []NTypes.FileInfo
	for _, entry := range strings.Split(bodyFileAsStr, "\n") {
		if len(entry) > 10 {
			files = append(files, na.convertBodyFileStringToFileInfo(entry))
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
	theSplitLine := strings.Split(input,"|")
	var myFile NTypes.FileInfo
	myFile.Filenames = append(myFile.Filenames, theSplitLine[1])
	mytmp, _ := strconv.Atoi(theSplitLine[2])
	myFile.Id = uint64(mytmp)
	myFile.Flags = theSplitLine[3]

	mytmptwo, _ := strconv.Atoi(theSplitLine[6])
	myFile.Filesize = uint64(mytmptwo)
	tmpTime,_ := strconv.Atoi(theSplitLine[7])
	myFile.Accesstime = time.Unix(int64(tmpTime),0)
	tmpTime,_ = strconv.Atoi(theSplitLine[8])
	myFile.Modifytime = time.Unix(int64(tmpTime),0)
	tmpTime,_ = strconv.Atoi(theSplitLine[9])
	myFile.Createtime = time.Unix(int64(tmpTime),0)
	tmpTime,_ = strconv.Atoi(theSplitLine[10])
	myFile.Emodifytime = time.Unix(int64(tmpTime),0)

	//fmt.Println("the filename: " + myFile.Filenames[0] + " and the size: " + strconv.Itoa(int(myFile.Filesize)))
	return myFile
}

func (na *ExtractNTFS) UploadData() {
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
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	na.beenUploaded = true
}

func getBodyFileFromTSK() string {
	client, err := rpc.DialHTTP("tcp", "192.168.1.198:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//load some data into tsk memory
	args := &NugArg{[]byte("")}
	var reply string
	err = client.Call("NugTSK.GetBodyFile", args, &reply)
	if err != nil {
		log.Fatal("vol load error:", err)
	}
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	return reply
}


func (na *ExtractNTFS) ExtractMetadataFromNTFS () []NTypes.FileInfo {
	file, err := os.Open(na.NTFSImageMetadataLocation)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	defer file.Close()
	errCount := 0
	lineScanner := bufio.NewScanner(file)
	var allfiles []NTypes.FileInfo

	for lineScanner.Scan() {
		// prepare an object to store a scarftypes.FileInfo in
		var fi NTypes.FileInfo
		var dr NTypes.DataRun
		var sr NTypes.RealOffsetRun
		co_set := false
		numc_set := false

		//scanner.Text() gives us a string of the line
		onBar := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			for i := 0; i < len(data); i++ {
				if data[i] == '|' {
					return i + 1, data[:i], nil
				}
			}
			// There is one final token to be delivered, which may be the empty string.
			// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
			// but does not trigger an error to be returned from Scan itself.
			return 0, data, bufio.ErrFinalToken
		}

		barScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		barScanner.Split(onBar)
		for barScanner.Scan() {
			keyValue := strings.Split(barScanner.Text(), ":")
			if cap(keyValue) == 2 {
				//fmt.Printf("[-] %s   matches   %s\n", keyValue[0], keyValue[1])
				key := strings.TrimSpace(keyValue[0])
				value := strings.TrimSpace(keyValue[1])
				var tmpint int64
				var tmpuint uint64
				var err error
				tmpint = 1
				if key == "m" {
					//fmt.Printf("[-] MFT ID: %s\n", value)
					tmpuint, err = strconv.ParseUint(value, 10, 64)
					fi.Id = tmpuint
				}
				if key == "fn" {
					//fmt.Printf("[-] Filename: %s\n", value)
					fi.Filenames = append(fi.Filenames, value)
				}
				if key == "ct" {
					//fmt.Printf("[-] Create Time: %s\n", value)
					tmpint, err = strconv.ParseInt(value, 10, 64)
					fi.Createtime = time.Unix((tmpint/10000000)-11644473600, 0)
				}
				if key == "mt" {
					//fmt.Printf("[-] Modify Time: %s\n", value)
					tmpint, err = strconv.ParseInt(value, 10, 64)
					fi.Modifytime = time.Unix((tmpint/10000000)-11644473600, 0)
				}
				if key == "at" {
					//fmt.Printf("[-] Access Time: %s\n", value)
					tmpint, err = strconv.ParseInt(value, 10, 64)
					fi.Accesstime = time.Unix((tmpint/10000000)-11644473600, 0)
				}
				if key == "emt" {
					//fmt.Printf("[-] E Modify Time: %s\n", value)
					tmpint, err = strconv.ParseInt(value, 10, 64)
					fi.Emodifytime = time.Unix((tmpint/10000000)-11644473600, 0)
				}
				if key == "faf" {
					//fmt.Printf("[-] FA Flags: %s\n", value)
					fi.Fflags = value
				}
				if key == "ds" {
					//fmt.Printf("[-] Data Size: %s\n", value)
					tmpuint, err = strconv.ParseUint(value, 10, 64)
					fi.Filesize = tmpuint
				}
				if key == "f" {
					//fmt.Printf("[-] Flag: %s\n", value)
					fi.Flags = value
				}
				if key == "o" {
					//fmt.Printf("[-] Offset: %s\n", value)
					tmpuint, err = strconv.ParseUint(value, 10, 64)
					dr.Clusteroffset = tmpuint
					co_set = true
				}
				if key == "s" {
					//fmt.Printf("[-] Size: %s\n", value)
					tmpuint, err = strconv.ParseUint(value, 10, 64)
					dr.Numclusters = tmpuint
					numc_set = true
				}

				if numc_set == true && co_set == true {
					fi.Dataruns = append(fi.Dataruns, dr)

					if fi.Id > 0 {
						sr.FileId = fi.Id
						sr.Clusteroffset = dr.Clusteroffset
						sr.NumBytesInRun = dr.Numclusters * uint64(4096)
						na.NTFSDataRuns = append(na.NTFSDataRuns, sr) //TODO optimize this.
					}

					numc_set = false
					co_set = false
				}

				//checkError(err)
				if err != nil {
					//fmt.Printf("error parsing joachim string: %s\n", lineScanner.Text())
					//fmt.Println(err.Error())
					errCount++
				}
			}
		}

		fraglimit := 20
		if len(fi.Dataruns) > fraglimit { // revisit -- registry files are causing crashes, extremely fragmented files with lots of data runs slow things down significantly
			fmt.Println("Won't add file ", GetAFilename(fi), "  id: ", fi.Id, " -- too fragmented (over ", fraglimit, " fragments)")
		} else {
			if GetAFilename(fi) != "null" {
				if na.doesFilePassFilters(fi) {
						allfiles = append(allfiles, fi)
				} else {
					//fmt.Println("the file: ", GetAFilename(fi), "did not match our filter")
				}
			}
 		}
	}

//	fmt.Println("Nonfatal Error: ", errCount, " records were not parsed from NTFS MFT Record")
	return allfiles
}

func (na *ExtractNTFS) doesFilePassFilters(fi NTypes.FileInfo) bool {
	passes := true
	for _,filter := range na.filters {
		switch filter.Field{
		case "filename":
			if filter.Op != "==" {
				fmt.Println("Error - Operation ", filter.Op, "not supported")
			} else {
				if strings.Contains(filter.Value, "*") {
					match, _ := regexp.MatchString(filter.Value, GetAFilename(fi))
					if !match {
						return false
					}
				} else {
					match := strings.Compare(filter.Value, GetAFilename(fi))
					if match != 0 {
						//did not match exactly
						return false
					}
				}
			}
		case "ctime":
			theTime := strings.Trim(filter.Value, "\"")
			layout := "01/02/06"
			t, err := time.Parse(layout, theTime)
			checkError(err)
			//fmt.Println("datetime of filter:", t.String(), " datetime of file: ", fi.Createtime.String())
			switch(filter.Op) {
			case ">":
				return t.Before(fi.Createtime)
			case ">=":
				return t.Equal(fi.Createtime) || t.Before(fi.Createtime)
			case "<=":
				return t.Equal(fi.Createtime) || t.After(fi.Createtime)
			case "<":
				return t.After(fi.Createtime)
			case "==":
				d := fi.Createtime.String()
				dt, err := time.Parse(layout,d)
				checkError(err)
				return t.Equal(dt)
			default:
				fmt.Println("Error - operation", filter.Op, " not recognized")
			}
		case "mtime":
			theTime := strings.Trim(filter.Value, "\"")
			layout := "01/02/06"
			t, err := time.Parse(layout, theTime)
			checkError(err)
			switch(filter.Op) {
			case ">":
				return t.Before(fi.Modifytime)
			case ">=":
				return t.Equal(fi.Modifytime) || t.Before(fi.Modifytime)
			case "<=":
				return t.Equal(fi.Modifytime) || t.After(fi.Modifytime)
			case "<":
				return t.After(fi.Modifytime)
			case "==":
				d := fi.Modifytime.String()
				dt, err := time.Parse(layout,d)
				checkError(err)
				return t.Equal(dt)
			default:
				fmt.Println("Error - operation", filter.Op, " not recognized")
			}
		default:
			fmt.Println("Error -- filter field not supported by FileInfo")
		}
	}
	return passes
}

// In case a file has multiple names, return the last one. Usually, the 8.3 name is at the 0 index and a 'regular' filename is at index 1
func GetAFilename(f NTypes.FileInfo) string {
	if len(f.Filenames) == 0 {
		return "null"
	}
	return f.Filenames[len(f.Filenames)-1]
}

func (na *ExtractNTFS) GetResults() interface{}{
	if na.executed == false {
		na.Execute()
	}
	return na.NTFSFiles
}

func readFile(imageFD *os.File, info NTypes.FileInfo) (int) {
	bytesRemaining := info.Filesize
	totalBytesRead := 0
	fileData := make([]byte, 0)

	for _, dr := range info.Dataruns {
		numBytesToRead := uint64(4096) * dr.Numclusters
		if numBytesToRead > bytesRemaining {
			numBytesToRead = bytesRemaining
		}
		chunkData := make([]byte, numBytesToRead)
		numRead, err := imageFD.ReadAt(chunkData,int64(4096)*int64(dr.Clusteroffset))
		checkError(err)
		bytesRemaining -= uint64(numRead)
		totalBytesRead += numRead

		fileData = append(fileData, chunkData...)
		checkError(err)
	}
	if int64(totalBytesRead) != int64(info.Filesize) {
		fmt.Println("Did not read the expect number of bytes for file: ", GetAFilename(info), "\t ID: ", info.Id)
		// maybe a sparse file if we get here?, maybe rewriting filelen instead of dealing with sparse will work for now
	}
	info.ReconstructedData = fileData
	return totalBytesRead
}

func (na *ExtractNTFS) streamImage() {
	myfiles := na.NTFSFiles
	fmt.Println("Found ", len(myfiles), " files")

	fmt.Println("Began streaming data from SSD..")

	ntfsimage, err := os.Open(na.NTFSImageDataLocation)
	checkError(err)
	totalBytesRead := 0
	beginTime := time.Now()

	defer ntfsimage.Close()
	// for every file in metadata,
	for idx, file := range myfiles {
		numRead := readFile(ntfsimage, file)
		fmt.Println("file info index:\t", idx, "\tfilename: ", GetAFilename(file))
		totalBytesRead += numRead
	}
	elapsed := time.Since(beginTime).Seconds()
	fmt.Println("Finished streaming data from SSD. Read ", totalBytesRead, " bytes in ", elapsed, " seconds. ", (float64(totalBytesRead)/1048576)/(elapsed), " (MB/s)")
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
