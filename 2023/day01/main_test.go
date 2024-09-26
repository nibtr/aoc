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
