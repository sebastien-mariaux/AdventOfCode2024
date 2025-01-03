package main

import "testing"

func TestSampleDataPart1(t *testing.T) {
	expected := 2
	result := part1(true)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart1(t *testing.T) {
	expected := 230
	result := part1(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestSampleDataPart2(t *testing.T) {
	expected := 4
	result := part2(true)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart2(t *testing.T) {
	expected := 0
	result := part2(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}
