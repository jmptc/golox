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
	text := "var x = 12.34; var id = \"Hello world\"" 
	fmt.Println("text: ", text)
	scanner := scanner.NewScanner(text)
	toks := scanner.ScanTokens()
	for i, tok := range toks {
		fmt.Println(i, tok)
	}
	fmt.Println(toks)
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
