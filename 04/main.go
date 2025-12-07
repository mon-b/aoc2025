package main

import (
	"fmt"
	"io"
	"iter"
	"os"
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

func checkRow(index int, row []string, nCols int) int {
	count := 0

	if index-1 >= 0 && row[index-1] == "@" {
		count++
	}

	if row[index] == "@" {
		count++
	}

	if index+1 < nCols && row[index+1] == "@" {
		count++
	}

	return count
}

func checkUpperRowBound(target int) bool {
	return target >= 0
}

func checkLowerRowBound(length int, target int) bool {
	return target < length
}

func main() {
	rowsSeq := readRows("input.txt")

	var rows [][]string

	for row := range rowsSeq {
		var rowSlice []string
		for _, col := range row {
			rowSlice = append(rowSlice, string(col))
		}
		rows = append(rows, rowSlice)
	}
	rowsUpdated := make([][]string, len(rows))
	for i := range rows {
		rowsUpdated[i] = make([]string, len(rows[i]))
	}

	nRows := len(rows)

	access := 9999
	finalCount := 0

	for access > 0 {
		for i := range rows {
			copy(rowsUpdated[i], rows[i])
		}
		access = 0
		for indexRow, row := range rows {
			nCols := len(row)
			for indexCol, _ := range row {
				if row[indexCol] != "@" {
					continue
				}

				count := 0

				upperRowIndex := indexRow - 1
				if checkUpperRowBound(upperRowIndex) {
					count += checkRow(indexCol, rows[upperRowIndex], nCols)
				}
				lowerRowIndex := indexRow + 1
				if checkLowerRowBound(nRows, lowerRowIndex) {
					count += checkRow(indexCol, rows[lowerRowIndex], nCols)
				}

				if indexCol-1 >= 0 {
					if row[indexCol-1] == "@" {
						count++
					}
				}

				if indexCol+1 < nCols {
					if row[indexCol+1] == "@" {
						count++
					}
				}

				if count < 4 {
					// can be accessed
					rowsUpdated[indexRow][indexCol] = "."
					access++
				}
			}
		}
		fmt.Println(access) // first print of this is p1
		copy(rows, rowsUpdated)
		finalCount += access

	}
	fmt.Println(finalCount) // p2

}
