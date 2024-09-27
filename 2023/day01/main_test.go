package main

import "testing"

type TestCase struct {
	name  string
	input string
	want  int
}

func Test_part1(t *testing.T) {
	tests := []TestCase{
		{
			name: "example1",
			input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := part1(tc.input); got != tc.want {
				t.Errorf("part1() = %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []TestCase{
		{
			name: "example1",
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := part2(tc.input); got != tc.want {
				t.Errorf("part1() = %v, want %v", got, tc.want)
			}
		})
	}
}
