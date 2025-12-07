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

func getRows(rowsSeq iter.Seq[string]) [][]string {
	var rows [][]string
	for row := range rowsSeq {

		splitRow := strings.Split(row, " ")
		var cols []string
		for _, col := range splitRow {
			trimmed := strings.TrimSpace(col)
			if trimmed != "" {
				cols = append(cols, trimmed)
			}
		}
		rows = append(rows, cols)

	}

	return rows
}

func getRowsPartTwo(rowSeq iter.Seq[string]) [][]string {
	var rows [][]string
	for i := range rowSeq {
		var col []string
		col = append(col, i)

		rows = append(rows, col)
	}
	return rows
}

func findColLength(row string) []int {
	var colLength []int
	count := 0
	for i := 0; i < len(row); i++ {
		currVal := row[i]

		if currVal == ' ' {
			count++
		} else {
			if count == 0 {
				continue
			}
			colLength = append(colLength, count)
			count = 0
		}

		if i == len(row)-1 {
			count++
			colLength = append(colLength, count)
		}
	}

	return colLength
}

func parseVal(val string) int {
	valInt, _ := strconv.Atoi(val)
	return valInt
}

type Col struct {
	slice []int
	op    string
}

func (c Col) calculateResult() int {
	var total int

	switch c.op {
	case "+":
		total = 0
		for i := range c.slice {
			total += c.slice[i]
		}
	case "*":
		total = 1
		for i := range c.slice {
			total *= c.slice[i]
		}
	}

	return total
}

func main() {
	rowsSeq := readRows("example.txt")

	rows := getRows(rowsSeq)

	nRows := len(rows)
	nCols := len(rows[0])
	total := 0

	for i := range nCols {
		col := Col{}
		for j := range nRows {
			val := rows[j][i]
			if val == "+" {
				col.op = "+"
				continue
			} else if val == "*" {
				col.op = "*"
				continue
			}

			valParsed := parseVal(val)
			col.slice = append(col.slice, valParsed)

		}
		total += col.calculateResult()
	}
	// p1
	fmt.Println(total)

	rowSeqP2 := readRows("input.txt")
	rows = getRowsPartTwo(rowSeqP2)

	colLength := findColLength(rows[len(rows)-1][0])
	currPos := 0
	totalTotal := 0

	for i := 0; i < len(colLength); i++ {
		endPos := currPos + colLength[i]

		var op string
		var numbers []int

		for charIdx := 0; charIdx < colLength[i]; charIdx++ {
			var numStr string

			for rowIdx := 0; rowIdx < len(rows)-1; rowIdx++ {
				char := rows[rowIdx][0][currPos+charIdx]

				if char != ' ' {
					numStr += string(char)
				}
			}

			opChar := rows[len(rows)-1][0][currPos+charIdx]
			if opChar == '*' || opChar == '+' {
				op = string(opChar)
			}

			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
		}

		c := Col{
			slice: numbers,
			op:    op,
		}

		totalTotal += c.calculateResult()
		currPos = endPos + 1
	}

	//p2
	fmt.Println(totalTotal)

}

// [0:3] [3:6] [7:11] [12:14]
