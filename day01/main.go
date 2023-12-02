package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var DIGITS = map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	res1 := part1(lines)
	res2 := part2(lines)
	fmt.Println("part #1: ", res1)
	fmt.Println("part #2: ", res2)
}

func part1(lines []string) int {
	var res int

	for _, line := range lines {
		var first, last int

		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == 0 {
					first, _ = strconv.Atoi(string(char))
					last, _ = strconv.Atoi(string(char))
				} else {
					last, _ = strconv.Atoi(string(char))
				}
			}

		}
		res += first*10 + last
	}

	return res
}

func part2(lines []string) int {
	var res int

	for _, line := range lines {
		digits := make(map[int]int)

		// find string digits
		for digitStr, digit := range DIGITS {
			if strings.Contains(line, digitStr) {
				digits[strings.Index(line, digitStr)] = digit
			}
			if strings.LastIndex(line, digitStr) >= 0 {
				digits[strings.LastIndex(line, digitStr)] = digit
			}
		}

		// find numeric digits
		for idx, char := range line {
			if unicode.IsDigit(char) {
				digits[idx], _ = strconv.Atoi(string(char))
			}
		}

		// sort keys
		var keys []int
		for key := range digits {
			keys = append(keys, key)
		}
		slices.Sort(keys)

		res += 10*digits[keys[0]] + digits[keys[len(keys)-1]]
	}

	return res
}
