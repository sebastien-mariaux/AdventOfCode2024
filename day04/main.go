package main

import (
	"aoc_2024/helpers"
	"fmt"
	"regexp"
	"strings"
)

func part1(sample bool) int {
	// Not a great solution, but it works
	// We make string from rows, cols and diagonals and count the number of matches for each
	fmt.Println("start")
	xmasRegex := `XMAS`
	samxRegex := `SAMX`

	input := helpers.OpenInput(sample)
	rows := strings.Split(input, "\n")
	inRowsCount := 0
	for _, row := range rows {
		// count number of matches for each row
		inRowsCount += len(regexp.MustCompile(xmasRegex).FindAllString(row, -1))
		inRowsCount += len(regexp.MustCompile(samxRegex).FindAllString(row, -1))
	}

	var cols []string
	for i := 0; i < len(rows[0]); i++ {
		col := ""
		for j := 0; j < len(rows); j++ {
			col += string(rows[j][i])
		}
		cols = append(cols, col)
	}

	incolsCount := 0
	for _, col := range cols {
		// count number of matches for each col
		incolsCount += len(regexp.MustCompile(xmasRegex).FindAllString(col, -1))
		incolsCount += len(regexp.MustCompile(samxRegex).FindAllString(col, -1))
	}

	// Diagonals top left to bottom right
	var diags []string
	for i := 0; i < len(rows); i++ {
		diag := ""
		for j := 0; j < len(rows); j++ {
			if i+j < len(rows) {
				diag += string(rows[i+j][j])
			}
		}
		diags = append(diags, diag)
	}
	for j := 1; j < len(rows[0]); j++ {
		diag := ""
		for i := 0; i < len(rows); i++ {
			if i+j < len(rows) {
				diag += string(rows[i][i+j])
			}
		}
		diags = append(diags, diag)
	}

	// Diagonals top right to bottom left
	var diags2 []string
	for i := 0; i < len(rows); i++ {
		diag := ""
		for j := 0; j < len(rows); j++ {
			if i+j < len(rows) {
				diag += string(rows[i+j][len(rows)-1-j])
			}
		}
		diags2 = append(diags2, diag)
	}
	for j := 1; j < len(rows[0]); j++ {
		diag := ""
		for i := 0; i < len(rows); i++ {
			if i+j < len(rows) {
				diag += string(rows[i][len(rows)-1-i-j])
			}
		}
		diags2 = append(diags2, diag)
	}

	inDiagsCount := 0
	for _, diag := range diags {
		// count number of matches for each diag
		inDiagsCount += len(regexp.MustCompile(xmasRegex).FindAllString(diag, -1))
		inDiagsCount += len(regexp.MustCompile(samxRegex).FindAllString(diag, -1))
	}
	for _, diag := range diags2 {
		// count number of matches for each diag
		inDiagsCount += len(regexp.MustCompile(xmasRegex).FindAllString(diag, -1))
		inDiagsCount += len(regexp.MustCompile(samxRegex).FindAllString(diag, -1))
	}

	return inRowsCount + incolsCount + inDiagsCount
}

func part2(sample bool) int {
	input := helpers.OpenInput(sample)
	rows := strings.Split(input, "\n")
	count := 0
	for i := 1; i < len(rows)-1; i++ {
		for j := 1; j < len(rows[i])-1; j++ {
			if rows[i][j] == 'A' &&
				(rows[i-1][j-1] == 'S' && rows[i+1][j+1] == 'M' || rows[i-1][j-1] == 'M' && rows[i+1][j+1] == 'S') &&
				(rows[i-1][j+1] == 'S' && rows[i+1][j-1] == 'M' || rows[i-1][j+1] == 'M' && rows[i+1][j-1] == 'S') {
				count++
			}
		}
	}

	return count
}
