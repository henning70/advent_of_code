package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"

    "golang.org/x/exp/slices"
)

var (
    digit  []string
    digits []int64

    numbers         = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
    str_numbers     = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    str_to_num      = make(map[string]string)
    index_of_digits = make(map[int]int64)
)

func main() {
    fmt.Println("Advent Of Code, Day 1 - Part 2")

    str_to_num["one"] = "1"
    str_to_num["two"] = "2"
    str_to_num["three"] = "3"
    str_to_num["four"] = "4"
    str_to_num["five"] = "5"
    str_to_num["six"] = "6"
    str_to_num["seven"] = "7"
    str_to_num["eight"] = "8"
    str_to_num["nine"] = "9"

    f, err := os.Open("data.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()

    f_scanner := bufio.NewScanner(f)
    f_scanner.Split(bufio.ScanLines)
    for f_scanner.Scan() {
        entry := f_scanner.Text()
        fmt.Println("\nProcessing entry:", entry)

        // first we will find the index of the word digits and put them in a map
        for _, n := range str_numbers {
            if strings.Contains(entry, n) {
                entry_cnt := strings.Count(entry, n)

                for i := 0; i < entry_cnt; i++ {
                    first_index := strings.Index(entry, n)
                    last_index := strings.LastIndex(entry, n)
                    // fmt.Println(n, "found with count", entry_cnt, "with index", first_index)
                    // fmt.Println(n, "found with count", entry_cnt, "with index", last_index)
                    index_of_digits[first_index], _ = strconv.ParseInt(str_to_num[n], 10, 0)
                    index_of_digits[last_index], _ = strconv.ParseInt(str_to_num[n], 10, 0)
                }
            }
        }

        // next we will find the index of the digits and put them in a map
        for n := 0; n < len(entry); n++ {
            str_entry := string(entry[n])
            // fmt.Println("substr is ", str_entry, "with index", n)
            if slices.Contains(numbers, str_entry) {
                // fmt.Println("Digit found", string(entry[n]), "and index is", n)
                index_of_digits[n], _ = strconv.ParseInt(str_entry, 10, 0)
            }
        }
        fmt.Println("word digits:", index_of_digits)
        fmt.Println("word digits and digits: ", index_of_digits)
        fmt.Println("len: ", len(index_of_digits))

        keys := make([]int, 0, len(index_of_digits))
        for key := range index_of_digits {
            keys = append(keys, key)
        }
        sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
        // fmt.Println("keys: ", keys, keys[0], keys[len(keys)-1])
        // fmt.Println(index_of_digits[keys[0]], index_of_digits[keys[len(keys)-1]])

        final_digit := fmt.Sprintf("%v%v", index_of_digits[keys[0]], index_of_digits[keys[len(keys)-1]])
        fmt.Println("Final digit:", final_digit)
        final_digit_int, _ := strconv.ParseInt(final_digit, 10, 0)
        digits = append(digits, final_digit_int)

        digit = nil
        index_of_digits = map[int]int64{}
    }
    fmt.Println("\n", digits)

    sum := 0
    for _, d := range digits {
        sum += int(d)
    }
    fmt.Println("\nSum:", sum)
}

func SubStringIndex(str, subStr string) int {
    for i := 0; i < len(str); i++ {
        if str[i:i+len(subStr)] == subStr {
            return i
        }
    }

    return -1
}
