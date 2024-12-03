package main

import (
	"fmt"
	"math"
	"mmissiato/aoc-2024-go/utils"
	"slices"
	"strings"
)

func isSafeReport(report string) bool {
	levels := strings.Split(report, " ")

	return isSafeReportInner(levels)
}

func isSafeReportInner(levels []string) bool {
	increasing := false

	for i := 0; i < len(levels)-1; i++ {
		currentLevel := utils.String2Int(levels[i])
		nextLevel := utils.String2Int(levels[i+1])

		// It would be better to compare sorted slices.
		if i == 0 {
			increasing = nextLevel > currentLevel
		}
		diff := math.Abs(float64(nextLevel) - float64(currentLevel))
		if increasing != (nextLevel > currentLevel) || diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

// Tolerate a single bad level, i.e., it doesn't exist.
func isSafeReportDampener(report string) bool {
	levels := strings.Split(report, " ")

	for i := 0; i < len(levels); i++ {
		levelsdumpener := make([]string, len(levels)-1)
		levelsdumpener = slices.Delete(slices.Clone(levels), i, i+1)
		if isSafeReportInner(levelsdumpener) {
			return true
		}
	}

	return false
}

func main() {

	reports := utils.GetFileContentByLines("input.txt")
	safeReports := 0

	// Part 1
	for _, report := range reports {
		if isSafeReport(report) {
			safeReports++
		}
	}

	fmt.Println(safeReports)

	// Part 2
	safeReports = 0

	for _, report := range reports {
		if isSafeReportDampener(report) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

