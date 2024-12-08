package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"strconv"
	"strings"
)

func canBeTrue(test int, numbers []int) bool {
	if len(numbers) == 1 {
		return test == numbers[0]
	}
	if canBeTrue(test, append([]int{numbers[0] + numbers[1]}, numbers[2:]...)) {
		return true
	}
	if canBeTrue(test, append([]int{numbers[0] * numbers[1]}, numbers[2:]...)) {
		return true
	}
	// Part 2
	if canBeTrue(test, append([]int{utils.String2Int(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))}, numbers[2:]...)) {
		return true
	}
	return false
}

func main() {

	equations := utils.GetFileContentByLines("input.txt")
	res := 0
	for _, equation := range equations {
		parts := strings.Split(equation, ":")

		test := utils.String2Int(parts[0])
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

		isTrue := canBeTrue(test, utils.ArrayMap(numbers, utils.String2Int))
		if isTrue {
			res += test
		}
	}

	fmt.Println(res)
}
