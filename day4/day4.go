package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"strings"
)

func main() {

	rows := utils.GetFileContentByLines("input.txt")
	n := len(rows)

	// Part 1
	res := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j+3 < n && rows[i][j] == 'X' && rows[i][j+1] == 'M' && rows[i][j+2] == 'A' && rows[i][j+3] == 'S' {
				res += 1
			}
			if j+3 < n && rows[i][j] == 'S' && rows[i][j+1] == 'A' && rows[i][j+2] == 'M' && rows[i][j+3] == 'X' {
				res += 1
			}
			if i+3 < n && rows[i][j] == 'X' && rows[i+1][j] == 'M' && rows[i+2][j] == 'A' && rows[i+3][j] == 'S' {
				res += 1
			}
			if i+3 < n && rows[i][j] == 'S' && rows[i+1][j] == 'A' && rows[i+2][j] == 'M' && rows[i+3][j] == 'X' {
				res += 1
			}
			if i-3 >= 0 && j+3 < n && rows[i][j] == 'X' && rows[i-1][j+1] == 'M' && rows[i-2][j+2] == 'A' && rows[i-3][j+3] == 'S' {
				res += 1
			}
			if i-3 >= 0 && j+3 < n && rows[i][j] == 'S' && rows[i-1][j+1] == 'A' && rows[i-2][j+2] == 'M' && rows[i-3][j+3] == 'X' {
				res += 1
			}
			if i+3 < n && j+3 < n && rows[i][j] == 'X' && rows[i+1][j+1] == 'M' && rows[i+2][j+2] == 'A' && rows[i+3][j+3] == 'S' {
				res += 1
			}
			if i+3 < n && j+3 < n && rows[i][j] == 'S' && rows[i+1][j+1] == 'A' && rows[i+2][j+2] == 'M' && rows[i+3][j+3] == 'X' {
				res += 1
			}
		}
	}

	// Part 1 alternative
	/* res = 0
	parts := splitMatrix(rows)
	for _, part := range parts {
		res += countXMAS(part)
	}*/

	fmt.Println(res)

	// Part 2
	res = 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+2 < n && j+2 < n {
				if rows[i][j] == 'M' && rows[i][j+2] == 'S' && rows[i+1][j+1] == 'A' && rows[i+2][j] == 'M' && rows[i+2][j+2] == 'S' {
					res += 1
				}
				if rows[i][j] == 'S' && rows[i][j+2] == 'M' && rows[i+1][j+1] == 'A' && rows[i+2][j] == 'S' && rows[i+2][j+2] == 'M' {
					res += 1
				}
				if rows[i][j] == 'M' && rows[i][j+2] == 'M' && rows[i+1][j+1] == 'A' && rows[i+2][j] == 'S' && rows[i+2][j+2] == 'S' {
					res += 1
				}
				if rows[i][j] == 'S' && rows[i][j+2] == 'S' && rows[i+1][j+1] == 'A' && rows[i+2][j] == 'M' && rows[i+2][j+2] == 'M' {
					res += 1
				}
			}
		}
	}

	fmt.Println(res)
}

func splitMatrix(rows []string) []string {

	n := len(rows)
	// rows, cols, diagonals
	parts := rows

	cols := make([]string, n)
	for i := 0; i < len(rows); i++ {
		for j := 0; j < n; j++ {
			cols[j] += string(rows[i][j])
		}
	}
	parts = append(parts, cols...)
	// Top diagonals
	for i := 0; i < n-1; i++ {
		diagonal := ""
		for k := n - 1 - i; k < n; k++ {
			diagonal += string(rows[i-n+k+1][k])
		}
		parts = append(parts, diagonal)
		diagonal = ""
		for k := 0; k <= i; k++ {
			diagonal += string(rows[i-k][k])
		}
		parts = append(parts, diagonal)
	}

	// Bottom diagonals
	for i := n - 1; i >= 0; i-- {
		diagonal := ""
		for k := 0; k <= n-i-1; k++ {
			diagonal += string(rows[i+k][k])
		}
		parts = append(parts, diagonal)
		diagonal = ""
		for k := n - 1; k >= i; k-- {
			diagonal += string(rows[i+n-k-1][k])
		}
		parts = append(parts, diagonal)
	}

	return parts
}

func countXMAS(s string) int {
	return strings.Count(s, "XMAS") + strings.Count(s, "SAMX")
}
