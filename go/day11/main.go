package main

import (
	"fmt"
	"maps"
	"math"
	"strings"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example", 1)
	part1("example2", 25)
	// part1("input", 45) //blows up somewhere above 40
	part2("example", 1)
	part2("example2", 25)
	part2("input", 75)
}

func part1(filename string, blinks int) {
	stones := util.GetLinesTransformed[[]int](filename, func(s string) ([]int, error) {
		numStrs := strings.Split(s, " ")
		nums := make([]int, len(numStrs))
		for i := range len(numStrs) {
			nums[i] = util.MustConvAtoi(numStrs[i])
		}

		return nums, nil
	})[0]

	for range blinks {
		newStones := make([]int, 0, len(stones))
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if numDigits(stone)%2 == 0 {
				first, second := splitInt(stone, numDigits(stone))
				newStones = append(newStones, first, second)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}

	fmt.Println(len(stones))
}

func part2(filename string, blinks int) {
	stones := util.GetLinesTransformed[[]int](filename, func(s string) ([]int, error) {
		numStrs := strings.Split(s, " ")
		nums := make([]int, len(numStrs))
		for i := range len(numStrs) {
			nums[i] = util.MustConvAtoi(numStrs[i])
		}

		return nums, nil
	})[0]

	stonesMap := make(map[int]int, len(stones))
	for _, stone := range stones {
		stonesMap[stone] = util.GetOrDefault(stonesMap, stone, 0) + 1
	}

	for range blinks {
		newMap := make(map[int]int, len(stonesMap))
		for stone, count := range stonesMap {
			if stone == 0 {
				newCount := util.GetOrDefault(newMap, 1, 0)
				newMap[1] = newCount + count
			} else if numDigits(stone)%2 == 0 {
				first, second := splitInt(stone, numDigits(stone))

				newCount := util.GetOrDefault(newMap, first, 0)
				newMap[first] = newCount + count
				newCount = util.GetOrDefault(newMap, second, 0)
				newMap[second] = newCount + count
			} else {
				val := stone * 2024
				newCount := util.GetOrDefault(newMap, val, 0)
				newMap[val] = newCount + count
			}
		}

		stonesMap = newMap
	}

	result := 0
	for v := range maps.Values(stonesMap) {
		result += v
	}

	fmt.Println(result)
}

func numDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func splitInt(num int, numDigits int) (int, int) {
	// Calculate the divisor to separate the halves
	divisor := int(math.Pow10(numDigits / 2))

	firstHalf := num / divisor
	secondHalf := num % divisor

	return firstHalf, secondHalf
}
