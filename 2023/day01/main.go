package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var part int
	var input string

	flag.IntVar(&part, "part", 1, "Run part 1 or 2")
	flag.StringVar(&input, "input", "", "Choose input")
	flag.Parse()

	fmt.Println("Run part", part)
	fmt.Println("Run input", input)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part1(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	data := strings.TrimRight(input, "\n")
	sum := 0
	for _, line := range strings.Split(data, "\n") {
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
