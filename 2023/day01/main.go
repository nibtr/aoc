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
	sum := 0
	for _, line := range strings.Split(input, "\n") {

	}

	return 0
}
