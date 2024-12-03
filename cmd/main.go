package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pawelbijok/advent2024/days/day_1"
	"github.com/pawelbijok/advent2024/days/day_2"
	"github.com/pawelbijok/advent2024/days/day_3"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Please provide day and part")
	}

	day, e := strconv.Atoi(args[1])

	if e != nil {
		fmt.Println("Invalid day arguments")
		os.Exit(1)
	}
	part, e := strconv.Atoi(args[2])
	if e != nil {
		fmt.Println("Invalid part arguments")
		os.Exit(1)
	}

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
	if day == 3 {
		if part == 1 {
			day_3.Part1("./days/day_3/input.txt")
		} else {
			day_3.Part2("./days/day_3/input.txt")
		}
	}
	end := time.Now().UnixNano()
	elapsed := end - start
	describeTime(elapsed)
}

func describeTime(elapsed int64) {
	if elapsed < 1000000 {
		fmt.Printf("Time: %.3f µs\n", float64(elapsed)/1000)
	} else {

		fmt.Printf("Time: %.3f ms\n", float64(elapsed)/1000000)
	}
}
