package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(`
    Enter following commands to control:
    q -- quit
    e -- echo input
`)

    r := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Enter command-> ")

        rawLine, _, _ := r.ReadLine()

        line := string(rawLine)

        if line == "q" {
            break
        }

        tokens := strings.Split(line, " ")

        fmt.Println("input tokens:", tokens)
    }
}

