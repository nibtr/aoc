package main

import "testing"

func Test_part1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	want := 142
	got := part1(input)

	if got != want {
		t.Fatalf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	want := 281
	got := part2(input)

	if got != want {
		t.Fatalf("part2() = %v, want %v", got, want)
	}
}
