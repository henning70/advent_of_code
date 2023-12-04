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

    games_possible     = make(map[int64]int64)
    games_not_possible []int64
)

func main() {
    fmt.Println("Advent Of Code, Day 2 - Part 1")

    // fill the map
    game_bag["red_cubes"] = 12
    game_bag["green_cubes"] = 13
    game_bag["blue_cubes"] = 14

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

    for f_scanner.Scan() {
        in_game := f_scanner.Text()
        fmt.Println("\nProcessing:", in_game)

        game_id := strings.Replace(strings.Split(in_game, ":")[0], "Game ", "", -1)
        game_id_int, _ := strconv.ParseInt(game_id, 10, 0)
        game_sets := strings.Split(in_game, ":")[1]

        fmt.Println("Found Game with id", game_id)
        fmt.Println("Found sets result", game_sets)

        i := 1
        for _, in_sets := range strings.Split(game_sets, ";") {
            game_result := make(map[string]int64)
            fmt.Println("Processing set", i, "with result", string(in_sets))
            i++
            for _, r := range strings.Split(in_sets, ",") {
                fmt.Println("Result:", string(r))
                if strings.Contains(string(r), "red") {
                    value, _ := strconv.ParseInt(strings.Split(string(r), " ")[1], 10, 0)
                    fmt.Println("Red Value:", value)
                    game_result["red"] = value + game_result["red"]
                } else if strings.Contains(string(r), "green") {
                    value, _ := strconv.ParseInt(strings.Split(string(r), " ")[1], 10, 0)
                    fmt.Println("Green Value:", value)
                    game_result["green"] = value + game_result["green"]
                } else if strings.Contains(string(r), "blue") {
                    value, _ := strconv.ParseInt(strings.Split(string(r), " ")[1], 10, 0)
                    fmt.Println("Blue Value:", value)
                    game_result["blue"] = value + game_result["blue"]
                }
            }
            fmt.Println("Game result:", game_result)
            fmt.Println("Game bag:", game_bag)
            fmt.Println("Game id:", game_id)

            if game_result["red"] > game_bag["red_cubes"] || game_result["green"] > game_bag["green_cubes"] || game_result["blue"] > game_bag["blue_cubes"] {
                fmt.Println("red cubes exceeded, game with id", game_id, "not possible")
                games_not_possible = append(games_not_possible, game_id_int)
            } else {
                fmt.Println("Game with id", game_id, "possible")
                games_possible[game_id_int] = game_id_int
            }
        }
        i = 1
    }

    for _, game_id_int := range games_not_possible {
        delete(games_possible, game_id_int)
    }

    fmt.Println("games possible", games_possible)
    fmt.Println("games not possible", games_not_possible)

    sum := 0
    for _, d := range games_possible {
        sum += int(games_possible[d])
    }
    fmt.Println("\nSum possible:", sum)

    sum = 0
    for _, d := range games_not_possible {
        sum += int(d)
    }
    fmt.Println("\nSum not possible:", sum)
}
