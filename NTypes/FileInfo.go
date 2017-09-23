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