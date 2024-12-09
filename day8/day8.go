package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"slices"
	"unicode"
)

type RowCol struct{ r, c int }

func main() {

	rows := utils.GetFileContentByLines("input.txt")
	n := len(rows)

	var run = func(part2 bool) {
		antinodes := make(map[RowCol]bool, 0)
		antennas := make(map[rune][]RowCol, 0)

		for i, row := range rows {
			for j, freq := range row {
				if unicode.IsDigit(freq) || unicode.IsLetter(freq) {
					if !slices.Contains(antennas[freq], RowCol{i, j}) {
						sameFreqRCs := antennas[freq]
						for _, rc := range sameFreqRCs {
							dr := rc.r - i
							dc := rc.c - j
							// Part 1 double the distance, Part 2 regardless of distance.
							var multiplier int
							if part2 {
								multiplier = 1
							} else {
								multiplier = 2
							}

							for {
								drm := dr * multiplier
								dcm := dc * multiplier
								antinodes1 := RowCol{i + drm, j + dcm}
								antinodes2 := RowCol{rc.r - drm, rc.c - dcm}
								if 0 <= antinodes1.r && antinodes1.r < n && 0 <= antinodes1.c && antinodes1.c < n {
									antinodes[antinodes1] = true
								}
								if 0 <= antinodes2.r && antinodes2.r < n && 0 <= antinodes2.c && antinodes2.c < n {
									antinodes[antinodes2] = true
								}
								isAntinodes1Out := antinodes1.r < 0 || antinodes1.r >= n || antinodes1.c < 0 || antinodes1.c >= n
								isAntinodes2Out := antinodes2.r < 0 || antinodes2.r >= n || antinodes2.c < 0 || antinodes2.c >= n
								if !part2 || isAntinodes1Out && isAntinodes2Out {
									break
								}
								multiplier += 1
							}
						}
						antennas[freq] = append(antennas[freq], RowCol{i, j})
					}
				}
			}
		}
		fmt.Println(len(antinodes))
	}

	run(false)
	run(true)
}
