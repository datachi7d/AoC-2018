package main

import (
    "flag"
    "fmt"
)

type XY struct {
    X int
    Y int
}

func reserveSheet(input []string, sheet map[XY][]int, taken map[int]bool) (xSize int, ySize int, overlapCount int, noOverlapID int) {

	overlapCount = 0
	xSize = 0
	ySize = 0
	noOverlapID = 0

    for _, sheetSpec := range input {
        var ID int
        var pos XY
        var area XY

        fmt.Sscanf(sheetSpec, "#%d @ %d,%d: %dx%d", &ID, &pos.X, &pos.Y, &area.X, &area.Y)

		xMax := (pos.X + area.X)
		yMax := (pos.Y + area.Y)

		if xSize < xMax {
			xSize = xMax
		}
		if ySize < yMax {
			ySize = yMax
		}
        taken[ID] = false

        for xpos := pos.X; xpos < xMax; xpos++ {
            for ypos := pos.Y; ypos < yMax; ypos++ {

                _, ok := sheet[XY{xpos, ypos}]
                if ok == true {
					if len(sheet[XY{xpos,ypos}]) == 1 {
						overlapCount += 1
					}
                    sheet[XY{xpos,ypos}] = append(sheet[XY{xpos,ypos}], ID)
                    for _,takenID := range sheet[XY{xpos,ypos}] {
                        taken[takenID] = true
                    }
                } else {
                    sheet[XY{xpos,ypos}] = append(sheet[XY{xpos,ypos}], ID)
                }
            }
        }
    }

    for takenID, isTaken := range taken {
        if isTaken == false {
            noOverlapID = takenID
        }
    }

	xSize += 1
	ySize += 1

	return
}

func main() {
    flag.Parse()
    sheet := make(map[XY][]int)
    taken := make(map[int]bool)

    xSize, ySize, overlapCount, noOverlapID := reserveSheet(flag.Args(), sheet, taken)
	fmt.Println("Overlap count: ", overlapCount, " ", xSize, " ", ySize)
	fmt.Println("No overlapID: ", noOverlapID)

}
