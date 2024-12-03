package day2

import (
	"log"
	"strconv"
	"strings"
)

func Part02(input []byte) {
	sumOfPowers := 0

	splitContent := strings.Split(string(input), "\n")
	for _, line := range splitContent {
		if len(line) == 0 { // blank line at the end
			continue
		}

		gameWithMinimumsNeeded := getMinimumCountNeededForEachColor(line)
		sumOfPowers += getPowerOfGame(gameWithMinimumsNeeded)
	}

	log.Println("Sum of powers:", sumOfPowers)
}

func getMinimumCountNeededForEachColor(line string) Set {
	resSet := Set{}

	record := strings.Split(line, ": ")[1]

	sets := strings.Split(record, "; ")
	for _, set := range sets {
		colors := strings.Split(set, ", ")
		for _, cube := range colors {
			cubeValStr := strings.Split(cube, " ")[0]
			cubeVal, _ := strconv.Atoi(cubeValStr)
			switch strings.Split(cube, " ")[1] {
			case "red":
				if cubeVal > resSet.Red {
					resSet.Red = cubeVal
				}
			case "green":
				if cubeVal > resSet.Green {
					resSet.Green = cubeVal
				}
			case "blue":
				if cubeVal > resSet.Blue {
					resSet.Blue = cubeVal
				}
			}
		}
	}

	return resSet
}

func getPowerOfGame(game Set) int {
	return game.Red * game.Green * game.Blue
}
