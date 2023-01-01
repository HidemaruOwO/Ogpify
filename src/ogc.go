package main

import (
	"fmt"
	"os"

	"github.com/HidemaruOwO/OGP-Generate-API/src/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
