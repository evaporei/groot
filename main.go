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
		grootCommand := os.Args[1]
		grootCommandArgs := os.Args[2:]

		fmt.Println(grootCommand)
		fmt.Println(grootCommandArgs)

		switch grootCommand {
		default:
			stringErr := fmt.Sprintf("%s command is not implemented", grootCommand)
			return errors.New(stringErr)
		}
	}

	return errors.New("you must specify a groot command to run!")
}
