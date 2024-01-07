package main

import (
	"bufio"
	"fmt"
	"log"
	"lox/internal/scanner"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: lox [script]")
	} else if len(args) == 1 {
		runScript(args[0])
	} else {
		runPrompt()
	}
}

func runScript(scriptPath string) {
	fmt.Println("running", scriptPath)
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		var line string
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalln("Failed to read the line", err)
		}
		s := scanner.NewScanner(line)
		fmt.Println(s.ScanTokens())
	}
}
