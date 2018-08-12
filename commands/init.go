package commands

import (
	"fmt"
	"os"

	"github.com/otaviopace/groot/files"
)

func Init(args []string) (err error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if files.IsARepoAlready(currentDir) {
		fmt.Println("already groot repository, nothing changed")
		return nil
	}

	err = createGrootFolderStructure(currentDir)
	if err != nil {
		return err
	}

	return nil
}

func createGrootFolderStructure(basePath string) (err error) {
	err = os.MkdirAll(basePath+"/.groot", 0777)
	if err != nil {
		return err
	}

	head, err := os.OpenFile(basePath+"/.groot/HEAD", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer head.Close()

	err = os.MkdirAll(basePath+"/.groot/objects", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll(basePath+"/.groot/refs", 0777)
	if err != nil {
		return err
	}

	refsHeads, err := os.OpenFile(basePath+"/.groot/refs/heads", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer refsHeads.Close()

	return nil
}
