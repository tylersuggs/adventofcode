package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var turns []string

	for scanner.Scan() {
		ln := scanner.Text()

		turns = append(turns, ln)
	}

	fmt.Println("Part 1:")
	countZeroLanding(turns)

	fmt.Println("\n\nPart 2:")
	countZeroCrossings(turns)
}

func countZeroLanding(turns []string) {
	position := 50
	zeroLandings := 0

	for _, turn := range turns {
		direction := turn[0]
		clicks := 0
		for _, ch := range turn[1:] {
			clicks = clicks*10 + int(ch-'0')
		}

		if direction == 'L' {
			position = ((position - clicks) % 100 + 100) % 100
		} else if direction == 'R' {
			position = (position + clicks) % 100
		}

		if position == 0 {
			zeroLandings++
		}
	}

	fmt.Printf("Number of times dial lands on 0: %d", zeroLandings)
}

func countZeroCrossings(turns []string) {
	position := 50
	totalCount := 0

	for _, turn := range turns {
		direction := turn[0]
		clicks := 0
		for _, ch := range turn[1:] {
			clicks = clicks*10 + int(ch-'0')
		}

		var count int
		if direction == 'L' {
			// Count how many times we hit position 0 during the clicks
			if position == 0 {
				// Starting at 0, moving left: hit 0 at clicks 100, 200, etc.
				count = clicks / 100
			} else {
				// Starting at P>0, moving left: hit 0 when we've moved P, P+100, P+200, etc. clicks
				count = (clicks + 100 - position) / 100
			}
			position = ((position - clicks) % 100 + 100) % 100
		} else if direction == 'R' {
			// Count how many times we hit position 0 during the clicks
			// Starting at P, moving right: hit 0 at clicks (100-P), (200-P), etc.
			count = (position + clicks) / 100
			position = (position + clicks) % 100
		}
		totalCount += count
	}

	fmt.Printf("Number of times dial crosses 0: %d", totalCount)
}
