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

func (di DiskImageInfo) String() string {
	var	myString string
	myString = fmt.Sprintf("Version: %s\n", di.Version)
	myString = myString + fmt.Sprintf("FSType: %s\n", di.FSType)
	myString = myString + fmt.Sprintf("Serial Number: %s\n",  di.SerialNumber)
	myString = myString + fmt.Sprintf("Cluster Range: %s\n", di.ClusterRange)
	return myString
}
/*FILE SYSTEM INFORMATION
--------------------------------------------
File System Type: NTFS
Volume Serial Number: 72582A3D582A0109
OEM Name: NTFS
Version: Windows XP

METADATA INFORMATION
--------------------------------------------
First Cluster of MFT: 786432
First Cluster of MFT Mirror: 1876588
Size of MFT Entries: 1024 bytes
Size of Index Records: 4096 bytes
Range: 0 - 29632
Root Directory: 5

CONTENT INFORMATION
--------------------------------------------
Sector Size: 512
Cluster Size: 4096
Total Cluster Range: 0 - 3753176
Total Sector Range: 0 - 30025420

$AttrDef Attribute Values:
$STANDARD_INFORMATION (16)   Size: 48-72   Flags: Resident
$ATTRIBUTE_LIST (32)   Size: No Limit   Flags: Non-resident
$FILE_NAME (48)   Size: 68-578   Flags: Resident,Index
$OBJECT_ID (64)   Size: 0-256   Flags: Resident
$SECURITY_DESCRIPTOR (80)   Size: No Limit   Flags: Non-resident
$VOLUME_NAME (96)   Size: 2-256   Flags: Resident
$VOLUME_INFORMATION (112)   Size: 12-12   Flags: Resident
$DATA (128)   Size: No Limit   Flags:
$INDEX_ROOT (144)   Size: No Limit   Flags: Resident
$INDEX_ALLOCATION (160)   Size: No Limit   Flags: Non-resident
$BITMAP (176)   Size: No Limit   Flags: Non-resident
$REPARSE_POINT (192)   Size: 0-16384   Flags: Non-resident
$EA_INFORMATION (208)   Size: 8-8   Flags: Resident
$EA (224)   Size: 0-65536   Flags:
$LOGGED_UTILITY_STREAM (256)   Size: 0-65536   Flags: Non-resident
*/