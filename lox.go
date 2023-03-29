package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    fmt.Println("Lox Interpreter")

    runPrompt()
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
