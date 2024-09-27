package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	utils "github.com/nibtr/aoc/utils"
)

func main() {
	var part int

	flag.IntVar(&part, "part", 1, "Run part 1 or 2")
	flag.Parse()

	content := utils.ReadFile("/2023/input/day01.txt")
	if part == 1 {
		ans := part1(content)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(content)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		var tens, ones int
		for i := 0; i < len(line); i++ {
			if strings.ContainsAny(line[i:i+1], "0123456789") {
				tens, _ = strconv.Atoi(string(line[i]))
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if strings.ContainsAny(line[i:i+1], "0123456789") {
				ones, _ = strconv.Atoi(string(line[i]))
				break
			}
		}

		sum += tens*10 + ones
	}
	return sum
}

func part2(input string) int {
	// sum := 0
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range strings.Split(input, "\n") {
		var tens int
		step := 1

		for i := 0; i < len(line); i += step {
			// 3-letter digit
			if v, ok := digits[line[i:i+3]]; ok {
				tens = v
				step = 3
			} else if v, ok := digits[line[i:i+4]]; ok {
				tens = v
				step = 4
			} else if v, ok := digits[line[i:i+5]]; ok {
				tens = v
				step = 5
			} else {
				step = 1
			}
		}

		fmt.Println(tens)
	}

	return 0
}

// aoeightwtwo

// one
// two
// six
// four
// five
// nine
// three
// seven
// eight
