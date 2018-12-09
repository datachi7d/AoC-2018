package main

import (
    "flag"
    "fmt"
    "sort"
    "strings"
)

func deleteTop(order map[string][]string) []string  {

    if(len(order) == 0) {
        return nil
    }

    counter := make(map[string]int)

    for k,vs := range order {
        if len(order) == 1  {
            return append([]string{k}, order[k]...)
        }

        _, ok := counter[k]
        if ok == false {
            counter[k] = 0
        }
        for _,v := range vs {
            _, ok := counter[v]
            if ok == false {
                counter[v] = 1
            } else {
                counter[v] += 1
            }
       }
    }

    var top []string
    for k,v := range counter {
        if v == 0 {
            top = append(top, k)
        }
    }

    sort.Strings(top)
    delete(order, top[0])

    fmt.Println(top)

    return  append([]string{top[0]}, deleteTop(order)...)
}


func main() {
    flag.Parse()
    inputs := flag.Args()

    order := make(map[string][]string)

    for _,input := range inputs {
        step1 := string(input[5])
        step2 := string(input[36])
        order[step1] = append(order[step1], step2)
    }

    fmt.Println(strings.Join(deleteTop(order),""))
}
