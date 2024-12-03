package day3

import (
	"log"
	"strings"
)

func Part01(input []byte) {
	splitContent := strings.Split(string(input), "\n")
	for _, line := range splitContent {
		if len(line) > 1 {
			log.Println(line)
		}
	}
}

// check for symbols with unicode.IsSymbol(r rune)
