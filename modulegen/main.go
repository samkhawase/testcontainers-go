package main

import (
	"fmt"
	"os"

	"github.com/samkhawase/testcontainers-go/modulegen/cmd"
)

func main() {
	err := cmd.NewRootCmd.Execute()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
