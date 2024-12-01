package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	b, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	var leftIds []int
	var rightIds []int

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		ids := strings.Split(line, "   ")

		leftId, err := strconv.Atoi(ids[0])
		if err == nil {
			leftIds = append(leftIds, leftId)
		}

		rigthId, err := strconv.Atoi(ids[1])
		if err == nil {
			rightIds = append(rightIds, rigthId)
		}

		slices.Sort(leftIds)
		slices.Sort(rightIds)
	}

	// Part 1
	distance := 0
	for i := 0; i < len(lines); i++ {
		distance += int(math.Abs(float64(leftIds[i]) - float64(rightIds[i])))
	}

	fmt.Println(distance)

	// Part 2
	counter := make(map[int]int)
	for _, id := range rightIds {
		counter[id]++
	}

	similarity := 0
	for _, id := range leftIds {
		similarity += counter[id] * id
	}

	fmt.Println(similarity)
}
