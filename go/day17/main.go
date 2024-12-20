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
	part2("example2")
	part2("input")
}

func part1(filename string) {
	a, b, c, program := readInput(filename)
	output := executeProgram(program, a, b, c)

	result := ""
	for _, i := range output {
		result += strconv.Itoa(i) + ","
	}
	fmt.Println(result[:len(result)-1])
}

func part2(filename string) {
	_, _, _, wanted := readInput(filename)
	a, b, c := 0, 0, 0
	for ; ; a++ {
		output := executeProgram(wanted, a, b, c)

		if len(wanted) > len(output) {
			a *= 2
		} else if len(output) > len(wanted) {
			a /= 2
		} else {
			if len(output) == len(wanted) {
				found := true
				for i := len(wanted) - 1; i >= 0; i-- {
					if wanted[i] != output[i] {
						found = false
						// needed BIG help on this one. it would have run for forever otherwise.
						// I guess this is hinted at in the op code descriptions...
						// https://www.reddit.com/r/adventofcode/comments/1hg38ah/comment/m2gkd6m/
						a += int(math.Pow(8, float64(i)))
						break
					}
				}

				if found {
					break // found a
				}
			}
		}
	}

	fmt.Println(a)
}

func readInput(filename string) (int, int, int, []int) {
	parseRegister := func(s string) int {
		parts := strings.Split(s, ":")
		return util.MustConvAtoi(strings.TrimSpace(parts[1]))
	}
	lines := util.GetLines(filename)
	a := parseRegister(lines[0])
	b := parseRegister(lines[1])
	c := parseRegister(lines[2])

	program := strings.Split(strings.TrimSpace(strings.Split(lines[4], ":")[1]), ",")
	programInts := make([]int, len(program))
	for idx, v := range program {
		programInts[idx] = util.MustConvAtoi(v)
	}

	return a, b, c, programInts
}

func getComboVal(combo, a, b, c int) int {
	switch combo {
	case 0, 1, 2, 3:
		return combo
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return -1
	}
}

func executeProgram(program []int, a, b, c int) []int {
	output := make([]int, 0)
	for i := 0; i < len(program); {
		opcode := program[i]
		arg := program[i+1]
		accum := 2

		switch opcode {
		case 0:
			a = a / int(math.Pow(2, float64(getComboVal(arg, a, b, c))))
		case 1:
			b = b ^ arg
		case 2:
			b = getComboVal(arg, a, b, c) % 8
		case 3:
			if a != 0 {
				accum = 0
				i = arg
			}
		case 4:
			b = b ^ c
		case 5:
			output = append(output, getComboVal(arg, a, b, c)%8)
		case 6:
			b = a / int(math.Pow(2, float64(getComboVal(arg, a, b, c))))
		case 7:
			c = a / int(math.Pow(2, float64(getComboVal(arg, a, b, c))))
		}

		i += accum
	}

	return output
}
