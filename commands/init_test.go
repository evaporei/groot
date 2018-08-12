package commands

import (
	"log"
	"os"
	"testing"
)

func TestCreateGrootFolderStructure(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	targetPath := currentDir + "/example_groot_structure"
	err = os.MkdirAll(targetPath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = createGrootFolderStructure(targetPath)
	if err != nil {
		log.Fatal(err)
	}

	if fileOrFolder, err := os.Stat(targetPath + "/.groot"); err == nil {
		if !fileOrFolder.IsDir() {
			t.Fail()
		}
		if head, err := os.Stat(targetPath + "/.groot/HEAD"); err == nil {
			if head.IsDir() {
				t.Fail()
			}
		}
		if objects, err := os.Stat(targetPath + "/.groot/objects"); err == nil {
			if !objects.IsDir() {
				t.Fail()
			}
		}
		if refs, err := os.Stat(targetPath + "/.groot/refs"); err == nil {
			if !refs.IsDir() {
				t.Fail()
			}
		}
		if refsHeads, err := os.Stat(targetPath + "/.groot/refs/heads"); err == nil {
			if refsHeads.IsDir() {
				t.Fail()
			}
		}
	} else {
		log.Fatal(err)
		t.Fail()
	}
}
