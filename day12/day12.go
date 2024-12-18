// https://adventofcode.com/2024/day/12
package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"slices"
)

type RowCol struct{ r, c int }
type RowColDir struct{ rc, dir RowCol }

var directions = [4]RowCol{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {

	gardensMap := utils.GetFileContentByLines("input.txt")

	n := len(gardensMap)

	part1 := func() {
		visited := make(map[RowCol]bool)
		res := 0
		for r, row := range gardensMap {
			for c, _ := range row {
				rc := RowCol{r, c}

				if visited[rc] {
					continue
				}

				queue := make([]RowCol, 0)
				queue = append(queue, rc)
				area := 0
				perimeter := 0

				for len(queue) > 0 {
					qrc := queue[0]
					queue = queue[1:]

					if visited[qrc] {
						continue
					}

					visited[qrc] = true
					area += 1

					for _, dr := range directions {
						nrc := RowCol{qrc.r + dr.r, qrc.c + dr.c}
						if nrc.r >= 0 && nrc.r < n && nrc.c >= 0 && nrc.c < n && gardensMap[nrc.r][nrc.c] == gardensMap[qrc.r][qrc.c] {
							queue = append(queue, nrc)
						} else {
							perimeter += 1
						}
					}
				}
				res += area * perimeter
			}
		}

		fmt.Println(res)
	}

	part2 := func() {
		visited := make(map[RowCol]bool)
		res := 0
		for r, row := range gardensMap {
			for c, _ := range row {
				rc := RowCol{r, c}

				if visited[rc] {
					continue
				}

				queue := make([]RowCol, 0)
				queue = append(queue, rc)
				area := 0
				sameSides := make(map[RowCol][]RowCol)

				for len(queue) > 0 {
					qrc := queue[0]
					queue = queue[1:]

					if visited[qrc] {
						continue
					}

					visited[qrc] = true
					area += 1

					for _, dr := range directions {
						nrc := RowCol{qrc.r + dr.r, qrc.c + dr.c}
						if nrc.r >= 0 && nrc.r < n && nrc.c >= 0 && nrc.c < n && gardensMap[nrc.r][nrc.c] == gardensMap[qrc.r][qrc.c] {
							queue = append(queue, nrc)
						} else {
							sameSides[dr] = append(sameSides[dr], nrc)
						}
					}
				}

				sides := 0

				for _, rcs := range sameSides {
					visitedSameSide := make(map[RowCol]bool)
					for _, rc := range rcs {
						if visitedSameSide[rc] == false {
							sides += 1
							queue = make([]RowCol, 0)
							queue = append(queue, rc)
							for len(queue) > 0 {
								qrc := queue[0]
								queue = queue[1:]
								if visitedSameSide[qrc] {
									continue
								}
								visitedSameSide[qrc] = true
								for _, dr := range directions {
									nrc := RowCol{qrc.r + dr.r, qrc.c + dr.c}
									if slices.Contains(rcs, nrc) {
										queue = append(queue, nrc)
									}
								}
							}
						}
					}
				}
				res += area * sides
			}
		}

		fmt.Println(res)
	}

	part1()
	part2()
}
