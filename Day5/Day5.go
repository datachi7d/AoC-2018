package main

import(
    "fmt"
    "io/ioutil"
    "bytes"
    "unicode"
)

func react(input []byte) (length int) {

    var i uint16

    i = 1

    for i < uint16(len(input)) {
        a := input[i]
        b := input[i-1]
        if ((a - b) == 32) || ((b - a) == 32) {
            input =  append(input[:i-1], input[i+1:]...)
            if i > 1 { i-=1 }
        } else {
            i++
        }
    }

    length = len(input)-1
    return
}

func reactC(input []byte, c rune, length chan int) {
    tmp := bytes.Map(func(r rune) rune { if r == c || r == unicode.ToUpper(c) { return -1 } else { return r } }, input)
    length <- react(tmp)
}


func main() {

    input, _ := ioutil.ReadFile("input.txt")

    input1 := make([]byte, len(input))
    copy(input1,input)

    fmt.Printf("Length part1: %d\n", react(input1))

    reactLengths := make(chan int)

    for i := 'a' ; i < 'z'; i++ {
       go reactC(input, i, reactLengths)
    }

    minLength := len(input)

    for i := 'a' ; i < 'z'; i++ {
        aLength := <- reactLengths
        if aLength < minLength {
            minLength = aLength
        }
    }

    fmt.Printf("Length part2: %d\n", minLength)
}
