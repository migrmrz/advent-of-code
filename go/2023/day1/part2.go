package day1

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

var days = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

const minDaysStringValue = 3

func Part02(input []byte) {
	// content from file splitted by breaklines
	splitContent := strings.Split(string(input), "\n")

	var total int // value to be returned as result

	for _, line := range splitContent {
		if len(line) == 0 {
			continue
		}
		first, last := findIntOrWordDigitInLine(line)

		total += (*first*10 + *last)
	}

	log.Println("total:", total)
}

func findIntOrWordDigitInLine(line string) (first, last *int) {
	pointer1 := 0
	pointer2 := len(line) - 1

	for {
		if first == nil {
			if unicode.IsDigit(rune(line[pointer1])) {
				val, _ := strconv.Atoi(string(line[pointer1]))

				first = &val

			} else {
				dig := findDigitWordInLine(line[pointer1:])
				if dig != nil {
					first = dig
				} else {
					pointer1 += 1
				}
			}
		} else {
			pointer1 += 1
		}

		if last == nil {
			if unicode.IsDigit(rune(line[pointer2])) {
				val, _ := strconv.Atoi(string(line[pointer2]))

				last = &val
			} else {
				dig := findDigitWordInLine(line[pointer2:])
				if dig != nil {
					last = dig
				} else {
					pointer2 -= 1
				}
			}
		} else {
			pointer2 -= 1
		}

		if first != nil && last != nil {
			return first, last
		}
	}
}

func findDigitWordInLine(line string) (digit *int) {
	var end int // variable to define the end index of which the word is going to be evaluated

	if len(line) < minDaysStringValue { // string to evaluate is not long enough to be a digit word
		return nil
	}

	for k, val := range days {
		if len(line) < len(val) {
			end = len(line)
		} else {
			end = len(val)
		}

		if val == line[:end] {
			return &k
		}
	}

	return nil
}
