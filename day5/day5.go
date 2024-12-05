package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"slices"
	"strings"
)

func main() {

	input := utils.GetFileContent("input.txt")
	sections := strings.Split(input, "\n\n")

	rules := strings.Split(sections[0], "\n")
	updates := strings.Split(sections[1], "\n")

	// Not really necessary but dict in go is good to know.
	var rulesSet map[string][]string
	rulesSet = make(map[string][]string)

	for _, rule := range rules {
		ns := strings.Split(rule, "|")
		rulesSet[ns[0]] = append(rulesSet[ns[0]], ns[1])
	}

	res1 := 0
	res2 := 0

	for _, update := range updates {
		valid := true
		pages := strings.Split(update, ",")

		for i := 0; i < len(pages); i++ {
			for j := 0; j < len(pages); j++ {
				// Is the page either before for the previous or after for the next?
				prev := slices.Contains(rulesSet[pages[j]], pages[i])
				next := slices.Contains(rulesSet[pages[i]], pages[j])
				if j > i && prev || j < i && next {
					valid = false
					// Part 2. Swap the unordered pages.
					if prev {
						pages[j], pages[i] = pages[i], pages[j]
					}
					if next {
						pages[i], pages[j] = pages[j], pages[i]
					}
				}
			}
		}

		middle := utils.String2Int(pages[len(pages)/2])

		if valid {
			res1 += middle
		} else {
			res2 += middle
		}
	}

	fmt.Println(res1)
	fmt.Println(res2)
}
