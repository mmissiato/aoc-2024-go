package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"strconv"
	"strings"
)

func canBeTrue(test int, numbers []int, part2 bool) bool {
	if len(numbers) == 1 {
		return test == numbers[0]
	}
	if canBeTrue(test, append([]int{numbers[0] + numbers[1]}, numbers[2:]...), part2) {
		return true
	}
	if canBeTrue(test, append([]int{numbers[0] * numbers[1]}, numbers[2:]...), part2) {
		return true
	}
	// Part 2
	if part2 && canBeTrue(test, append([]int{utils.String2Int(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))}, numbers[2:]...), part2) {
		return true
	}
	return false
}

func main() {

	equations := utils.GetFileContentByLines("input.txt")
	res1 := 0
	res2 := 0
	for _, equation := range equations {
		parts := strings.Split(equation, ":")

		test := utils.String2Int(parts[0])
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

		if canBeTrue(test, utils.ArrayMap(numbers, utils.String2Int), false) {
			res1 += test
		}

		if canBeTrue(test, utils.ArrayMap(numbers, utils.String2Int), true) {
			res2 += test
		}
	}

	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
}
