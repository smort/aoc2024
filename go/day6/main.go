package main

import (
	"fmt"
	"slices"
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
	lab := initializeLab(filename)

	fmt.Println(lab.Play())
}

func part2(filename string) {
	result := 0
	lab := initializeLab(filename)
	initialX, initialY := lab.Guard.X, lab.Guard.Y

	lab.Play()

	for y, row := range lab.Visited {
		for x, cell := range row {
			if !cell {
				continue
			}
			if x == initialX && y == initialY {
				continue
			}

			newLab1 := lab.Clone()
			newLab1.Grid[y][x] = "O"
			newLab1.Guard.X = initialX
			newLab1.Guard.Y = initialY
			newLab1.Guard.Direction = UP

			// keep track of squares visited, along with direction when visited
			visited := make(map[string]int)
			stuck := false
			for !stuck {
				if newLab1.IsOutOfLab(newLab1.Guard.X, newLab1.Guard.Y) {
					break
				}
				key := fmt.Sprintf("%d-%d-%d", newLab1.Guard.X, newLab1.Guard.Y, newLab1.Guard.Direction)
				if _, exists := visited[key]; exists {
					stuck = true
					break
				}

				visited[key] = 1

				isAgainst := newLab1.IsAgainstObject()
				for isAgainst {
					newLab1.Guard.Turn()
					isAgainst = newLab1.IsAgainstObject()
				}

				newLab1.Guard.Advance()
			}

			if stuck {
				result++
			}
		}
	}

	fmt.Println(result)
}

func initializeLab(filename string) Lab {
	lines := util.GetLinesTransformed(filename, func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})

	lab := Lab{
		Grid:    lines,
		Visited: make([][]bool, len(lines)),
	}

	for y, line := range lines {
		lab.Visited[y] = make([]bool, len(line))
		for x, c := range line {
			if c == "^" {
				lab.Guard = Guard{
					X:         x,
					Y:         y,
					Direction: UP,
				}
			}
		}
	}

	return lab
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Guard struct {
	X         int
	Y         int
	Direction Direction
}

func (p *Guard) Advance() {
	switch p.Direction {
	case UP:
		p.Y--
	case DOWN:
		p.Y++
	case RIGHT:
		p.X++
	case LEFT:
		p.X--
	}
}

func (p *Guard) Turn() {
	switch p.Direction {
	case UP:
		p.Direction = RIGHT
	case DOWN:
		p.Direction = LEFT
	case RIGHT:
		p.Direction = DOWN
	case LEFT:
		p.Direction = UP
	}
}

type Lab struct {
	Grid    [][]string
	Visited [][]bool
	Guard   Guard
}

func (l *Lab) ValueAt(x, y int) string {
	if l.IsOutOfLab(x, y) {
		return ""
	}
	return l.Grid[y][x]
}

func (l *Lab) Mark(x, y int) bool {
	if l.IsOutOfLab(x, y) {
		return false
	}

	if l.Visited[y][x] {
		return false
	}

	l.Visited[y][x] = true
	return true
}

func (l *Lab) IsAgainstObject() bool {
	switch l.Guard.Direction {
	case UP:
		val := l.ValueAt(l.Guard.X, l.Guard.Y-1)
		if val == "#" || val == "O" {
			return true
		}

	case DOWN:
		val := l.ValueAt(l.Guard.X, l.Guard.Y+1)
		if val == "#" || val == "O" {
			return true
		}
	case RIGHT:
		val := l.ValueAt(l.Guard.X+1, l.Guard.Y)
		if val == "#" || val == "O" {
			return true
		}
	case LEFT:
		val := l.ValueAt(l.Guard.X-1, l.Guard.Y)
		if val == "#" || val == "O" {
			return true
		}
	}

	return false
}

func (l *Lab) IsOutOfLab(x, y int) bool {
	if y < 0 || y > len(l.Grid)-1 {
		return true
	}

	row := l.Grid[y]
	if x < 0 || x > len(row)-1 {
		return true
	}

	return false
}

func (l *Lab) Play() int {
	result := 0
	for {
		if l.IsOutOfLab(l.Guard.X, l.Guard.Y) {
			return result
		}

		for {
			if isAgainst := l.IsAgainstObject(); !isAgainst {
				break
			}
			l.Guard.Turn()
		}

		if l.Mark(l.Guard.X, l.Guard.Y) {
			result++
		}

		l.Guard.Advance()
	}
}

func (l *Lab) Clone() Lab {
	cloned := Lab{
		Grid:    slices.Clone(l.Grid),
		Visited: make([][]bool, len(l.Grid)),
		Guard:   l.Guard,
	}

	for i := range cloned.Grid {
		cloned.Grid[i] = slices.Clone(l.Grid[i])
	}

	for i := range cloned.Visited {
		cloned.Visited[i] = make([]bool, len(l.Grid[0]))
	}

	return cloned
}
