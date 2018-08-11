package main

import (
	"fmt"
	"os"

	"github.com/otaviopace/groot/cli"
)

func main() {
	fmt.Println("I am groot")

	err := cli.Run(os.Args)

	if err != nil {
		fmt.Println(err)
	}
}
