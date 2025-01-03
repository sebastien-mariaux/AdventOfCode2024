package main

import (
	"aoc_2024/helpers"
	"regexp"
	"strconv"
)



func part1(sample bool) int {
	data := helpers.OpenInput(sample)

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// find all matches
	matches := regex.FindAllStringSubmatch(data, -1)
	count := 0
	for _, match := range matches {
		firstVal, _ := strconv.Atoi(match[1])
		secondVal, _ := strconv.Atoi(match[2])
		count += firstVal * secondVal
	}


	return count
}

func part2(sample bool) int {
	// data := helpers.OpenInput(sample)
	return 0
}
