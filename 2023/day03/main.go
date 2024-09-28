package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	utils "github.com/nibtr/aoc/utils"
)

var symbols = "`~!@#$%^&*()-_=+[{]}|\\;:'\"<,>?/"

// var dir = [8][2]int{
// 	{-1, -1},
// 	{-1, 0},
// 	{-1, 1},
// 	{0, -1},
// 	{0, 1},
// 	{1, -1},
// 	{1, 0},
// 	{1, 1},
// }

func main() {
	var part int

	flag.IntVar(&part, "part", 1, "Run part 1 or 2")
	flag.Parse()
	input := utils.ReadFile("/2023/input/day03.txt")

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for row, line := range lines {
		step := 1
		for i := 0; i < len(line); i += step {
			num := ""
			for j := i; j < len(line) && isDigit(line[j]); j++ {
				num += string(line[j])
			}

			if num != "" {
				if hasAdjacentSymbols(lines, row, i, i+len(num)-1) {
					c, _ := strconv.Atoi(num)
					sum += c
					fmt.Println(c)
				}
				step = len(num)
			} else {
				step = 1
			}
		}
	}

	return sum
}

func part2(input string) int {
	fmt.Println(input)
	return 0
}

func hasAdjacentSymbols(matrix []string, row int, startIdx int, endIdx int) bool {
	valid := false

	leftIdx := int(math.Max(0, float64(startIdx-1)))
	rightIdx := int(math.Min(float64(len(matrix[row])-1), float64(endIdx+1)))
	left := matrix[row][leftIdx]
	right := matrix[row][rightIdx]

	if strings.ContainsAny(string(left), symbols) || strings.ContainsAny(string(right), symbols) {
		valid = true
	}

	rowAbove := int(math.Max(0, float64(row-1)))
	rowBelow := int(math.Min(float64(len(matrix)-1), float64(row+1)))

	if strings.ContainsAny(matrix[rowAbove][leftIdx:rightIdx+1], symbols) ||
		strings.ContainsAny(matrix[rowBelow][leftIdx:rightIdx+1], symbols) {
		valid = true
	}

	return valid
}

func isDigit(c byte) bool {
	return strings.ContainsAny(string(c), "0123456789")
}
