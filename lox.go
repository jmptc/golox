package main

import (
	"bufio"
	"fmt"
	"os"

    "github.com/jmptc/golox/scanner"
)

var hadError bool

func main() {
	fmt.Println("Lox Interpreter")

	// runPrompt()
    
    scanner := scanner.NewScanner("hello world")
    fmt.Println(scanner)
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(line)
}
