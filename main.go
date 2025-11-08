package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/VictorHRRios/gordle/internal/repl"
	"github.com/fatih/color"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	color.Cyan("Hello %s! This is the Gordle REPL!\n", user.Username)
	color.Cyan("Type 'help' to see the available command")
	fmt.Printf("\n")
	repl.Start(os.Stdin, os.Stdout)
}
