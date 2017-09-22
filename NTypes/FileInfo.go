package NTypes

import (
	"time"
	"strings"
)

type FileInfo struct {
	Id          uint64
	Filenames   []string
	Createtime  time.Time
	Modifytime  time.Time
	Accesstime  time.Time
	Emodifytime time.Time
	Fflags      string
	Flags       string
	Filesize    uint64
	Dataruns    []DataRun
	ReconstructedData []byte
}

type DataRun struct {
	Runid         int
	Clusteroffset uint64
	Numclusters   uint64
}

type RealOffsetRun struct {
	FileId        uint64
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

// FileInfo implements sort interface
type ByCtime []FileInfo

func (a ByCtime) Len() int           { return len(a) }
func (a ByCtime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCtime) Less(i, j int) bool { return a[i].Createtime.Before(a[j].Createtime) }


// FileInfo implements sort interface
type ByFilename []FileInfo

func (a ByFilename) Len() int           { return len(a) }
func (a ByFilename) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFilename) Less(i, j int) bool {
	if a[i].Filenames == nil { return false }
	if a[j].Filenames == nil { return false }
	return strings.Compare(a[i].Filenames[0], a[j].Filenames[0]) == -1
}