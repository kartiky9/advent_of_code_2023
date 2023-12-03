package main

import (
	file_utils "aoc2023"
	"fmt"
	"log"
	"strings"
)

// import (
// 	file_utils "aoc2023"
// 	"strings"
// 	"unicode"
// )

type Game struct {
	idx         int
	listOfGames []RGB
}

type RGB struct {
	red   int
	green int
	blue  int
}

func parseRGBString(rgbString string) RGB {
	rgb := RGB{}
	for _, eachSet := range strings.Split(rgbString, ",") {

		var nNum int
		var colour string

		_, err := fmt.Sscanf(strings.TrimSpace(eachSet), "%d %s", &nNum, &colour)
		if err != nil {
			log.Fatal(err)
		}

		if colour == "red" {
			rgb.red = nNum
		} else if colour == "blue" {
			rgb.blue = nNum
		} else {
			rgb.green = nNum
		}
	}
	return rgb
}

func parseGame(line string) Game {
	strSplit := strings.Split(line, ":")
	gameNumString := strSplit[0]
	restOfString := strSplit[1]

	var gameIdx int
	_, err := fmt.Sscanf(gameNumString, "Game %d", &gameIdx)
	if err != nil {
		log.Fatal(err)
	}

	game := Game{idx: gameIdx, listOfGames: []RGB{}}

	for _, round := range strings.Split(restOfString, ";") {
		rgb := parseRGBString(round)
		game.listOfGames = append(game.listOfGames, rgb)
	}
	return game
}

func partOne(lines []string) int {
	var output int

	for _, eachLine := range lines {
		var isNotPossible bool
		game := parseGame(eachLine)
		for _, eachGame := range game.listOfGames {
			if eachGame.red > 12 || eachGame.green > 13 || eachGame.blue > 14 {
				isNotPossible = true
				break
			}
		}
		if !isNotPossible {
			output += game.idx
		}
	}

	return output
}

func partTwo(lines []string) int {
	var output int

	for _, eachLine := range lines {
		game := parseGame(eachLine)
		var minPossible RGB
		for _, eachGame := range game.listOfGames {
			minPossible.red = max(minPossible.red, eachGame.red)
			minPossible.green = max(minPossible.green, eachGame.green)
			minPossible.blue = max(minPossible.blue, eachGame.blue)
		}
		output += minPossible.red * minPossible.green * minPossible.blue
	}

	return output
}

func main() {
	lines := file_utils.ReadFile("./input.txt")

	println("partOne :- ", partOne(lines))
	println("partTwo :- ", partTwo(lines))
}
