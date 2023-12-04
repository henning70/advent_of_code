package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "unicode"
)

var (
    digit  []string
    digits []int64
)

func main() {
    fmt.Println("Advent Of Code, Day 1 - Part 1")

    f, err := os.Open("data.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()

    f_scanner := bufio.NewScanner(f)
    f_scanner.Split(bufio.ScanLines)
    for f_scanner.Scan() {
        fmt.Println("\nProcessing:", f_scanner.Text())
        for _, c := range f_scanner.Text() {
            if unicode.IsDigit(c) {
                fmt.Println("Digit found:", string(c))
                digit = append(digit, string(c))
            }
        }
        final_digit := fmt.Sprintf("%s%s", digit[0], digit[len(digit)-1])
        digit = nil
        final_digit_int, _ := strconv.ParseInt(final_digit, 10, 0)
        digits = append(digits, final_digit_int)
    }
    fmt.Println(digits)
    sum := 0
    for _, d := range digits {
        sum += int(d)
    }
    fmt.Println("\nSum:", sum)
}
