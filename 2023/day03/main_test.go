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
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 4361,
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
