package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetFileContent(name string) string {
	b, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	return string(b)
}

func SplitIntoLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func GetFileContentByLines(name string) []string {
	return SplitIntoLines(GetFileContent(name))
}

func String2Int(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

