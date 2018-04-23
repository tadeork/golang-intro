package main

import (
	"fmt"
	"os"
)

func main() {
	// the types are inferred in compiling time
	args := os.Args
	// sadly theres no ternery expressions in Go
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println("Hello, I'm Gopher")
	}
}
