package main

import "testing"

func TestSampleDataPart1(t *testing.T) {
	expected := 161
	result := solvePart1(("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"))
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart1(t *testing.T) {
	expected := 182619815
	result := part1(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestSampleDataPart2(t *testing.T) {
	expected := 48
	result := solvePart2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart2(t *testing.T) {
	expected := 80747545
	result := part2(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}
