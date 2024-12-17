package main

import (
	"fmt"
	"strings"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example", 11, 7)
	part1("input", 101, 103)
	part2("input", 101, 103)
}

func part1(filename string, x int, y int) {
	result := 0

	lines := readInput(filename)
	lobby := NewLobby(len(lines))
	for _, line := range lines {
		p := Point{
			X: line[0],
			Y: line[1],
		}
		g := Guard{
			XVelocity: line[2],
			YVelocity: line[3],
		}

		guards := util.GetOrDefault(lobby.Grid, p, []Guard{})
		lobby.Grid[p] = append(guards, g)
	}

	for range 100 {
		newLobby := NewLobby(len(lines))
		for point, guards := range lobby.Grid {
			for _, guard := range guards {
				newPoint := guard.Advance(point, x, y)
				guards := util.GetOrDefault(newLobby.Grid, newPoint, []Guard{})
				newLobby.Grid[newPoint] = append(guards, guard)
			}
		}

		lobby = newLobby
	}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	xMid := x / 2
	yMid := y / 2
	for point, guards := range lobby.Grid {
		if point.X == xMid || point.Y == yMid {
			continue
		}

		if point.X < xMid && point.Y < yMid {
			q1 += len(guards)
		} else if point.X > xMid && point.Y < yMid {
			q2 += len(guards)
		} else if point.X < xMid && point.Y > yMid {
			q3 += len(guards)
		} else if point.X > xMid && point.Y > yMid {
			q4 += len(guards)
		}
	}

	result = q1 * q2 * q3 * q4

	fmt.Println(result)
}

func part2(filename string, x, y int) {
	lines := readInput(filename)
	lobby := NewLobby(len(lines))

	for _, line := range lines {
		p := Point{
			X: line[0],
			Y: line[1],
		}
		g := Guard{
			XVelocity: line[2],
			YVelocity: line[3],
		}

		guards := util.GetOrDefault(lobby.Grid, p, []Guard{})
		lobby.Grid[p] = append(guards, g)
	}

	seconds := 0
	maxCount := 0
	maxTime := 0
	for range 20000 {
		// count the number of contiguous guards and get the max
		visited := make(map[Point]struct{}, len(lobby.Grid))
		for point := range lobby.Grid {
			count := visit(lobby.Grid, visited, point)
			if count > maxCount {
				maxCount = count
				maxTime = seconds
			}
		}

		// advance time and update grid
		seconds++
		newLobby := NewLobby(len(lines))
		for point, guards := range lobby.Grid {
			for _, guard := range guards {
				newPoint := guard.Advance(point, x, y)
				guards := util.GetOrDefault(newLobby.Grid, newPoint, []Guard{})
				newLobby.Grid[newPoint] = append(guards, guard)
			}
		}

		lobby = newLobby
	}

	fmt.Println(maxTime)
}

func visit(grid map[Point][]Guard, visited map[Point]struct{}, point Point) int {
	numGuards := 0
	if guards, exists := grid[point]; !exists {
		return 0
	} else {
		numGuards = len(guards)
	}
	if _, exists := visited[point]; exists {
		return 0
	}

	visited[point] = struct{}{}
	count := numGuards
	count += visit(grid, visited, Point{X: point.X, Y: point.Y - 1})
	count += visit(grid, visited, Point{X: point.X + 1, Y: point.Y})
	count += visit(grid, visited, Point{X: point.X, Y: point.Y + 1})
	count += visit(grid, visited, Point{X: point.X - 1, Y: point.Y})

	return count
}

func draw(grid map[Point][]Guard, x, y int) {
	for yy := range y {
		fmt.Printf("Y %d: ", yy)
		for xx := range x {
			if _, exists := grid[Point{X: xx, Y: yy}]; exists {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func readInput(filename string) [][4]int {
	return util.GetLinesTransformed[[4]int](filename, func(s string) ([4]int, error) {
		parts := strings.Split(s, " ")
		positions := strings.Split(parts[0][2:], ",")
		result := [4]int{util.MustConvAtoi(positions[0]), util.MustConvAtoi(positions[1])}

		velocities := strings.Split(parts[1][2:], ",")
		result[2] = util.MustConvAtoi(velocities[0])
		result[3] = util.MustConvAtoi(velocities[1])

		return result, nil
	})
}

type Point struct {
	X int
	Y int
}

type Guard struct {
	XVelocity int
	YVelocity int
}

func (g *Guard) Advance(starting Point, x, y int) Point {
	newX := (g.XVelocity + starting.X) % x
	newY := (g.YVelocity + starting.Y) % y

	if newX < 0 {
		newX += x
	}

	if newY < 0 {
		newY += y
	}

	return Point{
		X: newX,
		Y: newY,
	}
}

type Lobby struct {
	Grid map[Point][]Guard
}

func NewLobby(numGuards int) Lobby {
	l := Lobby{Grid: make(map[Point][]Guard, numGuards)}
	return l
}
