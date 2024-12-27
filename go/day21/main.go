package main

import (
	"fmt"
	"math"
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

var numericKeypad = make(map[string]util.Point, 11)
var directionalKeypad = make(map[string]util.Point, 5)

func part1(filename string) {
	input := readInput(filename)
	initializeKeypads()

	steps := getMoves(input, 2)
	fmt.Println(steps)
}

func part2(filename string) {
	input := readInput(filename)
	initializeKeypads()

	steps := getMoves(input, 25)
	fmt.Println(steps)
}

func readInput(filename string) []string {
	return util.GetLines(filename)
}

func initializeKeypads() {
	// numeric keypad
	numericKeypad["7"] = util.Point{X: 0, Y: 3}
	numericKeypad["8"] = util.Point{X: 1, Y: 3}
	numericKeypad["9"] = util.Point{X: 2, Y: 3}
	numericKeypad["4"] = util.Point{X: 0, Y: 2}
	numericKeypad["5"] = util.Point{X: 1, Y: 2}
	numericKeypad["6"] = util.Point{X: 2, Y: 2}
	numericKeypad["1"] = util.Point{X: 0, Y: 1}
	numericKeypad["2"] = util.Point{X: 1, Y: 1}
	numericKeypad["3"] = util.Point{X: 2, Y: 1}
	numericKeypad["0"] = util.Point{X: 1, Y: 0}
	numericKeypad["A"] = util.Point{X: 2, Y: 0}

	// directional keypad
	directionalKeypad["^"] = util.Point{X: 1, Y: 1}
	directionalKeypad["A"] = util.Point{X: 2, Y: 1}
	directionalKeypad["<"] = util.Point{X: 0, Y: 0}
	directionalKeypad["v"] = util.Point{X: 1, Y: 0}
	directionalKeypad[">"] = util.Point{X: 2, Y: 0}
}

func getNumericPath(code []string, start string) []string {
	curr := numericKeypad[start]
	moves := make([]string, 0)

	for _, c := range code {
		dest := numericKeypad[c]

		diff := dest.Sub(curr)

		// figure out if we need to move < or >
		var horizontalMovement []string
		for i := 0; i < int(math.Abs(float64(diff.X))); i++ {
			if diff.X >= 0 {
				horizontalMovement = append(horizontalMovement, ">")
			} else {
				horizontalMovement = append(horizontalMovement, "<")
			}
		}

		// figure out if we need to move ^ or v
		var verticalMovement []string
		for i := 0; i < int(math.Abs(float64(diff.Y))); i++ {
			if diff.Y >= 0 {
				verticalMovement = append(verticalMovement, "^")
			} else {
				verticalMovement = append(verticalMovement, "v")
			}
		}

		// need to special case so we avoid the corners where the empty space is (0,0) basically
		if curr.Y == 0 && dest.X == 0 {
			moves = append(moves, verticalMovement...)
			moves = append(moves, horizontalMovement...)
		} else if curr.X == 0 && dest.Y == 0 {
			moves = append(moves, horizontalMovement...)
			moves = append(moves, verticalMovement...)
		} else if diff.X < 0 {
			moves = append(moves, horizontalMovement...)
			moves = append(moves, verticalMovement...)
		} else if diff.X >= 0 {
			moves = append(moves, verticalMovement...)
			moves = append(moves, horizontalMovement...)
		}

		curr = dest
		moves = append(moves, "A")
	}

	return moves
}

func getDirectionalPath(input []string, start string) []string {
	curr := directionalKeypad[start]
	moves := make([]string, 0)

	for _, char := range input {
		dest := directionalKeypad[char]
		diff := dest.Sub(curr)

		horizontal, vertical := []string{}, []string{}

		for i := 0; i < int(math.Abs(float64(diff.X))); i++ {
			if diff.X >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < int(math.Abs(float64(diff.Y))); i++ {
			if diff.Y >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		// need to special case so we avoid the corners where the empty space is (0,1) basically
		if curr.X == 0 && dest.Y == 1 {
			moves = append(moves, horizontal...)
			moves = append(moves, vertical...)
		} else if curr.Y == 1 && dest.X == 0 {
			moves = append(moves, vertical...)
			moves = append(moves, horizontal...)
		} else if diff.X < 0 {
			moves = append(moves, horizontal...)
			moves = append(moves, vertical...)
		} else if diff.X >= 0 {
			moves = append(moves, vertical...)
			moves = append(moves, horizontal...)
		}
		curr = dest
		moves = append(moves, "A")
	}

	return moves
}

func getMoves(input []string, robots int) int {
	count := 0
	cache := make(map[string][]int)
	for _, line := range input {
		row := strings.Split(line, "")
		path := getNumericPath(row, "A")
		num := getCountAfterRobots(path, robots, 1, cache)
		codeNum := ""
		for _, c := range line {
			if _, err := strconv.Atoi(string(c)); err != nil {
				break
			}

			codeNum += string(c)
		}
		count += util.MustConvAtoi(codeNum) * num
	}
	return count
}

// needed a hint on this for more robots. completely rewrote. also didn't help i had a bug in sub
func getCountAfterRobots(input []string, maxRobots int, robot int, memo map[string][]int) int {
	inputCode := strings.Join(input, "")
	if val, ok := memo[inputCode]; ok {
		if val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		memo[inputCode] = make([]int, maxRobots)
	}

	path := getDirectionalPath(input, "A")
	memo[inputCode][0] = len(path)

	if robot == maxRobots {
		return len(path)
	}

	seqOfSteps := getSteps(path)

	count := 0
	for _, seq := range seqOfSteps {
		c := getCountAfterRobots(seq, maxRobots, robot+1, memo)
		if _, ok := memo[strings.Join(seq, "")]; !ok {
			memo[strings.Join(seq, "")] = make([]int, maxRobots)
		}
		memo[strings.Join(seq, "")][0] = c
		count += c
	}

	memo[strings.Join(input, "")][robot-1] = count
	return count
}

func getSteps(input []string) [][]string {
	output := [][]string{}
	current := []string{}
	for _, char := range input {
		current = append(current, char)

		if char == "A" {
			output = append(output, current)
			current = []string{}
		}
	}
	return output
}
