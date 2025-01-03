package helpers

import (
	"os"
)

func OpenInput(sample bool) string {
	var file string
	if sample {
		file = "sample.txt"
	} else {
		file = "data.txt"
	}
	content, _ := os.ReadFile(file)
	return string(content)
}

func AbsoluteDiff(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return a - b
	}
	return b - a
}

func SumList(list []int) int {
	var sum int
	for _, value := range list {
		sum += value
	}
	return sum
}
