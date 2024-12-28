package main

import (
	"fmt"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example", 2000)
	part1("input", 2000)
}

func part1(filename string, iterations int) {
	nums := readInput(filename)
	result := 0
	for _, n := range nums {
		for range iterations {
			n = ((64 * n) ^ n)
			n = n % 16777216
			n = int(n/32) ^ n
			n = n % 16777216
			n = (n * 2048) ^ n
			n = n % 16777216
		}

		result += n
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0

	fmt.Println(result)
}

func readInput(filename string) []int {
	return util.GetLinesTransformed(filename, func(s string) (int, error) {
		return util.MustConvAtoi(s), nil
	})
}
