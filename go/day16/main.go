package main

import (
	"fmt"
	"maps"
	"strings"

	"github.com/smort/aoc2024/util"
)

var possibleDirections = map[util.Point][]util.Point{
	util.Left:  {util.Left, util.Up, util.Down},
	util.Up:    {util.Up, util.Left, util.Right},
	util.Right: {util.Right, util.Up, util.Down},
	util.Down:  {util.Down, util.Left, util.Right},
}

func main() {
	part1("example")
	part1("input")
	part2("example")
	part2("example2")
	part2("input")
}

func part1(filename string) {
	lines := util.GetLinesTransformed(filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})

	grid := util.Grid{LenX: len(lines[0]) - 1, LenY: len(lines) - 1}
	adjMap := make(map[util.Point][]util.Point)
	var reindeer, end util.Point
	for y, row := range lines {
		for x, val := range row {
			if val == "#" {
				continue
			}

			p := util.Point{X: x, Y: y}
			if val == "S" {
				reindeer = p
			}
			if val == "E" {
				end = p
			}

			adjMap[p] = make([]util.Point, 0)
			for _, direction := range util.Directions {
				possibleNeighbor := p.Add(direction)
				if !p.In(grid) {
					continue
				}

				val := lines[possibleNeighbor.Y][possibleNeighbor.X]
				if val == "E" || val == "." {
					adjMap[p] = append(adjMap[p], possibleNeighbor)
				}
			}
		}
	}

	score, _ := djikstra(adjMap, reindeer, end)
	fmt.Println(score)
}

type position struct {
	loc   pointWithDirection
	score int
	path  map[util.Point]int
}

// direction matters
type pointWithDirection struct {
	point     util.Point
	direction util.Point
}

func djikstra(adjMap map[util.Point][]util.Point, start, end util.Point) (int, map[util.Point]int) {
	q := util.NewMinHeap(util.Item[position]{
		Value: position{
			loc:   pointWithDirection{start, util.Right},
			score: 0,
			path:  make(map[util.Point]int),
		},
		Priority: 0,
	})

	visited := make(map[pointWithDirection]struct{}, len(adjMap))
	for q.Len() > 0 {
		curr := q.Pop()
		if _, exists := visited[curr.loc]; exists {
			continue
		}

		curr.path[curr.loc.point] = curr.score

		if curr.loc.point == end {
			return curr.score, curr.path
		}

		for _, direction := range possibleDirections[curr.loc.direction] {
			possNext := curr.loc.point.Add(direction)
			if _, exists := adjMap[possNext]; !exists {
				continue
			}

			newScore := curr.score + 1
			if direction != curr.loc.direction {
				newScore += 1000
			}

			q.Push(position{
				loc:   pointWithDirection{possNext, direction},
				score: newScore,
				path:  maps.Clone(curr.path),
			}, newScore)
		}

		visited[curr.loc] = struct{}{}
	}

	return -1, nil
}

func countExciting(adjMap map[util.Point][]util.Point, start, end util.Point, optimal map[util.Point]int) int {
	q := util.NewMinHeap(util.Item[position]{
		Value: position{
			loc:   pointWithDirection{start, util.Right},
			score: 0,
			path:  make(map[util.Point]int),
		},
		Priority: 0,
	})

	excitingTiles := make(map[util.Point]struct{})
	visited := make(map[pointWithDirection]struct{}, len(adjMap))
	for q.Len() > 0 {
		curr := q.Pop()

		curr.path[curr.loc.point] = curr.score

		// if we get to a point, and it has the same score as that point on the optimal path, then this must be a valid path too
		if score, exists := optimal[curr.loc.point]; exists && score == curr.score {
			for p := range curr.path {
				excitingTiles[p] = struct{}{}
			}
		}

		if curr.loc.point == end {
			continue
		}

		for _, direction := range possibleDirections[curr.loc.direction] {
			possNext := curr.loc.point.Add(direction)
			if _, exists := visited[pointWithDirection{possNext, direction}]; exists {
				continue
			}
			if _, exists := adjMap[possNext]; !exists {
				continue
			}

			newScore := curr.score + 1
			if direction != curr.loc.direction {
				newScore += 1000
			}

			q.Push(position{
				loc:   pointWithDirection{possNext, direction},
				score: newScore,
				path:  maps.Clone(curr.path),
			}, newScore)
		}

		visited[curr.loc] = struct{}{}
	}

	return len(excitingTiles)
}

func getOpposite(d util.Point) util.Point {
	return util.Point{X: -d.X, Y: -d.Y}
}

func part2(filename string) {
	lines := util.GetLinesTransformed(filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})

	grid := util.Grid{LenX: len(lines[0]) - 1, LenY: len(lines) - 1}
	adjMap := make(map[util.Point][]util.Point)
	var reindeer, end util.Point
	for y, row := range lines {
		for x, val := range row {
			if val == "#" {
				continue
			}

			p := util.Point{X: x, Y: y}
			if val == "S" {
				reindeer = p
			}
			if val == "E" {
				end = p
			}

			adjMap[p] = make([]util.Point, 0)
			for _, direction := range util.Directions {
				possibleNeighbor := p.Add(direction)
				if !p.In(grid) {
					continue
				}

				val := lines[possibleNeighbor.Y][possibleNeighbor.X]
				if val == "E" || val == "." {
					adjMap[p] = append(adjMap[p], possibleNeighbor)
				}
			}
		}
	}

	_, optimal := djikstra(adjMap, reindeer, end)
	count := countExciting(adjMap, reindeer, end, optimal)
	fmt.Println(count)
}

func draw(grid [][]string, points map[util.Point]struct{}) {
	for y, row := range grid {
		fmt.Printf("%3d: ", y)
		for x, val := range row {
			if _, exists := points[util.Point{X: x, Y: y}]; exists {
				fmt.Print("O")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println()
	}
}
