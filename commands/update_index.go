package commands

import (
	"errors"
	"os"
	"strings"

	"github.com/otaviopace/groot/files"
)

func UpdateIndex(args []string) (err error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if !files.IsARepoAlready(currentDir) {
		return errors.New("not in groot repository, try using the `groot init` command first")
	}

	path, options := parseArgs(args)

	fileOrFolder, err := os.Stat(path)

	var isOnDisk bool

	if err == nil {
		isOnDisk = true
	} else {
		isOnDisk = false
	}

	isInIndex := index.IsInIndex(path)

	if isOnDisk && fileOrFolder.IsDir() {
		return errors.New(path + " is a directory - add files inside")
	}

	if options["remove"] && !isOnDisk && isInIndex {
		if index.IsFileInConflict(path) {
			return errors.New("unsupported")
		}

		return index.WriteRm(path)
	}

	if options["remove"] && !isOnDisk && isInIndex {
		return nil
	}

	if !options["add"] && isOnDisk && !isInIndex {
		return errors.New("cannot add " + path + " to index - use --add option")
	}

	if isOnDisk && (options["add"] || isInIndex) {
		return index.WriteNonConflict(path, files.Read(files.workingCopyPath(path)))
	}

	if !options["remove"] && !isOnDisk {
		return errors.New(path + " does not exist and --remove not passed")
	}

	return nil
}

func parseArgs(args []string) (string, map[string]bool) {
	path := args[0]

	options := make(map[string]bool)

	for _, option := range args[1:] {
		if strings.HasPrefix(option, "--") {
			options[option[2:]] = true
		} else if strings.HasPrefix(option, "-") {
			options[option[1:]] = true
		}
	}

	return path, options
}
