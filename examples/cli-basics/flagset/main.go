package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No Command Provided")
		os.Exit(2)
	}

	cmd := os.Args[1]

	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI-Basics", "greeted cmd message")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Welcome %s\n", *msgFlag)
	case "help":
		fmt.Println("Help Message")
	default:
		fmt.Printf("Unknown Command %s\n", cmd)
	}
}
