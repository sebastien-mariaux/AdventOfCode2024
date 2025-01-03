package main

import (
	"aoc_2024/helpers"
	"strconv"
	"strings"
)

func isSafe(values []int) bool {

	isDescending := values[len(values)-1] < values[0]
	isAscending := values[len(values)-1] > values[0]
	if !(isDescending || isAscending) {
		return false
	}

	// Loop over following values
	for i := 1; i < len(values); i++ {
		diff := helpers.AbsoluteDiff(values[i], values[i-1])
		if (isDescending && values[i] > values[i-1]) || (isAscending && values[i] < values[i-1]) || (diff < 1 || diff > 3) {
			return false
		}
	}
	return true
}

func rowToValues(row string) []int {
	var values []int
	for _, value := range strings.Split(row, " ") {
		int_value, _ := strconv.Atoi(value)
		values = append(values, int_value)
	}
	return values
}

func isSafeish(values []int) bool {
	// Loop over  values
	for i := 0; i < len(values); i++ {
		// Remove index i from slice
		newValues := helpers.DeleteAtIndex(values, i)
		if isSafe(newValues) {
			return true
		}
	}
	return false
}

func part1(sample bool) int {
	data := helpers.OpenInput(sample)

	split_string := strings.Split(data, "\n")
	safeCount := 0
	for _, row := range split_string {
		values := rowToValues(row)
		if isSafe(values) {
			safeCount++
		}
	}

	return safeCount
}

func part2(sample bool) int {
	data := helpers.OpenInput(sample)

	split_string := strings.Split(data, "\n")
	safeCount := 0
	for _, row := range split_string {
		values := rowToValues(row)
		if isSafe(values) || isSafeish(values) {
			safeCount++
		}
	}

	return safeCount
}
