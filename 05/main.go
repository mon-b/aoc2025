package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strconv"
	"strings"
)

func readRows(filepath string) iter.Seq[string] {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fileString := string(bytes)
	rows := strings.SplitSeq(fileString, "\n")

	return rows
}

func checkFresh(targetRange *Range, productID int) bool {
	if productID <= targetRange.upperBound && productID >= targetRange.lowerBound {
		return true
	}
	return false
}

type Range struct {
	lowerBound int
	upperBound int
}

func main() {
	rowsSeq := readRows("input.txt")
	var ranges []*Range

	processingRanges := true
	fresh := 0

	for row := range rowsSeq {
		if row == "" {
			processingRanges = false
			continue
		}

		if processingRanges {
			splitRow := strings.Split(row, "-")

			lowerBound, _ := strconv.Atoi(splitRow[0])
			upperBound, _ := strconv.Atoi(splitRow[1])

			intRange := Range{
				lowerBound: lowerBound,
				upperBound: upperBound,
			}

			ranges = append(ranges, &intRange)

		}

		if !processingRanges {
			productID, _ := strconv.Atoi(row)
			markedFresh := false
			for _, r := range ranges {
				isFresh := checkFresh(r, productID)
				if isFresh && !markedFresh {
					markedFresh = true
					fresh++

				}
			}

		}
	}

	fmt.Println(fresh)

}
