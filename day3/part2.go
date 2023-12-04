package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var (
    game_id     string
    game_id_int int64
    game_bag    = make(map[string]int64)

    game_power []int64
)

func main() {
    fmt.Println("Advent Of Code, Day 2 - Part 1")

    fmt.Println(game_bag)

    // f, err := os.Open("test_data.txt")
    f, err := os.Open("data.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()

    f_scanner := bufio.NewScanner(f)
    f_scanner.Split(bufio.ScanLines)

    game_result := make(map[string]int64)
    game_result["red"] = 0
    game_result["green"] = 0
    game_result["blue"] = 0
    for f_scanner.Scan() {
        in_game := f_scanner.Text()
        fmt.Println("\nProcessing:", in_game)

        game_id = strings.Replace(strings.Split(in_game, ":")[0], "Game ", "", -1)
        // game_id_int, _ := strconv.ParseInt(game_id, 10, 0)
        game_sets := strings.Split(in_game, ":")[1]

        fmt.Println("Found Game with id", game_id)
        fmt.Println("Found sets result", game_sets)

        i := 1
        for _, in_sets := range strings.Split(game_sets, ";") {
            fmt.Println("Processing set", i, "with result", string(in_sets))
            i++
            for _, r := range strings.Split(in_sets, ",") {
                fmt.Println("Result:", string(r))
                if strings.Contains(string(r), "red") {
                    value, _ := strconv.ParseInt(strings.Split(string(r), " ")[1], 10, 0)
                    fmt.Println("Red Value:", value)
                    if value > game_result["red"] {
                        game_result["red"] = value
                    }
                } else if strings.Contains(string(r), "green") {
                    value, _ := strconv.ParseInt(strings.Split(string(r), " ")[1], 10, 0)
                    fmt.Println("Green Value:", value)
                    if value > game_result["green"] {
                        game_result["green"] = value
                    }
                } else if strings.Contains(string(r), "blue") {
                    value, _ := strconv.ParseInt(strings.Split(string(r), " ")[1], 10, 0)
                    fmt.Println("Blue Value:", value)
                    if value > game_result["blue"] {
                        game_result["blue"] = value
                    }
                }
            }
        }
        i = 1
        fmt.Println("Game id:", game_id)
        fmt.Println("Game result:", game_result, "for id", game_id)

        // get the power value
        power := game_result["red"] * game_result["green"] * game_result["blue"]

        fmt.Println("Game power:", power)
        game_power = append(game_power, power)

        // clear game_result
        game_result["red"] = 0
        game_result["green"] = 0
        game_result["blue"] = 0
    }
    fmt.Println("Game power:", game_power)

    sum := 0
    for _, d := range game_power {
        sum += int(d)
    }
    fmt.Println("\nSum game power:", sum)
}
