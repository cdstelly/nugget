package NTypes

import "time"

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