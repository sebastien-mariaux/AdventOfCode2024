package day01

import "testing"

func TestSampleData(t *testing.T) {
  expected := 11
	result := solve(true)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}

func TestData(t *testing.T) {
  expected := 2580760
	result := solve(false)
	if result != expected {
		t.Errorf("Failed - Expected %d, Got %d", expected, result)
	}
}
