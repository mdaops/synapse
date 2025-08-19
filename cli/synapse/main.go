package main

import (
	"fmt"
	"os"
	"synapse/cmd"
)

func main() {
	root := cmd.NewRootCMD()
	if err := root.CMD.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
