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
	part1("input", 100)
	part2("example", 50)
	part2("input", 100)
}

func part1(filename string, requiredSavings int) {
	const cheatLength = 2

	// do djikstra's and get the fastest path, with the number of steps at each path
	y := 0
	var start, end util.Point
	maze := util.GetLinesTransformed[[]string](filename, func(s string) ([]string, error) {
		vals := strings.Split(s, "")
		for x, val := range vals {
			if val == "S" {
				start = util.Point{
					X: x,
					Y: y,
				}
			} else if val == "E" {
				end = util.Point{
					X: x,
					Y: y,
				}
			}
		}
		y++
		return vals, nil
	})

	_, path := djikstra(maze, start, end)

	// for step on the path, if there is a wall next to it, check if there is another path piece that is more of a shortcut
	cheats := map[util.Point]int{}
	result := 0
	for p, d := range path {
		for _, direction := range util.Directions {
			poss := p.Add(direction)
			if _, exists := cheats[poss]; exists || maze[poss.Y][poss.X] != "#" {
				continue
			}

			for _, direction := range util.Directions {
				maybePath := poss.Add(direction)
				if maybePath == p {
					continue
				}
				if cheatEnd, exists := path[maybePath]; exists {
					cheatSave := cheatEnd - d - cheatLength // need to account for travel distance through the wall aka cheat length
					if cheatSave >= requiredSavings {
						cheats[poss] = cheatEnd - d
						result++
					}
				}
			}
		}
	}

	fmt.Println(result)
}

type cheat struct {
	start util.Point
	end   util.Point
}

func part2(filename string, requiredSavings int) {
	const cheatLength = 20

	y := 0
	var start, end util.Point
	maze := util.GetLinesTransformed[[]string](filename, func(s string) ([]string, error) {
		vals := strings.Split(s, "")
		for x, val := range vals {
			if val == "S" {
				start = util.Point{
					X: x,
					Y: y,
				}
			} else if val == "E" {
				end = util.Point{
					X: x,
					Y: y,
				}
			}
		}
		y++
		return vals, nil
	})

	_, path := djikstra(maze, start, end)

	result := 0
	cheats := map[cheat]struct{}{}
	for p0, d0 := range path {
		for p1, d1 := range path {
			c := cheat{start: p0, end: p1}
			if _, exists := cheats[c]; exists {
				continue
			}
			travelDistance := manhattanDistance(p0, p1)
			// d0 is basically cheat start and d1 we'll assume is cheat end
			if d1-d0-travelDistance >= requiredSavings {
				if travelDistance <= cheatLength { // cant travel more than our cheat rule would allow
					cheats[c] = struct{}{}
					result++
				}
			}
		}
	}

	fmt.Println(result)
}

type journey struct {
	pos  util.Point
	path map[util.Point]int
}

func djikstra(maze [][]string, start util.Point, end util.Point) (int, map[util.Point]int) {
	pq := util.NewMinHeap[journey]()
	pq.Push(journey{start, make(map[util.Point]int)}, 0)

	grid := util.Grid{
		LenX: len(maze[0]),
		LenY: len(maze),
	}

	visited := make(map[util.Point]struct{}, 0)
	visited[start] = struct{}{}
	for pq.Len() > 0 {
		i := pq.PopItem()
		score := i.Priority
		j := i.Value

		j.path[j.pos] = score

		if j.pos == end {
			return score, j.path
		}

		newCount := score + 1
		for _, direction := range util.Directions {
			poss := j.pos.Add(direction)
			if !poss.In(grid) {
				continue
			}
			if maze[poss.Y][poss.X] == "#" {
				continue
			}
			if _, exists := visited[poss]; exists {
				continue
			}

			visited[poss] = struct{}{}
			// would run faster if we just kept the prev point, but then we have to reconstruct it and the runtime is fine
			pq.Push(journey{poss, maps.Clone(j.path)}, newCount)
		}
	}

	return -1, nil
}

func manhattanDistance(p0, p1 util.Point) int {
	return int(math.Abs(float64(p0.Y-p1.Y)) + math.Abs(float64(p0.X-p1.X)))
}
