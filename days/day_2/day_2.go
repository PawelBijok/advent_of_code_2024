package day_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part1(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	safeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		lineNumbers := []int{}
		for _, word := range words {

			num, err := strconv.Atoi(word)
			if err != nil {
				continue
			}
			lineNumbers = append(lineNumbers, num)
		}
		lastNumber := lineNumbers[0]
		// -1 goes down, 1 goes up, 0 is not set
		direction := 0
		isSafe := true
		for i := 1; i < len(lineNumbers); i++ {
			if direction == 0 {
				if lineNumbers[i] > lastNumber {
					direction = 1
				} else {
					direction = -1
				}
			}
			if direction == 1 && lineNumbers[i] < lastNumber {
				isSafe = false
				break
			}
			if direction == -1 && lineNumbers[i] > lastNumber {
				isSafe = false
				break
			}
			difference := abs(lineNumbers[i] - lastNumber)
			if difference > 3 || difference < 1 {
				isSafe = false
				break
			}
			lastNumber = lineNumbers[i]
		}
		if isSafe {
			safeReports += 1
		}
	}

	fmt.Println(safeReports)
}

func validateLine(lineNumbers []int) bool {
	lastNumber := lineNumbers[0]
	// -1 goes down, 1 goes up, 0 is not set
	direction := 0
	isSafe := true
	for i := 1; i < len(lineNumbers); i++ {
		if direction == 0 {
			if lineNumbers[i] > lastNumber {
				direction = 1
			} else {
				direction = -1
			}
		}
		if direction == 1 && lineNumbers[i] < lastNumber {
			isSafe = false
			break
		}
		if direction == -1 && lineNumbers[i] > lastNumber {
			isSafe = false
			break
		}
		difference := abs(lineNumbers[i] - lastNumber)
		if difference > 3 || difference < 1 {
			isSafe = false
			break
		}
		lastNumber = lineNumbers[i]
	}
	return isSafe
}

func Part2(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	safeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		lineNumbers := []int{}
		for _, word := range words {

			num, err := strconv.Atoi(word)
			if err != nil {
				continue
			}
			lineNumbers = append(lineNumbers, num)
		}
		isSafe := validateLine(lineNumbers)
		if !isSafe {
			for i := 0; i < len(lineNumbers); i++ {
				numbers := append([]int{}, lineNumbers...)
				numbers = append(numbers[:i], numbers[i+1:]...)
				newResult := validateLine(numbers)
				if newResult {
					isSafe = true
					break
				}
			}
		}
		if isSafe {
			safeReports += 1
		}
	}

	fmt.Println(safeReports)

}
