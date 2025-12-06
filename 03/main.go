package main

import (
	"fmt"
	"io"
	"iter"
	"math"
	"os"
	"strconv"
	"strings"
)

func readBanks(filepath string) iter.Seq[string] {
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
	banks := strings.SplitSeq(fileString, "\n")

	return banks
}

func findMax(slice []int) int {
	max := 0
	for _, n := range slice {
		if n > max {
			max = n
		}
	}
	return max
}

func checkValuePresence(val int, slice []int) bool {
	for _, n := range slice {
		if n == val {
			return true
		}
	}
	return false
}

func findValueIndex(val int, slice []int) int {
	for i, n := range slice {
		if n == val {
			return i
		}
	}
	// terrible but will only call if i checked presence
	return 9999
}

type Battery struct {
	combinedBattery int
	bankRemaining   []int
	addCandidate    bool
}

func findCombinedBatteryJoltage(maxBattery int, bank []int) *Battery {
	if !checkValuePresence(maxBattery, bank) {
		return &Battery{addCandidate: false}
	}

	maxIndex := findValueIndex(maxBattery, bank)
	if maxIndex == len(bank)-1 {
		// case where the biggest single battery value is right at the end
		// would need to repeat process with the entire slice minus the last digit
		// but then continue with the rest of the slide on the right side of the new max
		newMax := findMax(bank[:len(bank)-1])
		maxIndex = findValueIndex(newMax, bank[:len(bank)-1])

		secondMax := findMax(bank[maxIndex+1:])
		secondMaxIndex := findValueIndex(secondMax, bank[maxIndex+1:])

		combinedBatteryValue := (bank[maxIndex] * 10) + bank[maxIndex+1:][secondMaxIndex]

		return &Battery{
			combinedBattery: combinedBatteryValue,
			addCandidate:    true,
		}
	}
	secondMax := findMax(bank[maxIndex+1:])
	secondMaxIndex := findValueIndex(secondMax, bank[maxIndex+1:])

	combinedBatteryValue := (bank[maxIndex] * 10) + bank[maxIndex+1:][secondMaxIndex]

	return &Battery{
		combinedBattery: combinedBatteryValue,
		bankRemaining:   bank[maxIndex+2:],
		addCandidate:    true,
	}

}

func part1(banks iter.Seq[string]) int {
	totalJoltage := 0
	for bankString := range banks {
		var candidates []int
		bankSeq := strings.SplitSeq(bankString, "")
		var bank []int

		for battery := range bankSeq {
			intBattery, _ := strconv.Atoi(battery)
			bank = append(bank, intBattery)
		}

		updatedBank := bank
		bankMax := findMax(updatedBank)

		for checkValuePresence(bankMax, updatedBank) {
			newBattery := findCombinedBatteryJoltage(bankMax, updatedBank)

			if newBattery.addCandidate {
				updatedBank = newBattery.bankRemaining
			} else {
				updatedBank = newBattery.bankRemaining
				bankMax = findMax(updatedBank)
			}

			candidates = append(candidates, newBattery.combinedBattery)
		}

		totalJoltage += findMax(candidates)
	}
	return totalJoltage
}

func part2(banks iter.Seq[string]) int {
	totalJotalge := 0

	for bankString := range banks {
		bankSeq := strings.SplitSeq(bankString, "")
		var bank []int

		for battery := range bankSeq {
			intBattery, _ := strconv.Atoi(battery)
			bank = append(bank, intBattery)
		}

		updatedBank := bank

		digitsLeft := 12
		result := 0
		for digitsLeft > 0 {
			bankMax := findMax(updatedBank[:len(updatedBank)-digitsLeft+1])
			indexMax := findValueIndex(bankMax, updatedBank[:len(updatedBank)-digitsLeft+1])

			result += int(math.Pow(float64(10), float64(digitsLeft-1))) * bankMax
			updatedBank = updatedBank[indexMax+1:]
			digitsLeft--

		}
		totalJotalge += result
	}

	return totalJotalge
}

func main() {
	banks := readBanks("input.txt")

	fmt.Println(part1(banks))
	banks2 := readBanks("input.txt")
	fmt.Println(part2(banks2))
}
