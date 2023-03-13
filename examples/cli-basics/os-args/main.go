package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No Command Provided")
	}
	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "Remingers cli -cli Basics"
		f := strings.Split(os.Args[2], "=")
		if len(f) == 2 && f[0] == "--msg" {
			msg = f[1]
		}
		fmt.Printf("Hello welcome to %s \n", msg)
	case "help":
		fmt.Println("Help Message")
	default:
		fmt.Println("No Commands")
	}
}
