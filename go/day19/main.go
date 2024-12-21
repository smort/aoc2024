package main

import (
	"fmt"
	"strings"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example")
	part1("input")
	part2("example")
	part2("input")
}

func part1(filename string) {
	result := 0

	towels, designs := readInput(filename)
	for _, design := range designs {
		if isDesignPossible(design, towels) {
			result++
		}
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0
	towels, designs := readInput(filename)

	memo := make(map[string]int)
	for _, design := range designs {
		result += getTowelComboCount(design, towels, memo)
	}

	fmt.Println(result)
}

func readInput(filename string) ([]string, []string) {
	lines := util.GetLines(filename)
	towels := strings.Split(lines[0], ", ")

	designs := make([]string, 0, len(lines))
	for _, design := range lines[2:] {
		designs = append(designs, design)
	}

	return towels, designs
}

func isDesignPossible(design string, towels []string) bool {
	if len(design) == 0 {
		return true
	}

	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			if isDesignPossible(design[len(towel):], towels) {
				return true
			}
		}
	}

	return false
}

func getTowelComboCount(design string, towels []string, memo map[string]int) int {
	if count, exists := memo[design]; exists {
		return count
	}

	count := 0
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			if len(towel) == len(design) { // DONE
				count++
				continue
			}
			count += getTowelComboCount(design[len(towel):], towels, memo)
		}
	}

	memo[design] = count
	return count
}
