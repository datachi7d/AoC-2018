package main

import (
    "flag"
    "fmt"
	"time"
)

type guardEvent struct {
    minute int
    asleep bool
}

type guardEventRange struct {
    start int
    stop int
}

func processEvents(input []string, gaurdData map[int]map[time.Time][]guardEvent) {


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

        count,_ := fmt.Sscanf(EventString, "Guard #%d", &GuardID)

        _, ok := gaurdData[GuardID]

        if ok == false {
            gaurdData[GuardID] = make(map[time.Time][]guardEvent)
        }

        _, ok = gaurdData[GuardID][dayTime]

        if count != 1 {
            guardE := guardEvent{Minute, false}
            if EventString == "wakes up" {
               guardE.asleep = false
            } else if EventString == "falls asleep" {
               guardE.asleep = true
            } else {
                panic("Input error")
            }

            gaurdData[GuardID][dayTime] = append(gaurdData[GuardID][dayTime], guardE)
        }
    }
    
    gaurdMinutes := make(map[int][]int)
    maxMinutes := 0
    maxMinutesValue := 0
    maxGuard := 0
    maxTotalGuard := 0
    maxTotalGuardID := 0

    for GuardID, allEvents := range gaurdData {
        gaurdMinutes[GuardID] = make([]int, 120)
        totalGuard := 0
        for day,dayEvents := range allEvents {
           fmt.Println(GuardID, " ", day, " ", dayEvents)
           for i := 0; i < len(dayEvents); i+=2 {
                start := dayEvents[i]
                end := dayEvents[i+1]
                if (start.asleep == true) && (end.asleep == false) {
                    for ii := start.minute; ii < end.minute; ii++ {
                        gaurdMinutes[GuardID][ii] += 1
                        totalGuard += 1
                        if gaurdMinutes[GuardID][ii] > maxMinutes {
                            maxMinutes =  gaurdMinutes[GuardID][ii]
                            maxMinutesValue = ii
                            maxGuard = GuardID
                        }
                    }
                } else {
                    panic("Input error")
                }
           }
        }
        if totalGuard > maxTotalGuard {
            maxTotalGuard = totalGuard
            maxTotalGuardID = GuardID
        }
    }


    maxTotalMinutesValue  := 0
    maxTotalMinutes := 0

    for i, minutes := range gaurdMinutes[maxTotalGuardID] {
        if minutes > maxTotalMinutes {
            maxTotalMinutesValue = i
            maxTotalMinutes = minutes
        }
    }

    fmt.Println("maxGuard ",maxTotalGuardID, " maxTotalMinutesValue", maxTotalMinutesValue, " answer ", maxTotalGuardID*maxTotalMinutesValue)
    fmt.Println("maxGuard ",maxGuard, " maxMinutesValue", maxMinutesValue, " answer ", maxGuard*maxMinutesValue)

}

func main() {
    flag.Parse()

    var guardData = map[int]map[time.Time][]guardEvent{}

    processEvents(flag.Args(), guardData)
}
