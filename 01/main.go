package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func rightTurn(number int, dial int) int {
	newDial := (dial + number) % 100
	return newDial
}

func leftTurn(number int, dial int) int {
	newDial := (dial - number%100 + 100) % 100
	return newDial
}

func rotate(rotation string, dial int) [2]int {
	direction := rotation[0]
	n, _ := strconv.Atoi(rotation[1:])
	counter := n / 100

	var newDial int

	if direction == 'L' {
		newDial = leftTurn(n, dial)
		remainder := n % 100
		if newDial != 0 && dial-remainder < 0 && dial != 0 {
			counter++
		}

	} else {
		newDial = rightTurn(n, dial)
		remainder := n % 100
		if newDial != 0 && dial+remainder >= 100 && dial != 0 {
			counter++
		}
	}

	return [2]int{newDial, counter}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	dial := 50
	var dialZeroCounter int
	var rotationZeroCounter int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		rotation := scanner.Text()
		result := rotate(rotation, dial)

		dial = result[0]

		if dial == 0 {
			dialZeroCounter++
		}
		rotationZeroCounter += result[1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	// solution to 1
	fmt.Println(dialZeroCounter)
	// solution to 2
	fmt.Println(dialZeroCounter + rotationZeroCounter)
}
