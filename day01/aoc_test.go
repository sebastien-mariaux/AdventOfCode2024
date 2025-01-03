package main

import "testing"

func TestSampleDataPart1(t *testing.T) {
	expected := 11
	result := part1(true)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart1(t *testing.T) {
	expected := 2580760
	result := part1(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestSampleDataPart2(t *testing.T) {
	expected := 31
	result := part2(true)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart2(t *testing.T) {
	expected := 25358365
	result := part2(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}
