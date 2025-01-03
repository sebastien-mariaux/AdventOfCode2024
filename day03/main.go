package main

import (
	"aoc_2024/helpers"
	"regexp"
	"strconv"
)

func part1(sample bool) int {
	data := helpers.OpenInput(sample)
	return solvePart1(data)
}

func solvePart1(data string) int {
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

func solvePart2(data string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|(don't\(\))|(do\(\))`)

	// find all matches
	matches := regex.FindAllStringSubmatch(data, -1)
	count := 0
	canMultiply := true
	for _, match := range matches {
		if match[0] == "don't()" {
			canMultiply = false
			continue
		}
		if match[0] == "do()" {
			canMultiply = true
			continue
		}
		if canMultiply {
			firstVal, _ := strconv.Atoi(match[1])
			secondVal, _ := strconv.Atoi(match[2])
			count += firstVal * secondVal
		}
	}

	return count
}

func part2(sample bool) int {
	data := helpers.OpenInput(sample)
	return solvePart2(data)
}
