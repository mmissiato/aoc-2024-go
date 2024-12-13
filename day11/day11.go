// https://adventofcode.com/2024/day/11
package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"strconv"
	"strings"
)

func blink(stones *map[int]int) {
	newStones := make(map[int]int)
	for stone, count := range *stones {
		stoneStr := strconv.Itoa(stone)
		switch {
		case stone == 0:
			newStones[1] += count
		case len(stoneStr)%2 == 0:
			left := utils.String2Int(stoneStr[:len(stoneStr)/2])
			right := utils.String2Int(stoneStr[len(stoneStr)/2:])
			newStones[left] += count
			newStones[right] += count
		default:
			newStones[stone*2024] += count
		}
	}
	*stones = newStones
	return
}

func iterBlink(stones *map[int]int, iterations int) {
	if iterations == 0 {
		return
	} else {
		blink(stones)
		iterBlink(stones, iterations-1)
	}
}

func main() {
	initial := utils.GetFileContent("input.txt")
	fmt.Println(initial)

	initStones := strings.Split(initial, " ")
	var run = func(iterations int) {
		stones := make(map[int]int)

		for _, stone := range initStones {
			stones[utils.String2Int(stone)]++
		}

		iterBlink(&stones, iterations)

		totalCount := 0
		for _, count := range stones {
			totalCount += count
		}
		fmt.Println(totalCount)
	}

	run(25)
	run(75)
}
