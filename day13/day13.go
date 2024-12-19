// https://adventofcode.com/2024/day/13
package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"strings"
)

func main() {
	puzzle := utils.GetFileContent("input.txt")

	machines := strings.Split(puzzle, "\n\n")
	tokensPart1 := 0
	tokensPart2 := 0

	for _, machine := range machines {
		parts := strings.Split(machine, "\n")
		var buttonAX, buttonAY, buttonBX, buttonBY, prizeX, prizeY int
		fmt.Sscanf(parts[0], "Button A: X+%d, Y+%d", &buttonAX, &buttonAY)
		fmt.Sscanf(parts[1], "Button B: X+%d, Y+%d", &buttonBX, &buttonBY)
		fmt.Sscanf(parts[2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)

		/*
			PX = A*AX + B*BX
			PY = A*AY + B*BY

			B = (AX*PY - AY*PX) / (BY*AX-BX*AY)
			A = (PX - B*BX) / AX
		*/

		if (buttonAX*prizeY-buttonAY*prizeX)%(buttonBY*buttonAX-buttonBX*buttonAY) == 0 {
			timesB := (buttonAX*prizeY - buttonAY*prizeX) / (buttonBY*buttonAX - buttonBX*buttonAY)
			timesA := (prizeX - timesB*buttonBX) / buttonAX

			if timesA >= 0 && timesA <= 100 && timesB >= 0 && timesB <= 100 {
				tokensPart1 += timesA*3 + timesB
			}
		}

		// Part 2
		prizeX += 10000000000000
		prizeY += 10000000000000

		if (buttonAX*prizeY-buttonAY*prizeX)%(buttonBY*buttonAX-buttonBX*buttonAY) == 0 {
			timesB := (buttonAX*prizeY - buttonAY*prizeX) / (buttonBY*buttonAX - buttonBX*buttonAY)
			timesA := (prizeX - timesB*buttonBX) / buttonAX

			if timesA >= 0 && timesA > 100 && timesB >= 0 && timesB > 100 {
				tokensPart2 += timesA*3 + timesB
			}
		}
	}

	fmt.Println(tokensPart1)
	fmt.Println(tokensPart2)
}
