package NTypes

import "testing"
import (
	"fmt"
	"../NTypes"
)

func TestFileInfo (t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func ExampleFileInfo_GetFilename() {
	var fi NTypes.FileInfo
	fi.Filenames = []string{"index.dat"}
	fmt.Println(fi.GetFilename())
	// Output: index.dat
}