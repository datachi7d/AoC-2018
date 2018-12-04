package main

import (
        "flag"
        "fmt"
)

func Hamming(s1 string, s2 string) (distance int, diffIndex int) {
	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	for i, v := range r1 {
		if r2[i] != v {
			distance += 1
			diffIndex = i
		}
	}
	return
}


func checkCount(letterCount map[byte]int, count int) bool {
    for _,v := range letterCount {
        if v == count {
            return true
        }
    }

    return false
}

func findID(inputList []string) {

    for boxIndex, boxID := range inputList {
        fmt.Println("BoxIndex: ", boxIndex, " BoxID: ", boxID)

        for i := boxIndex+1; i < len(inputList); i++ {
            distance, diffIndex := Hamming(boxID, inputList[i])

			if(distance == 1) {
				boxID2 := inputList[i]
				fmt.Println("BoxID_1: ", boxID, " BoxID_2: ", boxID2)
                fmt.Println("diffIndex: ", diffIndex)
				result1 := boxID[0:diffIndex] + boxID[diffIndex+1:]
				result2 := boxID2[0:diffIndex] + boxID2[diffIndex+1:]
                if (result1 == result2) && (len(result1)+1 == len(boxID)) {
                    fmt.Println("Result: ", result1)
                }
                return
			}
        }

    }
}

func main() {
    flag.Parse()
    inputList := flag.Args()

    findID(inputList)
}
