package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"lipi/repl"
	"lipi/runner"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		currentUser, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Println(`
██╗     ██╗██████╗ ██╗
██║     ██║██╔══██╗██║
██║     ██║██████╔╝██║
██║     ██║██╔═══╝ ██║
███████╗██║██║     ██║
╚══════╝╚═╝╚═╝     ╚═╝
`)

		fmt.Printf("Hello %s! This is the Lipi programming language!\n", currentUser.Username)
		fmt.Println("Type 'exit' to quit.")
		fmt.Println()

		repl.Start(os.Stdin, os.Stdout, currentUser.Username)
		return
	}

	if len(args) == 2 {
		filename := args[1]

		if filepath.Ext(filename) != ".lipi" {
			fmt.Fprintln(os.Stderr, "Error: file must have .lipi extension")
			os.Exit(1)
		}

		if err := runner.RunFile(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		return
	}

	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  lipi")
	fmt.Fprintln(os.Stderr, "  lipi file.lipi")
	os.Exit(1)
}