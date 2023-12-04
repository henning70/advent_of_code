package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    line string
)

func main() {
    fmt.Println("Advent Of Code, Day 3 - Part 1")

    f, err := os.Open("test_data.txt")
    // f, err := os.Open("data.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()

    f_scanner := bufio.NewScanner(f)
    f_scanner.Split(bufio.ScanLines)

    for f_scanner.Scan() {
        line = f_scanner.Text()
        fmt.Println("\nProcessing:", line)
    }
}
