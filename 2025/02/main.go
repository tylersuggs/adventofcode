package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(data))
	ranges := strings.Split(input, ",")

	fmt.Println("Part 1:")
	sum := findInvalidProductIDs(ranges)
	fmt.Printf("Sum of invalid product IDs: %d\n", sum)

	fmt.Println("\n\nPart 2:")
	sum2 := findInvalidProductIDsPart2(ranges)
	fmt.Printf("Sum of invalid product IDs (expanded): %d\n", sum2)
}

func findInvalidProductIDsPart2(ranges []string) int64 {
	var sum int64

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for id := start; id <= end; id++ {
			if isInvalidIDPart2(id) {
				sum += int64(id)
			}
		}
	}

	return sum
}

func isInvalidIDPart2(id int) bool {
	s := strconv.Itoa(id)
	n := len(s)

	// Try each possible pattern length from 1 to n/2
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		// Check if the length is divisible by pattern length
		if n%patternLen == 0 {
			pattern := s[:patternLen]
			// Check if repeating this pattern creates the whole string
			if strings.Repeat(pattern, n/patternLen) == s {
				return true
			}
		}
	}

	return false
}

func findInvalidProductIDs(ranges []string) int64 {
	var sum int64

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for id := start; id <= end; id++ {
			if isInvalidID(id) {
				sum += int64(id)
			}
		}
	}

	return sum
}

func isInvalidID(id int) bool {
	s := strconv.Itoa(id)

	// Must have even length to be split in half
	if len(s)%2 != 0 {
		return false
	}

	// Split in half and check if both halves are equal
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf == secondHalf
}
