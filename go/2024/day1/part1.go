package day1

import (
	"fmt"
	"log"
	"strings"
)

func Part01(input []byte) {
	// content from file splitted by breaklines
	splitContent := strings.Split(string(input), "\n")

	var total int // value to be returned as result

	for _, line := range splitContent {
		if len(line) == 0 {
			continue
		}

		fmt.Printf("line: %s, length: %d\n", line, len(line))
	}

	log.Println("total:", total)

}
