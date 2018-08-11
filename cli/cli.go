package cli

import (
	"errors"
	"fmt"
)

func Run(args []string) (err error) {
	if len(args) > 1 {
		grootCommand := args[1]
		grootCommandArgs := args[2:]

		fmt.Println(grootCommand)
		fmt.Println(grootCommandArgs)

		switch grootCommand {
		case "init":
			return initCommand(grootCommandArgs)
		default:
			stringErr := fmt.Sprintf("%s command is not implemented", grootCommand)
			return errors.New(stringErr)
		}
	}

	return errors.New("you must specify a groot command to run!")
}

func initCommand(args []string) (err error) {
	fmt.Println("yo init yo")
	return nil
}
