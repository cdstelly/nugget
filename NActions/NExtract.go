package NActions

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
	"../NTypes"
	"fmt"
)

//don't export this type so that we can force users of it to use the 'New' method, thereby initializing values
type ExtractNTFS struct {
	executed  bool
	dependsOn BaseAction

	NTFSImageLocation string
	NTFSDataRuns []NTypes.RealOffsetRun
}

func NewNExtractNTFS() ExtractNTFS {
	n := ExtractNTFS{}
	n.dependsOn = nil
	n.executed = false
	return n
}

func (na *ExtractNTFS) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractNTFS) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractNTFS) SetDependency(action BaseAction) {
	na.dependsOn = nil
}

func (na *ExtractNTFS) Exeucte() {
	na.ExtractFromNTFS()
	na.executed = true
}

func (na *ExtractNTFS) ExtractFromNTFS () []NTypes.FileInfo {
	file, _ := os.Open(na.NTFSImageLocation)

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
			fmt.Println("Won't add file ", getAFilename(fi), "  id: ", fi.Id, " -- too fragmented (over ", fraglimit, " fragments)")
		} else {
			//fileInfoMapping[fi.Id] = fi //TODO see if we can get rid of array and just keep mapping
			allfiles = append(allfiles, fi)
		}
	}

	fmt.Println("Nonfatal Error: ", errCount, " records were not parsed from NTFS MFT Record")
	return allfiles
}

// In case a file has multiple names, return the last one. Usually, the 8.3 name is at the 0 index and a 'regular' filename is at index 1
func getAFilename(f NTypes.FileInfo) string {
	if len(f.Filenames) == 0 {
		return "null"
	}
	return f.Filenames[len(f.Filenames)-1]
}