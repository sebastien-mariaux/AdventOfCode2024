package main

import (
	"aoc_2024/helpers"
	"sort"
	"strconv"
	"strings"
)

func create_lists(data string) ([]int, []int) {
	split_string := strings.Split(data, "\n")
	var left []int
	var right []int

	// loop over split_string
	for _, row := range split_string {
		// trimed_string := strings.ReplaceAll(row, " ", ",")
		split_row := strings.SplitN(row, " ", -1)
		// append first item to left
		int_value, _ := strconv.Atoi(split_row[0])
		left = append(left, int_value)

		// append last item to right
		int_value, _ = strconv.Atoi(split_row[len(split_row)-1])
		right = append(right, int_value)
	}
	return left, right
}

func part1(sample bool) int {
	data := helpers.OpenInput(sample)
	left, right := create_lists(data)
	// sort arrays
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var distances []int
	for i := 0; i < len(left); i++ {
		distances = append(distances, helpers.AbsoluteDiff(left[i], right[i]))
	}

	return helpers.SumList(distances)
}

func part2(sample bool) int {
	data := helpers.OpenInput(sample)
	left, right := create_lists(data)

	var similarities []int
	for _, left_value := range left {
		occurrences := 0
		for _, right_value := range right {
			if left_value == right_value {
				occurrences++
			}
		}
		similarities = append(similarities, left_value*occurrences)
	}
	return helpers.SumList(similarities)
}
