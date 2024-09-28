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
	input := utils.ReadFile("/2023/input/day02.txt")

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	gameConfig := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	lines := strings.Split(input, "\n")
	sum := 0
	i := 1
	for _, game := range lines {
		sets := strings.Split(game, ": ")[1]
		sets = strings.ReplaceAll(sets, ";", ",")
		picks := strings.Split(sets, ", ")

		for _, pick := range picks {
			num := strings.Split(pick, " ")[0]
			color := strings.Split(pick, " ")[1]
			if c, _ := strconv.Atoi(num); c > gameConfig[color] {
				sum -= i
				break
			}
		}

		sum += i
		i++
	}

	return sum
}

func part2(input string) int {
	input = ""
	return 0
}
