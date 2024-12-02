package day_1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	l1 := []int{}
	l2 := []int{}

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
		l1 = append(l1, lineNumbers[0])
		l2 = append(l2, lineNumbers[1])
	}
	sort.Ints(l1)
	sort.Ints(l2)
	var totalDistance int
	for i := 0; i < len(l1); i++ {
		distance := l1[i] - l2[i]
		totalDistance += abs(distance)
	}
	fmt.Println(totalDistance)
}
func Part2(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	l1 := []int{}
	l2 := []int{}

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
		l1 = append(l1, lineNumbers[0])
		l2 = append(l2, lineNumbers[1])
	}
	var similarityScore int = 0
	for _, number := range l1 {
		timesAppeared := 0
		for _, number2 := range l2 {
			if number == number2 {
				timesAppeared += 1
			}
		}
		similarityScore += timesAppeared * number
	}
	fmt.Println(similarityScore)
}
