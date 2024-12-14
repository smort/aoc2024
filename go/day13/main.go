package main

import (
	"fmt"
	"math"
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

const (
	costA = 3
	costB = 1
)

type machine struct {
	Ax     int
	Ay     int
	Bx     int
	By     int
	PrizeX int
	PrizeY int
}

func part1(filename string) {
	result := 0
	machines := readInput(filename)

	for _, machine := range machines {
		possAnsX := [][2]int{}
		for i := 0; i < 101; i++ {
			for j := 100; j > -1; j-- {
				if (machine.Ax*i)+(machine.Bx*j) == machine.PrizeX {
					possAnsX = append(possAnsX, [2]int{i, j})
				}
			}
		}

		possAnsY := [][2]int{}
		for i := 0; i < 101; i++ {
			for j := 100; j > -1; j-- {
				if (machine.Ay*i)+(machine.By*j) == machine.PrizeY {
					possAnsY = append(possAnsY, [2]int{i, j})
				}
			}
		}

		fewestTokens := math.MaxInt
		for _, ans := range possAnsX {
			if slices.Contains(possAnsY, ans) {
				tokens := ans[0]*costA + ans[1]*costB

				fewestTokens = min(fewestTokens, tokens)
			}
		}

		if fewestTokens != math.MaxInt {
			result += fewestTokens
		}
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0
	machines := readInput(filename)

	for idx, machine := range machines {
		machine.PrizeX += 10000000000000
		machine.PrizeY += 10000000000000
		machines[idx] = machine // range copies, so we need to reassign
	}

	for _, machine := range machines {
		a, b := solve(machine)

		if machine.Ax*a+machine.Bx*b == machine.PrizeX && machine.Ay*a+machine.By*b == machine.PrizeY {
			result += a*costA + b*costB
		}
	}

	fmt.Println(result)
}

// matrix approach to solving system of 2 equations
func solve(m machine) (int, int) {
	mat := [2][3]float64{
		{float64(m.Ax), float64(m.Bx), float64(m.PrizeX)},
		{float64(m.Ay), float64(m.By), float64(m.PrizeY)},
	}

	divisor := mat[0][0]
	mat[0][0] = 1 // divide by itself
	mat[0][1] /= divisor
	mat[0][2] /= divisor

	// mult top row by inverse of 1 0 then add it to the bottom row
	mult := -mat[1][0]
	mat[1][0] = 0 // we're shooting for 0, if it's not we're in big trouble
	mat[1][1] += mult * mat[0][1]
	mat[1][2] += mult * mat[0][2]

	// [1][0] will already be 0 so we can skip it
	mult = 1 / mat[1][1]
	mat[1][1] *= mult
	mat[1][2] *= mult

	b := mat[1][2]

	mat[0][2] -= b * mat[0][1]
	a := mat[0][2]

	return int(math.Round(a)), int(math.Round(b))
}

func readInput(filename string) []machine {
	parseButton := func(line string) (int, int) {
		xPos := strings.Index(line, "X+")
		commaPos := strings.Index(line, ",")
		yPos := strings.Index(line, "Y+")

		x := line[xPos+2 : commaPos]
		y := line[yPos+2:]

		return util.MustConvAtoi(x), util.MustConvAtoi(y)
	}

	parsePrize := func(line string) (int, int) {
		xPos := strings.Index(line, "X=")
		commaPos := strings.Index(line, ",")
		yPos := strings.Index(line, "Y=")

		x := line[xPos+2 : commaPos]
		y := line[yPos+2:]

		return util.MustConvAtoi(x), util.MustConvAtoi(y)
	}

	machines := make([]machine, 0)
	lines := util.GetLines(filename)
	for i := 0; i < len(lines); i += 4 {
		lineA := lines[i]
		lineB := lines[i+1]
		lineP := lines[i+2]

		aX, aY := parseButton(lineA)
		bX, bY := parseButton(lineB)
		pX, pY := parsePrize(lineP)

		machines = append(machines, machine{
			Ax:     aX,
			Ay:     aY,
			Bx:     bX,
			By:     bY,
			PrizeX: pX,
			PrizeY: pY,
		})
	}

	return machines
}
