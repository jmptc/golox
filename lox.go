package main

import (
	"bufio"
	"fmt"
	"os"

    "github.com/jmptc/golox/token"
)

var hadError bool

func main() {
	fmt.Println("Lox Interpreter")

	// runPrompt()

    t := token.Token{ TokenType: "DOT", Lexeme: ".", Line: 0}
    fmt.Println(t)
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
