package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getRow(rowCords string) int {
	rowsRange := []int{0, 127}

	for _, cord := range rowCords {
		if cord == 'F' {
			rowsRange[1] = (rowsRange[0] + rowsRange[1] - 1) / 2
		} else {
			rowsRange[0] = (rowsRange[0] + rowsRange[1] + 1) / 2
		}
	}

	return rowsRange[0]
}

func getCol(colCords string) int {
	colsRange := []int{0, 7}

	for _, cord := range colCords {
		if cord == 'L' {
			colsRange[1] = (colsRange[0] + colsRange[1] - 1) / 2
		} else {
			colsRange[0] = (colsRange[0] + colsRange[1] + 1) / 2
		}
	}

	return colsRange[0]
}

func getSeatID(boardingPass string) int {
	rowCords := boardingPass[:7]
	colCords := boardingPass[7:]

	row := getRow(rowCords)
	col := getCol(colCords)

	return row*8 + col
}

func getSeatIDs(boardingPasses []string) []int {
	var seatIDs []int

	for _, passes := range boardingPasses {
		seatID := getSeatID(passes)
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)
	return seatIDs
}

func part1(boardingPasses []string) {
	seatIDs := getSeatIDs(boardingPasses)
	fmt.Printf("Part 1: %d\n", seatIDs[len(seatIDs)-1])
}

func part2(boardingPasses []string) {
	seatIDs := getSeatIDs(boardingPasses)

	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1]-seatIDs[i] == 2 {
			fmt.Printf("Part 2: %d\n", seatIDs[i]+1)
			return
		}
	}
	fmt.Println("Part 2: No gap found")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	boardingPasses := []string{}
	for scanner.Scan() {
		boardingPasses = append(boardingPasses, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}

	part1(boardingPasses)
	part2(boardingPasses)
}
