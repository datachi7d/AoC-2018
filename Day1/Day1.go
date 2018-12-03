package main

import (
        "flag"
        "fmt"
        "strconv"
        "strings"
)

func calcFreq(changesList []string) int {

    freq := 0

    for _, changeStr := range changesList {
        convertStr := strings.TrimSpace(changeStr)
        changeInt, err := strconv.Atoi(convertStr)
        if err != nil {
            panic(err)
        }
        freq += changeInt
    }

    return freq
}

func main() {
    var input string;
    flag.StringVar(&input, "input", "", "Input frequency changes")
    flag.Parse()
    inputList := strings.Split(input, ",")
    fmt.Println("Input: ", inputList)
    fmt.Println("Frequency: ", calcFreq(inputList))
}
