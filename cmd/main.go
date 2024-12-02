package main

import (
	"github.com/pawelbijok/advent2024/days/day_1"
	"github.com/pawelbijok/advent2024/days/day_2"
)

func main() {
	day := 2
	part := 2
	if day == 1 {
		if part == 1 {
			day_1.Part1("./days/day_1/input.txt")
		} else {
			day_1.Part2("./days/day_1/input.txt")
		}
	}
	if day == 2 {
		if part == 1 {
			day_2.Part1("./days/day_2/input.txt")
		} else {
			day_2.Part2("./days/day_2/input.txt")
		}
	}
}
