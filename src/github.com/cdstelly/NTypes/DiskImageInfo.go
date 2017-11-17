package NTypes

import "fmt"

type DiskImageInfo struct {
	FSType string
	SerialNumber string
	OEMName string
	Version string
	FirstCluster string
	SectorSize string
	ClusterSize	 string
	ClusterRange string
	SectorRange	string
}

type OffsetInfo struct {
	OffsetBytes int
	ClusterSize int
}

func (di DiskImageInfo) String() string {
	var	myString string
	myString = fmt.Sprintf("Version: %s\n", di.Version)
	myString = myString + fmt.Sprintf("FSType: %s\n", di.FSType)
	myString = myString + fmt.Sprintf("Serial Number: %s\n",  di.SerialNumber)
	myString = myString + fmt.Sprintf("Cluster Range: %s\n", di.ClusterRange)
	return myString
}