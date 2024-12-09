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

func part1(filename string) {
	result := 0
	equations := readInput(filename)
	for _, eq := range equations {
		// -1 because we need 1 operator less than there are numbers; -1 + -1 = -2
		possibilities := generatePossibilities(len(eq.Numbers)-1, []string{"*", "+"})
		for _, possibility := range possibilities {
			if evaluate(eq.Numbers, possibility) == eq.Sum {
				result += eq.Sum
				break
			}
		}
	}

	fmt.Println(result)
}

func part2(filename string) {
	result := 0
	equations := readInput(filename)
	for _, eq := range equations {
		// -1 because we need 1 operator less than there are numbers; -1 + -1 = -2
		possibilities := generatePossibilities(len(eq.Numbers)-1, []string{"*", "+", "||"})
		for _, possibility := range possibilities {
			if evaluate(eq.Numbers, possibility) == eq.Sum {
				result += eq.Sum
				break
			}
		}
	}

	fmt.Println(result)
}

func generatePossibilities(n int, operators []string) [][]string {
	possibilities := make([][]string, 0)
	backtrack(n, operators, make([]string, 0, n), &possibilities, 0)

	return possibilities
}

func backtrack(n int, options []string, poss []string, possibilities *[][]string, k int) {
	if len(poss) == n {
		*possibilities = append(*possibilities, poss)
		return
	}

	for i := 0; i < len(options); i++ {
		poss = append(poss, options[i])
		backtrack(n, options, slices.Clone(poss), possibilities, k+1)
		poss = poss[:len(poss)-1]
	}

	return
}

func evaluate(numbers []int, operators []string) int {
	num1 := numbers[0]
	var num2 int
	for idx, operator := range operators {
		num2 = numbers[idx+1]
		switch operator {
		case "*":
			num1 = num1 * num2
		case "+":
			num1 = num1 + num2
		case "||":
			digits := numDigits(num2)
			num1 = int(math.Pow10(digits))*num1 + num2
		}
	}

	return num1
}

type equation struct {
	Sum     int
	Numbers []int
}

func readInput(filename string) []equation {
	return util.GetLinesTransformed[equation](filename, func(s string) (equation, error) {
		parts := strings.Split(s, ":")
		eq := equation{
			Sum: util.MustConvAtoi(parts[0]),
		}

		numStrs := strings.Split(strings.TrimSpace(parts[1]), " ")
		numbers := make([]int, len(numStrs))
		for idx, numStr := range numStrs {
			numbers[idx] = util.MustConvAtoi(numStr)
		}

		eq.Numbers = numbers

		return eq, nil
	})
}

func numDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
