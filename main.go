// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"seahorsesyntax.villdyr.dev/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Seahorse programming language!\n", user.Username)
	fmt.Printf("Type the Seahorse commands that your heart desires.\n")
	repl.Start(os.Stdin, os.Stdout)
}
