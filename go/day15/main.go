package main

import (
	"fmt"
	"strings"

	"github.com/smort/aoc2024/util"
)

var (
	left  = Point{X: -1, Y: 0}
	right = Point{X: 1, Y: 0}
	up    = Point{X: 0, Y: -1}
	down  = Point{X: 0, Y: 1}
)

func main() {
	part1("example")
	part1("example2")
	part1("input")
	part2("example3")
	part2("example2")
	part2("input")
}

func part1(filename string) {
	result := 0
	g := make(grid)
	directions := make([]string, 0)
	readingGrid := true
	robot := Point{}
	y := 0
	util.GetLinesTransformed(filename, func(s string) (string, error) {
		if s == "" {
			readingGrid = false
			return "", nil
		}

		if readingGrid {
			for idx, c := range strings.Split(s, "") {
				g[Point{X: idx, Y: y}] = c
				if c == "@" {
					robot = Point{X: idx, Y: y}
				}
			}
			y++
			return "", nil
		}

		directions = append(directions, strings.Split(s, "")...)

		return "", nil
	})

	for _, d := range directions {
		direction := Point{}
		switch d {
		case "<":
			direction = left
		case ">":
			direction = right
		case "^":
			direction = up
		case "v":
			direction = down
		}

		if canMove(g, direction, robot) {
			move(g, direction, robot)
			robot = robot.Add(direction)
		}
	}

	for p, val := range g {
		if val != "O" {
			continue
		}

		result += 100*p.Y + p.X
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0
	g := make(grid)
	directions := make([]string, 0)
	readingGrid := true
	robot := Point{}
	y := 0
	util.GetLinesTransformed(filename, func(s string) (string, error) {
		if s == "" {
			readingGrid = false
			return "", nil
		}

		if readingGrid {
			for idx, c := range strings.Split(s, "") {
				scaledIdx := idx * 2
				g[Point{X: scaledIdx, Y: y}] = c
				if c == "@" {
					g[Point{X: scaledIdx + 1, Y: y}] = "."
					robot = Point{X: scaledIdx, Y: y}
					continue
				}

				if c == "O" {
					g[Point{X: scaledIdx, Y: y}] = "["
					g[Point{X: scaledIdx + 1, Y: y}] = "]"
				} else {
					g[Point{X: scaledIdx + 1, Y: y}] = c
				}
			}
			y++
			return "", nil
		}

		directions = append(directions, strings.Split(s, "")...)

		return "", nil
	})

	for _, d := range directions {
		direction := Point{}
		switch d {
		case "<":
			direction = left
		case ">":
			direction = right
		case "^":
			direction = up
		case "v":
			direction = down
		}

		if canMove(g, direction, robot) {
			move(g, direction, robot)
			robot = robot.Add(direction)
		}
	}

	for p, val := range g {
		if val != "[" {
			continue
		}

		result += 100*p.Y + p.X
	}

	fmt.Println(result)
}

func move(g grid, direction Point, curr Point) {
	newPos := curr.Add(direction)
	newValue := g[newPos]
	if newValue == "#" {
		return
	}

	if newValue == "O" {
		move(g, direction, newPos)
	}

	if newValue == "[" || newValue == "]" {
		if direction == up || direction == down {
			moveBox(g, direction, newPos)
		} else {
			move(g, direction, newPos)
		}
	}

	g[newPos] = g[curr]
	g[curr] = "."
}

func moveBox(g grid, direction Point, curr Point) {
	left, right := getBoxPos(g, curr)

	newLeft := left.Add(direction)
	newRight := right.Add(direction)
	newLeftValue := g[newLeft]
	newRightValue := g[newRight]

	if newLeftValue == "#" || newRightValue == "#" {
		return
	}

	if newLeftValue == "[" && newRightValue == "]" {
		moveBox(g, direction, newLeft)
	} else {
		if newLeftValue != "." {
			moveBox(g, direction, newLeft)
		}
		if newRightValue != "." {
			moveBox(g, direction, newRight)
		}
	}

	g[newLeft] = g[left]
	g[left] = "."
	g[newRight] = g[right]
	g[right] = "."
}

func canMove(g grid, direction Point, curr Point) bool {
	newPos := curr.Add(direction)
	newValue := g[newPos]
	isMovable := true
	switch newValue {
	case ".":
		isMovable = true
	case "#":
		isMovable = false
	case "O":
		isMovable = canMove(g, direction, newPos)
	case "[", "]":
		if direction == left || direction == right {
			isMovable = canMove(g, direction, newPos)
		} else {
			left, right := getBoxPos(g, newPos)
			isMovable = canMove(g, direction, left) && canMove(g, direction, right)
		}
	}

	return isMovable
}

func getBoxPos(g grid, side Point) (Point, Point) {
	currValue := g[side]
	var l, r Point
	if currValue == "[" {
		l = side
		r = side.Add(right)
	} else {
		l = side.Add(left)
		r = side
	}

	return l, r
}

func draw(g grid, x, y int) {
	for yy := range y {
		for xx := range x {
			if v, exists := g[Point{X: xx, Y: yy}]; exists {
				fmt.Print(v)
				continue
			}

			fmt.Print(".")
		}
		fmt.Println()
	}
}

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

type grid map[Point]string
