package main

import (
	"fmt"
	"time"

	"github.com/pawelbijok/advent2024/days/day_1"
	"github.com/pawelbijok/advent2024/days/day_2"
)

func main() {
	day := 2
	part := 2
	start := time.Now().UnixNano()
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
	end := time.Now().UnixNano()
	elapsed := end - start
	describeTime(elapsed)
}

func describeTime(elapsed int64) {
	if elapsed < 1000000 {
		fmt.Printf("Time: %d ns\n", elapsed)
	} else {

		fmt.Printf("Time: %.3f ms\n", float64(elapsed)/1000000)
	}
}
