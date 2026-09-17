package main

import (
	"fmt"
	"lipi/repl"
	"os"
	"os/user"
)

// func main() {
// 	user, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Hello %s! This is the Lipi programming language!\n", user.Username)
// 	fmt.Printf("Feel free to type in commands\n")
// 	repl.Start(os.Stdin, os.Stdout)
// }
func main() {
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
}