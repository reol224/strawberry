package main

import (
	"fmt"
	"os"
	"os/user"
	"strawberryInterpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Strawberry programming language!\n",
		user.Name)
	fmt.Printf("Functionality is reduced for now but stay tuned for future updates!\n")
	repl.Start(os.Stdin, os.Stdout)

}
