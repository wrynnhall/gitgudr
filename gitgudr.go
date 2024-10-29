package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	wiprCmd := flag.NewFlagSet("wipr", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Expected 'wipr' subcommand")
	}

	switch os.Args[1] {
	case "wipr":
		fmt.Println("wipr")
		fmt.Println("tail", wiprCmd.Args())
	default:
		os.Exit(1)
	}
}
