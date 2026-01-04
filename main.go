package main

import (
	"bat-go/repl"
	"fmt"
	"os"
	"os/user"
)

//     _______  ________  ________  _______   ________
//   ╱╱   ╱   ╲╱    ╱   ╲╱    ╱   ╲╱       ╲╲╱    ╱   ╲
//  ╱╱        ╱         ╱_       _╱        ╱╱_       _╱
// ╱         ╱╲__     ╱╱         ╱         ╱         ╱
// ╲__╱_____╱   ╲____╱╱╲___╱___╱╱╲________╱╲___╱___╱╱

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Bat programming language!\n", user.Username)
	fmt.Printf("REPL: Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
