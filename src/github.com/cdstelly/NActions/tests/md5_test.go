package NActions

import (
	"testing"
	"../NActions"
	"../NTypes"
	"strings"
)

func TestPerformHashing(t *testing.T) {

	var f NTypes.FileInfo
	f.SetFileData([]byte("abc"))
	files := []NTypes.FileInfo{f}

	var ma NActions.MD5Action
	ma.OperateOn = files
	result := "900150983CD24FB0D6963F7D28E17F72"

	ma.PerformHashing()

	for _,x := range ma.Results {
		x.Digest = strings.ToUpper(x.Digest)
		result = strings.ToUpper(result)
		if x.Digest != result {
			t.Error(
				"For", f,
				"expected", result,
				"got", x.Digest,
			)
		}
	}
}