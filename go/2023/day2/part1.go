package day2

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type Set struct {
	Red   int
	Green int
	Blue  int
}

func Part01(input []byte) {
	gamesSum := 0

	splitContent := strings.Split(string(input), "\n")
	for _, line := range splitContent {
		if len(line) > 1 {
			err := validateGameSets(line)
			if err == nil {
				gamesSum += getGameID(line)
			}
		}
	}

	log.Println("Result:", gamesSum)
}

func getGameID(line string) int {
	title := strings.Split(line, ":")[0]
	idStr := strings.Split(title, " ")[1]

	id, _ := strconv.Atoi(idStr)

	return id
}

func validateGameSets(line string) error {
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
				if cubeVal > 12 {
					return errors.New("red cube exceeded max value")
				}
				resSet.Red = cubeVal
			case "green":
				if cubeVal > 13 {
					return errors.New("green cube exceeded max value")
				}
				resSet.Green = cubeVal
			case "blue":
				if cubeVal > 14 {
					return errors.New("blue cube exceeded max value")
				}
				resSet.Blue = cubeVal
			}
		}

		// gameSets = append(gameSets, resSet)
		resSet = Set{}
	}

	return nil
}
