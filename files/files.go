package files

import (
	"os"
)

func IsARepoAlready(dirPath string) bool {
	if fileOrFolder, err := os.Stat(dirPath + "/.groot"); err == nil {
		if fileOrFolder.IsDir() {
			return true
		}
	}
	return false
}
