package day01

import (
	"aoc_2024/helpers"
	"sort"
	"strconv"
	"strings"
)

func solve(sample bool) int {
	data := helpers.OpenInput(sample)
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
	// sort arrays
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
  sort.Slice(right, func(i, j int) bool {
    return right[i] < right[j]
  })

  var distances []int
  for i := 0; i < len(left); i++ {
    var distance int
    if left[i] == right[i] {
      distance = 0
    } else if left[i] > right[i] {
      distance = left[i] - right[i]
    } else {
      distance = right[i] - left[i]
    }
    distances = append(distances, distance)
  }

  // sum all distances
  var sum int
  for _, distance := range distances {
    sum += distance
  }

	return sum
}
