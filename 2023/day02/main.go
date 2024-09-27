package main

import (
	"flag"
	"fmt"

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
	fmt.Println(input)
	return 0
}

func part2(input string) int {
	input = ""
	return 0
}
