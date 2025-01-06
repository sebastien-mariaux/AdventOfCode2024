package main

import "testing"

func TestSampleDataPart1(t *testing.T) {
	expected := 18
	result := part1(true)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart1(t *testing.T) {
	expected := 2524
	result := part1(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestSampleDataPart2(t *testing.T) {
	expected := 9
	result := part2(true)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestDataPart2(t *testing.T) {
	expected := 1873
	result := part2(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}
