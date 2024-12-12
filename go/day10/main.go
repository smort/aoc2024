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
	adjMap, trailheads := readInput(filename)
	for _, trailhead := range trailheads {
		traverseUniq(adjMap, trailhead, make(map[Point]struct{}), &result)
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0
	adjMap, trailheads := readInput(filename)
	for _, trailhead := range trailheads {
		traverse(adjMap, trailhead, &result)
	}

	fmt.Println(result)
}

func traverseUniq(adjMap map[Point][]Point, curr Point, visited map[Point]struct{}, counter *int) {
	if curr.Val == 9 {
		if _, exists := visited[curr]; !exists {
			visited[curr] = struct{}{}
			*counter++
		}
		return
	}

	if neighbors, exists := adjMap[curr]; !exists || len(neighbors) == 0 {
		return
	} else {
		for _, neighbor := range neighbors {
			traverseUniq(adjMap, neighbor, visited, counter)
		}
	}
}

func traverse(adjMap map[Point][]Point, curr Point, counter *int) {
	if curr.Val == 9 {
		*counter++
		return
	}

	if neighbors, exists := adjMap[curr]; !exists || len(neighbors) == 0 {
		return
	} else {
		for _, neighbor := range neighbors {
			traverse(adjMap, neighbor, counter)
		}
	}
}

type Point struct {
	X   int
	Y   int
	Val int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) IsValid(boundX, boundY int) bool {
	return p.X > -1 && p.X < boundX && p.Y > -1 && p.Y < boundY
}

func readInput(filename string) (map[Point][]Point, []Point) {
	lines := util.GetLinesTransformed[[]string](filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})

	boundsX := len(lines[0])
	boundsY := len(lines)
	directions := []Point{Point{X: 0, Y: -1}, Point{X: 1, Y: 0}, Point{X: 0, Y: 1}, Point{X: -1, Y: 0}} // N, E, S, W
	adjMap := make(map[Point][]Point, boundsY*boundsX)
	trailheads := make([]Point, 0)

	for y, row := range lines {
		for x, cell := range row {
			cellValue := util.MustConvAtoi(cell)
			p := Point{X: x, Y: y, Val: cellValue}

			neighbors := make([]Point, 0)
			for _, direction := range directions {
				neighbor := p.Add(direction)
				if neighbor.IsValid(boundsX, boundsY) {
					val := util.MustConvAtoi(lines[neighbor.Y][neighbor.X])
					if val == cellValue+1 {
						neighbor.Val = val
						neighbors = append(neighbors, neighbor)
					}
				}
			}

			if cellValue == 0 {
				trailheads = append(trailheads, p)
			}

			adjMap[p] = neighbors
		}
	}

	return adjMap, trailheads
}
