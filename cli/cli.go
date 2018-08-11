package cli

import (
	"errors"
	"fmt"

	"github.com/otaviopace/groot/commands"
)

func Run(args []string) (err error) {
	if len(args) > 1 {
		grootCommand := args[1]
		grootCommandArgs := args[2:]

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
