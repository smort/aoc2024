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
	rules, updates := parseInput(filename)

	result := 0
	checkUpdate(updates, rules, func(s []string) {
		numStr := s[len(s)/2]
		result += util.MustConvAtoi(numStr)
	}, func(s []string) {})

	fmt.Println(result)
}

func part2(filename string) {
	rules, updates := parseInput(filename)

	result := 0
	checkUpdate(updates, rules, func(s []string) {
	}, func(s []string) {
		fixed := reorder(s, rules)

		numStr := fixed[len(fixed)/2]
		result += util.MustConvAtoi(numStr)
	})

	fmt.Println(result)
}

func reorder(line []string, rules map[string][]string) []string {
	slices.SortFunc(line, func(a, b string) int {
		if slices.Contains(rules[a+"a"], b) {
			return 1
		}

		if slices.Contains(rules[a+"b"], b) {
			return -1
		}

		return 0
	})

	return line
}

func checkUpdate(updates [][]string, rules map[string][]string, onValid func([]string), onInvalid func([]string)) {
	for _, update := range updates {

		allValid := true
		for idx, num := range update {
			afterRules := rules[num+"a"]
			beforeRules := rules[num+"b"]

			if util.HasIntersection(update[:idx], beforeRules) || util.HasIntersection(update[idx:], afterRules) {
				allValid = false
				break
			}
		}

		if allValid {
			onValid(update)
		} else {
			onInvalid(update)
		}
	}
}

func parseInput(filename string) (map[string][]string, [][]string) {
	ruleMap := make(map[string][]string)
	parseRule := func(s string) {
		numbers := strings.Split(s, "|")
		num1 := numbers[0]
		num2 := numbers[1]
		if _, exists := ruleMap[num1+"b"]; !exists {
			ruleMap[num1+"b"] = make([]string, 0, 1)
		}
		ruleMap[num1+"b"] = append(ruleMap[num1+"b"], num2)

		if _, exists := ruleMap[num2+"a"]; !exists {
			ruleMap[num2+"a"] = make([]string, 0, 1)
		}
		ruleMap[num2+"a"] = append(ruleMap[num2+"a"], num1)
	}
	updates := make([][]string, 0)
	parseUpdate := func(s string) {
		numbers := strings.Split(s, ",")
		updates = append(updates, numbers)
	}

	lines := util.GetLines(filename)

	parser := parseRule
	for _, line := range lines {
		if line == "" {
			parser = parseUpdate
			continue
		}
		parser(line)
	}

	return ruleMap, updates
}
