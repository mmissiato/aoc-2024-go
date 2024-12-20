// https://adventofcode.com/2024/day/14
package main

import (
	"fmt"
	"math"
	"mmissiato/aoc-2024-go/utils"
)

func main() {

	robots := utils.GetFileContentByLines("input.txt")

	xs := 101
	ys := 103

	robotAtT := func(robot string, t int) (int, int) {

		var px, py, vx, vy int
		fmt.Sscanf(robot, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		px = (px + vx*t) % xs
		py = (py + vy*t) % ys

		if px < 0 {
			px = xs - int(math.Abs(float64(px)))
		}
		if px >= xs {
			px = px - xs
		}
		if py < 0 {
			py = ys - int(math.Abs(float64(py)))
		}
		if py >= ys {
			py = py - ys
		}

		return px, py
	}

	part1 := func() {
		var q1, q2, q3, q4 int

		for _, robot := range robots {
			t := 100
			px, py := robotAtT(robot, t)

			switch {
			case (px < xs/2 && py < ys/2):
				q1 += 1
			case (px > xs/2 && py < ys/2):
				q2 += 1
			case (px < xs/2 && py > ys/2):
				q3 += 1
			case (px > xs/2 && py > ys/2):
				q4 += 1
			}
		}

		fmt.Println(q1 * q2 * q3 * q4)
	}

	// Kind of cheating.
	part2 := func() {
		// for t := 1; t < 10000; t++ {
		for t := 7916; t < 7917; t++ {
			robotsMap := make([][]string, 0)
			for i := 0; i < xs; i++ {
				line := make([]string, 0)
				for j := 0; j < ys; j++ {
					line = append(line, ".")
				}
				robotsMap = append(robotsMap, line)
			}
			for _, robot := range robots {
				px, py := robotAtT(robot, t)
				robotsMap[px][py] = "#"
			}

			fmt.Println("@@@@@@ Iter", t)
			for _, line := range robotsMap {
				fmt.Println(line)
			}
			fmt.Println("@@@@@@ Iter End", t)
		}
	}

	part1()
	part2()
}
