package day_3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func findMuls(s string) []string {
	mulPattern := `mul\(\d{1,3},\d{1,3}\)`
	mulRegexp := regexp.MustCompile(mulPattern)

	m := mulRegexp.FindAllString(s, -1)

	return m
}
func getDigitsFromMul(s string) []int {

	digitPattern := `\d{1,3}`
	digitRegexp := regexp.MustCompile(digitPattern)

	digits := digitRegexp.FindAllString(s, -1)
	result := []int{}
	for _, digit := range digits {
		digitInt, _ := strconv.Atoi(digit)
		result = append(result, digitInt)
	}
	return result
}

func Part1(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		muls := findMuls(line)
		for _, mul := range muls {
			digits := getDigitsFromMul(mul)
			partialMult := 1
			for _, d := range digits {
				partialMult *= d
			}
			sum += partialMult
		}
	}
	fmt.Println(sum)
}

func findInstructions(s string) []string {
	instructionPattern := `(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`
	instructionRegexp := regexp.MustCompile(instructionPattern)

	instructions := instructionRegexp.FindAllString(s, -1)

	return instructions
}

func Part2(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allInstructions := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		instructions := findInstructions(line)
		allInstructions = append(allInstructions, instructions...)

	}

	sum := 0
	enabled := true
	for _, instruction := range allInstructions {
		if instruction == "don't()" {
			enabled = false
			continue
		}
		if instruction == "do()" {
			enabled = true
			continue
		}

		if enabled {
			numbers := getDigitsFromMul(instruction)
			partialSum := 1
			for _, n := range numbers {
				partialSum *= n
			}
			sum += partialSum
		}
	}
	fmt.Println(sum)
}
