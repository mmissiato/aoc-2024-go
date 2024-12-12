// https://adventofcode.com/2024/day/10
package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
)

type RowCol struct{ r, c int }
type Trail struct {
	n  int
	rc RowCol
}

var directions = [4]RowCol{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func evalHikingRating(rc RowCol, topomap []string, visited map[RowCol]int) int {
	n := len(topomap)
	if string(topomap[rc.r][rc.c]) == "0" {
		return 1
	}
	if val, contains := visited[rc]; contains {
		return val
	}
	rating := 0
	for _, dr := range directions {
		nrc := RowCol{rc.r + dr.r, rc.c + dr.c}
		if nrc.r >= 0 && nrc.r < n && nrc.c >= 0 && nrc.c < n && utils.String2Int(string(topomap[nrc.r][nrc.c])) == utils.String2Int(string(topomap[rc.r][rc.c]))-1 {
			rating += evalHikingRating(nrc, topomap, visited)
		}
	}
	return rating
}

func main() {

	topomap := utils.GetFileContentByLines("input.txt")
	n := len(topomap)

	part1 := func() int {
		score := 0
		for r, row := range topomap {
			for c, ns := range row {
				if string(ns) == "9" {
					queue := make([]Trail, 0)
					queue = append(queue, Trail{9, RowCol{r, c}})
					visited := make(map[RowCol]bool)

					for len(queue) > 0 {
						var pos Trail
						pos, queue = queue[0], queue[1:]

						if visited[pos.rc] {
							continue
						}

						visited[pos.rc] = true

						if pos.n == 0 {
							score += 1
						}

						for _, dr := range directions {
							nrc := RowCol{pos.rc.r + dr.r, pos.rc.c + dr.c}
							if nrc.r >= 0 && nrc.r < n && nrc.c >= 0 && nrc.c < n && utils.String2Int(string(topomap[nrc.r][nrc.c])) == pos.n-1 {
								queue = append(queue, Trail{utils.String2Int(string(topomap[nrc.r][nrc.c])), nrc})
							}
						}
					}
				}
			}
		}
		return score
	}

	part2 := func() int {
		visited := make(map[RowCol]int, 0)
		rating := 0

		for r, row := range topomap {
			for c, ns := range row {
				if string(ns) == "9" {
					r := evalHikingRating(RowCol{r, c}, topomap, visited)
					rating += r
				}
			}
		}
		return rating
	}

	fmt.Println(part1())
	fmt.Println(part2())
}
