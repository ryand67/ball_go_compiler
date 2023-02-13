package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ryand67/ball_go_compiler/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
