package day1

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func Part01(input []byte) {
	// content from file splitted by breaklines
	splitContent := strings.Split(string(input), "\n")

	var total int // value to be returned as result

	for _, line := range splitContent {
		if len(line) == 0 {
			continue
		}
		first, last := findIntDigitsInLine(line)

		total += (*first*10 + *last)
		// content from file splitted by breaklines
	}

	log.Println("total:", total)

}

func findIntDigitsInLine(line string) (first, last *int) {
	pointer1 := 0
	pointer2 := len(line) - 1

	for {
		if first == nil && unicode.IsDigit(rune(line[pointer1])) {
			val, _ := strconv.Atoi(string(line[pointer1]))

			first = &val

		} else {
			pointer1 += 1
		}

		if last == nil && unicode.IsDigit(rune(line[pointer2])) {
			val, _ := strconv.Atoi(string(line[pointer2]))

			last = &val
		} else {
			pointer2 -= 1
		}

		if first != nil && last != nil {
			return first, last
		}
	}
}
