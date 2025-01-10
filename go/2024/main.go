package main

import (
	"aoc/day1"
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	// variables declaration
	var (
		day      string
		part     string
		filePath string
	)

	flag.StringVar(&day, "day", "01", "day")
	flag.StringVar(&part, "part", "01", "part (01 or 02)")
	flag.StringVar(&filePath, "filePath", "", "file path from the example or test data")

	flag.Parse()

	if day == "" || filePath == "" || part == "" {
		log.Panic("either 'day', 'part' or 'filePath' params were not provided")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic("not able to read file:", err.Error())
	}

	runDay(day, part, content)
}

func runDay(day, part string, content []byte) {
	startTime := time.Now()

	log.Println("**************************************")

	switch day {
	case "01":
		switch part {
		case "01":
			day1.Part01(content)
			// case "02":
			// 	day1.Part02(content)
		}
		// case "02":
		// 	switch part {
		// 	case "01":
		// 		day2.Part01(content)
		// 	case "02":
		// 		day2.Part02(content)
		// 	}
		// case "03":
		// 	switch part {
		// 	case "01":
		// 		day3.Part01(content)
		// 		// case "02":
		// 		//		day3.Part02(content)
		// 	}
	}

	log.Println("elapsed time:", time.Since(startTime))

	log.Println("**************************************")
}
