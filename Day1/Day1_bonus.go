package main

import (
        "flag"
        "fmt"
        "strconv"
        "strings"
)

func calcFreq(freqStart int, changesList []string, freqCount (map[int]int)) (int, bool) {

    freq := freqStart

    for _, changeStr := range changesList {
        convertStr := strings.TrimSpace(changeStr)
        changeInt, err := strconv.Atoi(convertStr)
        if err != nil {
            panic(err)
        }
        freq += changeInt

        _, ok := freqCount[freq]

        if ok == true {
            freqCount[freq] += 1
            return freq, true
        } else {
            freqCount[freq] = 1
        }
    }

    return freq, false
}

func main() {
    var input string;
    flag.StringVar(&input, "input", "", "Input frequency changes")
    flag.Parse()
    inputList := strings.Split(input, ",")
    fmt.Println("Input: ", inputList)

    freqCount := make(map[int]int)
    freq := 0
    found := false
    calcIterations := 0

    for (found == false) && (calcIterations < 1000) {
        freq, found = calcFreq(freq, inputList, freqCount)
        calcIterations += 1
    }


    fmt.Println("Frequency: ", freq, "Found: ", found, " Iterations: ", calcIterations)
}
