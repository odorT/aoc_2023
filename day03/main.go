package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var (
	SYMBOLS = "@#$%&*-+=/"
	LENGTH int
	matrix [][]rune
	digitCoordinates [][]int
)

func main() {
	data, err := os.ReadFile("./little.txt")
	if err != nil {
		panic(err)
	}

	// create 2d matrix from input
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		var lineSlice []rune

		for _, char := range line {
			lineSlice = append(lineSlice, char)
		}
		
		matrix = append(matrix, lineSlice)
	}
	
	// store coordinates of each digit
	for idxI, i := range matrix {
		var numbers [][]int
		for idxJ, j := range i {
			if unicode.IsDigit(j) {
				digitCoordinates = append(digitCoordinates, []int{idxI, idxJ})
			}
		}
	}

	// set length of each line
	LENGTH = len(matrix[0])

	res1 := part1()
	fmt.Println("part #1: ", res1)
}

func part1() int {

	// fmt.Println(matrix)
	// for _, i := range matrix {
	// 	for _, j := range i {
	// 		fmt.Println(string(j))
	// 	}
	// }

	for _, coordinate := range digitCoordinates {
		x := coordinate[0]
		y := coordinate[1]
		fmt.Println(string(matrix[x][y]))
	}

	fmt.Println(digitCoordinates)
	return 0
}

// .....
// .345.
// .....

// #1 -> 8
// #2 -> 10
// #2 -> 12

// use X/Y system to find number of dots and if less than full-dots count, add up
