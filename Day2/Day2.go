package main

import (
        "flag"
        "fmt"
//        "strconv"
//        "strings"
)

func checkCount(letterCount map[byte]int, count int) bool {
    for _,v := range letterCount {
        if v == count {
            return true
        }
    }

    return false
}

func calcFreq(inputList []string) (int, int) {

    threeCount := 0;
    twoCount := 0;

    for _, boxID := range inputList {
        letterCount := make(map[byte]int)
        fmt.Println("BoxID: ", boxID)

        for i := 0; i < len(boxID); i++ {
            boxChar := boxID[i]
            
            _, ok := letterCount[boxChar]
            
            if ok == true {
                letterCount[boxChar] += 1
            } else {
                letterCount[boxChar] = 1
            }
        }
        
        if checkCount(letterCount, 3) {
            threeCount += 1
        }
        
        if checkCount(letterCount, 2) {
            twoCount += 1
        }

        fmt.Println(letterCount)
    }

    return twoCount, threeCount
}

func main() {
    flag.Parse()
    inputList := flag.Args()

    twoCount, threeCount := calcFreq(inputList)
    fmt.Println("twoCount: ", twoCount, "threeCount", threeCount)
    fmt.Println("Checksum: ", twoCount * threeCount)
}
