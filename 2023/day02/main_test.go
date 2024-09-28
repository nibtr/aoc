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
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			want: 8,
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
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			want: 2286,
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
