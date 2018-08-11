package cli

import (
	"errors"
	"fmt"

	"github.com/otaviopace/groot/commands"
)

const minimumArgsLength = 1

func areArgsFilled(args []string) bool {
	return len(args) > minimumArgsLength
}

func parseArgs(args []string) (string, []string) {
	command := args[1]
	commandArgs := args[2:]

	return command, commandArgs
}

func Run(args []string) (err error) {
	if areArgsFilled(args) {
		grootCommand, grootCommandArgs := parseArgs(args)

		fmt.Println(grootCommand)
		fmt.Println(grootCommandArgs)

		switch grootCommand {
		case "init":
			return commands.Init(grootCommandArgs)
		default:
			stringErr := fmt.Sprintf("%s command is not implemented", grootCommand)
			return errors.New(stringErr)
		}
	}

	return errors.New("you must specify a groot command to run!")
}
