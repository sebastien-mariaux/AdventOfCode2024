package main

import (
	"aoc_2024/helpers"
	"strconv"
	"strings"
)

func isSafe(values []int, canRemove bool) bool {
	// Check if safe without first value
	if canRemove {
		withoutFirstValue := helpers.DeleteAtIndex(values, 0)
		if isSafe(withoutFirstValue, false) {
			return true
		}
	}

	isDescending := values[len(values)-1] < values[0]
	isAscending := values[len(values)-1] > values[0]
	if !canRemove && !(isDescending || isAscending) {
		return false
	}

	// Loop over following values
	for i := 1; i < len(values); i++ {
		diff := helpers.AbsoluteDiff(values[i], values[i-1])
		if (isDescending && values[i] > values[i-1]) || (isAscending && values[i] < values[i-1]) || (diff < 1 || diff > 3) {
			newValues := helpers.DeleteAtIndex(values, i)
			if canRemove && isSafe(newValues, false) {
				return true
			}
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

func part1(sample bool) int {
	data := helpers.OpenInput(sample)

	split_string := strings.Split(data, "\n")
	safeCount := 0
	for _, row := range split_string {
		values := rowToValues(row)
		if isSafe(values, false) {
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
		if isSafe(values, true) {
			safeCount++
		}
	}

	return safeCount
}
