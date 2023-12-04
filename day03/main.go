package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	ROWLENGTH        int
	COLUMNLENGTH     int
	matrix           [][]rune
	digitCoordinates [][][]int
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
	// digitCoordinates: [ [ [], [] ], [ [] ], [ [], [], [] ] ]
	for idxI, i := range matrix {
		var number [][]int

		for idxJ, j := range i {

			if unicode.IsDigit(j) {
				number = append(number, []int{idxI, idxJ})
			} else if len(number) != 0 {
				digitCoordinates = append(digitCoordinates, number)
				number = [][]int{}
			}
		}
	}

	// set length of each line
	ROWLENGTH = len(matrix[0])
	COLUMNLENGTH = len(matrix)

	res1 := part1()
	fmt.Println("part #1: ", res1)
}

func part1() int {

	var partNumbersSum int

	for _, numbers := range digitCoordinates {
		fmt.Println(numbers)
		var nearCoordinates [][]int

		// find all near coordinates of each digit logically. e.g. all x's coordinates will be stored.
		// ......
		// .xxxxx.
		// .x467x.
		// .xxxxx.
		// .......
		for idx, number := range numbers {
			if idx == 0 {
				// --
				// | 
				// --
				if number[1]-1 >= 0 {
					nearCoordinates = append(nearCoordinates, []int{number[0], number[1] - 1}) // left
				}
				if number[0]-1 >= 0 && number[1]-1 >= 0 {
					nearCoordinates = append(nearCoordinates, []int{number[0] - 1, number[1] - 1}) // leftUpper
				}
				if number[0]-1 >= 0 {
					nearCoordinates = append(nearCoordinates, []int{number[0] - 1, number[1]}) // up
				}
				if number[0]+1 < COLUMNLENGTH {
					nearCoordinates = append(nearCoordinates, []int{number[0] + 1, number[1]}) // down
				}
				if number[0]+1 < COLUMNLENGTH && number[1]-1 >= 0 {
					nearCoordinates = append(nearCoordinates, []int{number[0] + 1, number[1] - 1}) // leftDown
				}
			} else if idx == len(numbers) - 1 {
				// --
				//  |
				// --
				if number[0]-1 >= 0 {
					nearCoordinates = append(nearCoordinates, []int{number[0] - 1, number[1]}) // up
				}
				if number[0]-1 >= 0 && number[1]+1 < ROWLENGTH {
					nearCoordinates = append(nearCoordinates, []int{number[0] - 1, number[1] + 1}) // rightUpper
				}
				if number[1]+1 < ROWLENGTH {
					nearCoordinates = append(nearCoordinates, []int{number[0], number[1] + 1}) // right
				}
				if number[0]+1 < COLUMNLENGTH && number[1]+1 < ROWLENGTH {
					nearCoordinates = append(nearCoordinates, []int{number[0] + 1, number[1] + 1}) // rightDown
				}
				if number[0]+1 < COLUMNLENGTH {
					nearCoordinates = append(nearCoordinates, []int{number[0] + 1, number[1]}) // down
				}
			} else {
				// --
				//
				// --
				if number[0]-1 >= 0 {
					nearCoordinates = append(nearCoordinates, []int{number[0] - 1, number[1]}) // up
				}
				if number[0]+1 < COLUMNLENGTH {
					nearCoordinates = append(nearCoordinates, []int{number[0] + 1, number[1]}) // down
				}
			}
		}

		// fmt.Println(digitCoordinates)
		// fmt.Println(numbers, nearCoordinates)

		for _, near := range nearCoordinates {
			x := near[0]
			y := near[1]
			if string(matrix[x][y]) != "." && ! unicode.IsDigit(matrix[x][y]) {
				partNumbersSum += coordinatesToNumber(numbers)
			}
		}
	}

	return partNumbersSum
}

func coordinatesToNumber(coordinates [][]int) int {
	var num int

	for _, coordinate := range coordinates {
		val, err := strconv.Atoi(string(matrix[coordinate[0]][coordinate[1]]))
		if err != nil {
			fmt.Errorf("something unexpected happened, %v", err)
		}
		num = num * 10 + val
	}
	// fmt.Println(num)
	return num
}

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..