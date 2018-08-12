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

	return nil
}
