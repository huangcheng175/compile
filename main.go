package main

import (
	"compile/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	println("Hello " + user.Username + "!")
	println("This is the Monkey programming language!")
	println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
