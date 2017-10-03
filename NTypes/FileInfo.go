package NTypes

import (
	"time"
	"strings"
	"net/rpc"
	"log"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"regexp"
)

type FileInfo struct {
	Id          string 		//tsk represents as string w/ dashes
	Filenames   []string
	Createtime  time.Time
	Modifytime  time.Time
	Accesstime  time.Time
	Emodifytime time.Time
	Fflags      string
	Flags       string
	Filesize    uint64
	Dataruns    []DataRun
	reconstructedData []byte
	beenReconstructed bool
}

func getFileFromTSK(inode string) []byte {
	client, err := rpc.DialHTTP("tcp", "192.168.1.198:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//load some data into tsk memory
	args := &NugArg{[]byte(""),inode}
	var reply string
	err = client.Call("NugTSK.GetFileData", args, &reply)
	if err != nil {
		log.Fatal("tsk get file data error:", err)
	}
	decodedreply, err := base64.StdEncoding.DecodeString(reply)
	if err != nil {
		panic(err)
	}
	hasher := md5.New()
	hasher.Write(decodedreply)
//	myhash := fmt.Sprintf("%x", hasher.Sum(nil))

	return decodedreply
}

func (fi *FileInfo) ReconstructTheData() {
	fi.beenReconstructed = true
	fi.reconstructedData = getFileFromTSK(fi.Id)
}

func (fi *FileInfo) GetFileData() []byte {
	if fi.beenReconstructed == false {
		fi.ReconstructTheData()
	}
	return fi.reconstructedData
}

func (fi *FileInfo) SetFileData(data []byte) {
	fi.reconstructedData = data
}

func (fi *FileInfo) DoesPassFilter(theFilters []Filter) bool {
	passes := true
	for _,filter := range theFilters {
		switch filter.Field{
		case "filename":
			if filter.Op != "==" {
				fmt.Println("Error - Operation ", filter.Op, "not supported")
			} else {
				if strings.Contains(filter.Value, "*") {
					filterStr := strings.TrimSpace(filter.Value)
					filterStr = strings.Trim(filterStr, `"`)
					match, _ := regexp.MatchString(filterStr, fi.Filenames[0])
					if !match {
						//fmt.Println("Target: ", fi.Filenames[0], " did not match Fitler: ", filterStr)
						return false
					}
				} else {
					match := strings.Compare(filter.Value, fi.Filenames[0])
					if match != 0 {
						//did not match exactly
						return false
					}
				}
			}
		case "ctime":
			theTime := strings.Trim(filter.Value, "\"")
			layout := "01/02/2006"
			t, err := time.Parse(layout, theTime)
			if err != nil {
				fmt.Errorf("Error! %s", err)
			}
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
				if err != nil {
					fmt.Errorf("Error! %s", err)
				}
				return t.Equal(dt)
			default:
				fmt.Println("Error - operation", filter.Op, " not recognized")
			}
		case "mtime":
			theTime := strings.Trim(filter.Value, "\"")
			layout := "01/02/2006"
			t, err := time.Parse(layout, theTime)
			if err != nil {
				fmt.Errorf("Error! %s", err)
			}
			switch(filter.Op) {
			case "<":
				return t.Before(fi.Modifytime)
			case "<=":
				return t.Equal(fi.Modifytime) || t.Before(fi.Modifytime)
			case ">=":
				return t.Equal(fi.Modifytime) || t.After(fi.Modifytime)
			case ">":
				return t.After(fi.Modifytime)
			case "==":
				d := fi.Modifytime.String()
				dt, err := time.Parse(layout,d)
				if err != nil {
					fmt.Errorf("Error! %s", err)
				}
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

type DataRun struct {
	Runid         int
	Clusteroffset uint64
	Numclusters   uint64
}

type RealOffsetRun struct {
	FileId        string
	NumBytesInRun uint64
	Clusteroffset uint64
}

type ByOffset []RealOffsetRun

func (cbo ByOffset) Len() int {
	return len(cbo)
}

func (cbo ByOffset) Swap(i, j int) {
	cbo[i], cbo[j] = cbo[j], cbo[i]
}

func (cbo ByOffset) Less(i, j int) bool {
	return cbo[i].Clusteroffset < cbo[j].Clusteroffset
}

// Sorting interfaces
type ByCtime []FileInfo
func (a ByCtime) Len() int           { return len(a) }
func (a ByCtime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCtime) Less(i, j int) bool { return a[i].Createtime.Before(a[j].Createtime) }
type ByMtime []FileInfo
func (a ByMtime) Len() int           { return len(a) }
func (a ByMtime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMtime) Less(i, j int) bool { return a[i].Modifytime.Before(a[j].Modifytime) }
type ByAtime []FileInfo
func (a ByAtime) Len() int           { return len(a) }
func (a ByAtime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAtime) Less(i, j int) bool { return a[i].Accesstime.Before(a[j].Accesstime) }
type ByEtime []FileInfo
func (a ByEtime) Len() int           { return len(a) }
func (a ByEtime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEtime) Less(i, j int) bool { return a[i].Emodifytime.Before(a[j].Emodifytime) }
type ByFilesize []FileInfo
func (a ByFilesize) Len() int           { return len(a) }
func (a ByFilesize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFilesize) Less(i, j int) bool { return a[i].Filesize < a[j].Filesize }

type ByFilename []FileInfo
func (a ByFilename) Len() int           { return len(a) }
func (a ByFilename) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFilename) Less(i, j int) bool {
	if len(a[i].Filenames) == 0 { return false }
	if len(a[j].Filenames) == 0 { return false }
	return strings.Compare(a[i].Filenames[0], a[j].Filenames[0]) == -1
}