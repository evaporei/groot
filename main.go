package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("I am groot")

	err := runCli()

	if err != nil {
		fmt.Println(err)
	}
}

func runCli() (err error) {
	if len(os.Args) > 1 {
		gitCommand := os.Args[1]
		gitCommandArgs := os.Args[2:]

		fmt.Println(gitCommand)
		fmt.Println(gitCommandArgs)

		switch gitCommand {
		default:
			stringErr := fmt.Sprintf("%s command is not implemented", gitCommand)
			return errors.New(stringErr)
		}
	}

	return errors.New("you must specify a Git command to run!")
}
