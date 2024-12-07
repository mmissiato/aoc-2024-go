package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
)

type RowCol struct{ r, c int }

// Part 2
type RowColDir struct{ rc, dir RowCol }

func main() {
	rows := utils.GetFileContentByLines("input.txt")

	var rc RowCol
	directions := [4]RowCol{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	i := 0

	for i, row := range rows {
		for j, col := range row {
			if col == '^' {
				rc = RowCol{i, j}
			}
		}
	}

	n := len(rows)
	visited := make(map[RowCol]bool)
	visited[rc] = true

	// Part 2
	start := rc
	obstacles := make(map[RowCol]bool)
	var isLoop = func(obstacle RowCol) bool {
		visitedLoop := make(map[RowColDir]bool)
		rc := start
		i := 0

		for rc.r >= 0 && rc.r < n && rc.c >= 0 && rc.c < n {
			direction := directions[i]
			if rows[rc.r][rc.c] == '#' || rc.r == obstacle.r && rc.c == obstacle.c {
				rc.r -= direction.r
				rc.c -= direction.c
				i = (i + 1) % 4
			} else {
				if visitedLoop[RowColDir{rc, direction}] {
					return true
				} else {
					visitedLoop[RowColDir{rc, direction}] = true
					rc.r += direction.r
					rc.c += direction.c
				}
			}
		}

		return false
	}

	for rc.r >= 0 && rc.r < n && rc.c >= 0 && rc.c < n {
		direction := directions[i]
		if rows[rc.r][rc.c] == '#' {
			rc.r -= direction.r
			rc.c -= direction.c
			i = (i + 1) % 4
		} else {
			visited[RowCol{rc.r, rc.c}] = true
			// Part 2
			// Is this spot a candidate to stuck the guard in a loop?
			if isLoop(RowCol{rc.r, rc.c}) {
				obstacles[rc] = true
			}
			rc.r += direction.r
			rc.c += direction.c
		}
	}

	fmt.Println(len(visited))
	fmt.Println(len(obstacles))
}
