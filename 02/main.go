package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strconv"
	"strings"
)

func readIDs(filepath string) iter.Seq[string] {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fileString := string(bytes)
	giftIDs := strings.SplitSeq(fileString, ",")

	return giftIDs
}

func getRange(rangeString string) []int {
	parts := strings.Split(rangeString, "-")

	lowerBound, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	upperBound, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

	return []int{lowerBound, upperBound}
}

func isRepeatedPattern(id int) bool {
	idString := strconv.Itoa(id)
	n := len(idString)

	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen != 0 {
			continue
		}

		pattern := idString[:patternLen]
		isMatch := true

		for i := patternLen; i < n; i += patternLen {
			if idString[i:i+patternLen] != pattern {
				isMatch = false
				break
			}
		}

		if isMatch {
			return true
		}
	}

	return false
}

func findInvalidIDsPart1(lowerBound int, upperBound int) int {
	sum := 0

	for id := lowerBound; id < upperBound+1; id++ {
		idString := strconv.Itoa(id)
		half := len(idString) / 2

		firstHalf := idString[:half]
		secondHalf := idString[half:]

		if firstHalf == secondHalf {
			sum += id
		}
	}

	return sum
}

func findInvalidIDsPart2(lowerBound int, upperBound int) int {
	sum := 0

	for id := lowerBound; id <= upperBound; id++ {
		if isRepeatedPattern(id) {
			sum += id
		}
	}

	return sum
}

func main() {

	giftIDs := readIDs("input.txt")
	sumInvalidPart1 := 0
	sumInvalidPart2 := 0

	for id := range giftIDs {
		lowerBound := getRange(id)[0]
		upperBound := getRange(id)[1]

		sumInvalidPart1 += findInvalidIDsPart1(lowerBound, upperBound)
		sumInvalidPart2 += findInvalidIDsPart2(lowerBound, upperBound)

	}

	fmt.Println(sumInvalidPart1)
	fmt.Println(sumInvalidPart2)

}
