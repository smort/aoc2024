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

type antenna struct {
	X     int
	Y     int
	Value string
}

func part1(filename string) {
	result := 0
	lines := util.GetLinesTransformed[[]string](filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})
	adjacencyMap := makeAdjacencyMap(lines)

	yBound := len(lines)
	xBound := len(lines[0])
	antiAntennae := make([][]bool, yBound)
	for idx, _ := range antiAntennae {
		antiAntennae[idx] = make([]bool, xBound)
	}

	for a1, adjacencies := range adjacencyMap {
		for _, a2 := range adjacencies {
			xDiff := a1.X - a2.X
			yDiff := a1.Y - a2.Y
			antiX, antiY := a2.X+-xDiff, a2.Y+-yDiff

			inBounds := antiX > -1 && antiX < xBound && antiY > -1 && antiY < yBound
			if !inBounds {
				continue
			}
			// check if we've already put an antenna there
			if antiAntennae[antiY][antiX] {
				continue
			}

			antiAntennae[antiY][antiX] = true
			lines[antiY][antiX] = "#" // for visualization
			result++
		}
	}

	for _, v := range lines {
		fmt.Println(strings.Join(v, ""))
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0
	lines := util.GetLinesTransformed[[]string](filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})
	adjacencyMap := makeAdjacencyMap(lines)

	yBound := len(lines)
	xBound := len(lines[0])
	antiAntennae := make([][]bool, yBound)
	for idx := range antiAntennae {
		antiAntennae[idx] = make([]bool, xBound)
	}

	for a1, adjacencies := range adjacencyMap {
		for _, a2 := range adjacencies {
			xDiff := a1.X - a2.X
			yDiff := a1.Y - a2.Y

			// start from a1 instead of a2 this time bc of updated rules
			antiX, antiY := a1.X+-xDiff, a1.Y+-yDiff
			for {
				inBounds := antiX > -1 && antiX < xBound && antiY > -1 && antiY < yBound
				if !inBounds {
					break
				}

				// check if we've already put an antenna there
				if !antiAntennae[antiY][antiX] {
					antiAntennae[antiY][antiX] = true
					lines[antiY][antiX] = "#" // for visualization
					result++
				}

				antiX, antiY = antiX+-xDiff, antiY+-yDiff
			}
		}
	}

	for _, v := range lines {
		fmt.Println(strings.Join(v, ""))
	}

	fmt.Println(result)
}

func makeAdjacencyMap(lines [][]string) map[antenna][]antenna {
	adjacencyMap := make(map[antenna][]antenna, 0)
	for y, line := range lines {
		for x, value := range line {
			if value == "." {
				continue
			}

			a := antenna{
				X:     x,
				Y:     y,
				Value: value,
			}

			// do the adding to other antennae first, so we don't add this antenna to itself
			adjacencies := make([]antenna, 0)
			for k, v := range adjacencyMap {
				if k.Value != a.Value {
					continue
				}
				adjacencyMap[k] = append(v, a)
				adjacencies = append(adjacencies, k)
			}

			adjacencyMap[a] = adjacencies
		}
	}

	return adjacencyMap
}
