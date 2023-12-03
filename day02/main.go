package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type Pick struct {
	Blue  int
	Green int
	Red   int
}

type Game struct {
	GameID int
	Picks  []Pick
}

type Games struct {
	Games []Game
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	games := parseGames(lines)

	fmt.Println("part #1: ", part1(*games))
	fmt.Println("part #2: ", part2(*games))
}

func part1(games Games) int {

	var possibleGamesSum int

	for _, game := range games.Games {
		flag := true
		for _, pick := range game.Picks {
			if pick.Blue > maxBlue || pick.Green > maxGreen || pick.Red > maxRed {
				flag = false
			}
		}
		if flag {
			possibleGamesSum += game.GameID
		}
	}

	return possibleGamesSum
}

func part2(games Games) int {

	var res int

	for _, game := range games.Games {
		var blue, green, red []int

		for _, pick := range game.Picks {
			blue = append(blue, pick.Blue)
			green = append(green, pick.Green)
			red = append(red, pick.Red)
		}
		product := slices.Max(blue) * slices.Max(green) * slices.Max(red)
		res += product
	}

	return res
}

func parseGames(lines []string) *Games {
	var games Games

	for _, line := range lines {

		var newGame Game

		splitLines := strings.Split(line, ":")

		// set gameID
		gameID, err := strconv.Atoi(strings.Split(splitLines[0], " ")[1])
		if err != nil {
			fmt.Errorf("could not convert to int %v", err)
		}

		newGame.GameID = gameID

		// grab picks(recorded picks by elf)
		picksDraft := strings.Split(splitLines[1], ";")
		for _, pickDraft := range picksDraft {
			var newPick Pick

			// grab color palette at each pick
			palette := strings.Split(pickDraft, ",")

			for _, paletteItem := range palette {

				// get color count of each palette
				splitColor := strings.Split(paletteItem, " ")

				color := splitColor[2]
				switch color {
				case "green":
					newPick.Green, err = strconv.Atoi(splitColor[1])
					if err != nil {
						fmt.Errorf("could not fill pick %v", err)
					}
				case "red":
					newPick.Red, err = strconv.Atoi(splitColor[1])
					if err != nil {
						fmt.Errorf("could not fill pick %v", err)
					}
				case "blue":
					newPick.Blue, err = strconv.Atoi(splitColor[1])
					if err != nil {
						fmt.Errorf("could not fill pick %v", err)
					}
				}
			}
			newGame.Picks = append(newGame.Picks, newPick)
		}
		games.Games = append(games.Games, newGame)
	}

	return &games
}
