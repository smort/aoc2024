package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
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
	grid := util.GetLinesTransformed[[]int](filename, func(s string) ([]int, error) {
		vals := strings.Split(s, " ")
		converted := make([]int, len(vals))
		for idx, val := range vals {
			i, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			converted[idx] = i
		}

		return converted, nil
	})

	numValid := 0
	for _, report := range grid {
		valid := isValid(report, getDirection(report[0], report[1]), 0)
		if valid {
			numValid++
		}
	}

	fmt.Println(numValid)
}

func part2(filename string) {
	grid := util.GetLinesTransformed[[]int](filename, func(s string) ([]int, error) {
		vals := strings.Split(s, " ")
		converted := make([]int, len(vals))
		for idx, val := range vals {
			i, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			converted[idx] = i
		}

		return converted, nil
	})

	numValid := 0
	for _, report := range grid {
		valid := isValid(report, getDirection(report[0], report[1]), 1)
		if valid {
			numValid++
		}
	}

	fmt.Println(numValid)
}

func getDirection(prev, curr int) string {
	if prev-curr < 0 {
		return "asc"
	}

	return "desc"
}

func isValid(nums []int, desiredDirection string, tolerance int) bool {
	valid := true
	i := 1
	for ; i < len(nums); i++ {
		prev := nums[i-1]
		curr := nums[i]
		// check for same direction
		if desiredDirection != getDirection(prev, curr) {
			valid = false
			break
		}

		// check for distance greater than 0 and less than 4
		diff := math.Abs(float64(prev - curr))
		if diff < 1 || diff > 3 {
			valid = false
			break
		}
	}

	if !valid && tolerance > 0 {
		for i := 0; i < len(nums); i++ {
			withoutCurr := slices.Concat(nums[:i], nums[i+1:])
			valid = isValid(withoutCurr, getDirection(withoutCurr[0], withoutCurr[1]), tolerance-1)
			if valid {
				break
			}
		}
	}

	return valid
}
