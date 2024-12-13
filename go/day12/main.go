package main

import (
	"fmt"

	"github.com/smort/aoc2024/util"
)

func main() {
	part1("example")
	part1("example2")
	part1("input")
	part2("example")
	part2("example2")
	part2("input")
}

var above = Point{X: 0, Y: -1}
var right = Point{X: 1, Y: 0}
var below = Point{X: 0, Y: 1}
var left = Point{X: -1, Y: 0}
var directions = []Point{above, right, below, left}

type Point struct {
	X   int
	Y   int
	Val string
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

func part1(filename string) {
	lines := util.GetLines(filename)
	points := make([]Point, 0, len(lines)*len(lines[0]))
	for y, line := range lines {
		for x, c := range line {
			p := Point{Y: y, X: x, Val: string(c)}
			points = append(points, p)
		}
	}

	boundX := len(lines[0])
	boundY := len(lines)

	// make adjacency map
	size := len(points)
	adjMap := make(map[Point][]Point, size)
	for _, point := range points {
		neighbors := []Point{}
		for _, direction := range directions {
			possNeighbor := point.Add(direction)
			if possNeighbor.IsValid(boundX, boundY) {
				possNeighbor.Val = string(lines[possNeighbor.Y][possNeighbor.X])
				if possNeighbor.Val == point.Val {
					neighbors = append(neighbors, possNeighbor)
				}
			}
		}

		adjMap[point] = neighbors
	}

	result := 0
	visited := make(map[Point]struct{}, size)
	for p := range adjMap {
		area, perimeter := visit(p, visited, adjMap, 0, 0)
		result += area * perimeter
	}

	fmt.Println(result)
}

func part2(filename string) {
	// every time it turns there is another side?
	lines := util.GetLines(filename)
	points := make([]Point, 0, len(lines)*len(lines[0]))
	for y, line := range lines {
		for x, c := range line {
			p := Point{Y: y, X: x, Val: string(c)}
			points = append(points, p)
		}
	}

	boundX := len(lines[0])
	boundY := len(lines)

	// make adjacency map
	size := len(points)
	adjMap := make(map[Point][]Point, size)
	for _, point := range points {
		neighbors := []Point{}
		for _, direction := range directions {
			possNeighbor := point.Add(direction)
			if possNeighbor.IsValid(boundX, boundY) {
				possNeighbor.Val = string(lines[possNeighbor.Y][possNeighbor.X])
				if possNeighbor.Val == point.Val {
					neighbors = append(neighbors, possNeighbor)
				}
			}
		}

		adjMap[point] = neighbors
	}

	result := 0
	visited := make(map[Point]struct{}, size)
	for p := range adjMap {
		area, corners := visit2(p, visited, adjMap, lines, 0, 0)
		result += area * corners
	}

	fmt.Println(result)
}

func visit(p Point, visited map[Point]struct{}, adjMap map[Point][]Point, area int, perimeter int) (int, int) {
	if _, exists := visited[p]; exists {
		return area, perimeter
	}

	visited[p] = struct{}{}
	neighbors := adjMap[p]
	area++
	perimeter += 4 - len(neighbors)
	for _, neighbor := range neighbors {
		area, perimeter = visit(neighbor, visited, adjMap, area, perimeter)
	}

	return area, perimeter
}

func visit2(p Point, visited map[Point]struct{}, adjMap map[Point][]Point, lines []string, area int, corners int) (int, int) {
	if _, exists := visited[p]; exists {
		return area, corners
	}

	visited[p] = struct{}{}
	neighbors := adjMap[p]
	area++
	if ok, count := isCorner(p, lines); ok {
		corners += count
	}

	for _, neighbor := range neighbors {
		area, corners = visit2(neighbor, visited, adjMap, lines, area, corners)
	}

	return area, corners
}

func isCorner(p0 Point, grid []string) (bool, int) {
	boundsX := len(grid[0])
	boundsY := len(grid)
	addWithVal := func(point Point) Point {
		newPoint := p0.Add(point)
		if newPoint.IsValid(boundsX, boundsY) {
			newPoint.Val = string(grid[newPoint.Y][newPoint.X])
		}

		return newPoint
	}
	isNotSame := func(point Point) bool {
		if !point.IsValid(boundsX, boundsY) {
			return true
		}
		return point.Val != p0.Val
	}

	decision := false
	count := 0

	// upper right
	p1 := addWithVal(above)
	p2 := addWithVal(right)
	if isNotSame(p1) && isNotSame(p2) {
		decision = true
		count++
	}

	// upper left
	p1 = addWithVal(above)
	p2 = addWithVal(left)
	if isNotSame(p1) && isNotSame(p2) {
		decision = true
		count++
	}

	// bottom right
	p1 = addWithVal(below)
	p2 = addWithVal(right)
	if isNotSame(p1) && isNotSame(p2) {
		decision = true
		count++
	}

	// bottom left
	p1 = addWithVal(below)
	p2 = addWithVal(left)
	if isNotSame(p1) && isNotSame(p2) {
		decision = true
		count++
	}

	// concave corners
	// bottom right
	p1 = addWithVal(below)
	p2 = addWithVal(right)
	if p1.Val == p0.Val && p2.Val == p0.Val {
		p3 := addWithVal(Point{Y: 1, X: 1})
		if p3.Val != "" && p3.Val != p0.Val {
			decision = true
			count++
		}
	}

	// upper right
	p1 = addWithVal(above)
	p2 = addWithVal(right)
	if p1.Val == p0.Val && p2.Val == p0.Val {
		p3 := addWithVal(Point{Y: -1, X: 1})
		if p3.Val != "" && p3.Val != p0.Val {
			decision = true
			count++
		}
	}

	// upper left
	p1 = addWithVal(above)
	p2 = addWithVal(left)
	if p1.Val == p0.Val && p2.Val == p0.Val {
		p3 := addWithVal(Point{Y: -1, X: -1})
		if p3.Val != "" && p3.Val != p0.Val {
			decision = true
			count++
		}
	}

	// bottom left
	p1 = addWithVal(below)
	p2 = addWithVal(left)
	if p1.Val == p0.Val && p2.Val == p0.Val {
		p3 := addWithVal(Point{Y: 1, X: -1})
		if p3.Val != "" && p3.Val != p0.Val {
			decision = true
			count++
		}
	}

	return decision, count
}
