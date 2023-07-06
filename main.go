package main

import (
	"fmt"
	"github.com/AYGA2K/interpreter_go/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the AyScript programming language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
