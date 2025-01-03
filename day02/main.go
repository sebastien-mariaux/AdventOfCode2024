package main

import (
	"aoc_2024/helpers"
	"strconv"
	"strings"
)

func isSafe(row string) bool {
	var values []int
	for _, value := range strings.Split(row, " ") {
		int_value, _ :=strconv.Atoi(value)
		values = append(values, int_value)
	}

	isDescending := values[len(values)-1] < values[0]
	isAscending := values[len(values)-1] > values[0]

	if !(isDescending || isAscending) {
		return false
	}

	for i := 1; i < len(values); i++ {
		if isDescending && values[i] > values[i-1] {
			return false
		}
		if isAscending && values[i] < values[i-1] {
			return false
		}
		diff := helpers.AbsoluteDiff(values[i], values[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func part1(sample bool) int {
	data := helpers.OpenInput(sample)

	split_string := strings.Split(data, "\n")
	safeCount := 0
	for _, row := range split_string {
		if isSafe(row) {
			safeCount++
		}
	}

	return safeCount
}

func part2(sample bool) int {
	return 0
}
