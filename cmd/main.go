package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Daviddlh1/writing-an-interpreter-in-go/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	fmt.Printf("Type 'exit' to quit the REPL\n")
	repl.Start(os.Stdin, os.Stdout)
}
