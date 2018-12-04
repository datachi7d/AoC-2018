package main

import (
    "flag"
    "fmt"
	"time"
)

type XY struct {
    X int
    Y int
}




func processEvents(input []string) {


    var GuardID int
    for _, timeEvent := range input {
        var Year int
        var Month int
        var Day int
        var Hour int
        var Minute int
        var EventString string

        fmt.Sscanf(timeEvent, "[%d-%d-%d %d:%d]", &Year, &Month, &Day, &Hour, &Minute)

        EventString = timeEvent[19:]
        EventTime := (Hour*60)+Minute

        dayTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.UTC)

        if EventTime > 1000 {
            EventTime = 0
            dayTime = dayTime.AddDate(0,0,1)
            Day = dayTime.Day()
            Month = int(dayTime.Month())
        }

        fmt.Printf("%d %d ", dayTime.Unix(), EventTime)


        count,_ := fmt.Sscanf(EventString, "Guard #%d", &GuardID)

        fmt.Printf("%d ", GuardID )
        if count == 1 {
            fmt.Printf("\n")
        } else {
            if EventString == "wakes up" {
               fmt.Printf("true\n")
            } else if EventString == "falls asleep" {
               fmt.Printf("false\n")
            } else {
                panic("Input error")
            }
        }
    }
}

func main() {
    flag.Parse()

    processEvents(flag.Args())
}
