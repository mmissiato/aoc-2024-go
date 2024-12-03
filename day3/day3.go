package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"regexp"
	"strings"
)

func main() {

	memory := utils.GetFileContent("input.txt")

	r := regexp.MustCompile(`mul\([\d]+,[\d]+\)`)
	muls := r.FindAllString(memory, -1)

	// Part 1
	// memory = utils.GetFileContent("example.txt")
	res := 0
	for _, mul := range muls {
		var x, y int
		fmt.Sscanf(mul, "mul(%d,%d)", &x, &y)
		res += x * y
	}

	fmt.Println(res)

	// Part 2
	//memory = utils.GetFileContent("example2.txt")
	enabled := true
	res = 0

	r = regexp.MustCompile(`mul\([\d]+,[\d]+\)|don't\(\)|do\(\)`)
	fragments := r.FindAllString(memory, -1)

	for _, fragment := range fragments {
		switch {
		case fragment == "do()":
			enabled = true
		case fragment == "don't()":
			enabled = false
		case strings.HasPrefix(fragment, "mul"):
			if enabled {
				var x, y int
				fmt.Sscanf(fragment, "mul(%d,%d)", &x, &y)
				res += x * y
			}
		}
	}

	fmt.Println(res)
}
