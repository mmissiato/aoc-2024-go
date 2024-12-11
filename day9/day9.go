package main

import (
	"fmt"
	"mmissiato/aoc-2024-go/utils"
	"strconv"
)

type Block struct {
	c           string
	start, size int
}

func main() {
	diskmap := utils.GetFileContent("input3.txt")
	runPart1(diskmap)
	runPart2(diskmap)
}

func makeDiskMap(diskmap string) []string {

	var fragmentedmap []string
	fileId := 0

	for i := 0; i < len(diskmap); i++ {
		c := utils.String2Int(string(diskmap[i]))
		if i%2 == 0 {
			for j := 0; j < c; j++ {
				fragmentedmap = append(fragmentedmap, strconv.Itoa(fileId))
			}
			fileId += 1
		} else {
			for j := 0; j < c; j++ {
				fragmentedmap = append(fragmentedmap, ".")
			}
		}
	}
	return fragmentedmap
}

func runPart1(diskmap string) {
	fragmentedmap := makeDiskMap(diskmap)
	i := 0
	k := len(fragmentedmap) - 1

	for i < k {
		if fragmentedmap[i] != "." {
			i += 1
		} else {
			if fragmentedmap[k] == "." {
				k -= 1
			} else {
				fragmentedmap[i], fragmentedmap[k] = fragmentedmap[k], fragmentedmap[i]
			}
		}
	}

	checksum := 0
	for i, fileId := range fragmentedmap {
		if fileId == "." {
			break
		}
		checksum += i * utils.String2Int(fileId)
	}

	fmt.Println(checksum)

}

func runPart2(diskmap string) {
	fragmentedmap := makeDiskMap(diskmap)
	checksum := 0

	freeBlocks := make([]Block, 0)
	fileBlocks := make([]Block, 0)

	k := len(fragmentedmap) - 1
	i := 0

	for k >= 0 {
		if fragmentedmap[k] == "." {
			k -= 1
		} else {
			l := 0
			for k-l >= 0 && fragmentedmap[k-l] == fragmentedmap[k] {
				l++
			}
			fileBlocks = append(fileBlocks, Block{fragmentedmap[k], k - l + 1, l})
			k -= l
		}
		if k == 0 {
			break
		}
	}

	for i < len(fragmentedmap) {
		if fragmentedmap[i] != "." {
			i += 1
		} else {
			l := 0
			for i+l < len(fragmentedmap) && fragmentedmap[i+l] == fragmentedmap[i] {
				l++
			}
			freeBlocks = append(freeBlocks, Block{".", i, l})
			i += l
		}
	}

	fileBlockIdx := 0
	for fileBlockIdx < len(fileBlocks) {
		freeBlockIdx := 0
		for freeBlockIdx < len(freeBlocks) {
			if fileBlocks[fileBlockIdx].start > freeBlocks[freeBlockIdx].start && fileBlocks[fileBlockIdx].size <= freeBlocks[freeBlockIdx].size {
				for j := 0; j < fileBlocks[fileBlockIdx].size; j++ {
					fragmentedmap[freeBlocks[freeBlockIdx].start+j] = fileBlocks[fileBlockIdx].c
					fragmentedmap[fileBlocks[fileBlockIdx].start+j] = freeBlocks[freeBlockIdx].c
				}
				freeBlocks[freeBlockIdx].start += fileBlocks[fileBlockIdx].size
				freeBlocks[freeBlockIdx].size -= fileBlocks[fileBlockIdx].size
				break
			} else {
				freeBlockIdx += 1
			}
		}
		fileBlockIdx += 1
	}

	for i, fileId := range fragmentedmap {
		if fileId != "." {
			checksum += i * utils.String2Int(fileId)
		}
	}

	fmt.Println(checksum)
}
