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
	var grid []string

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	fmt.Println("Part 1:")
	splits := simulateBeams(grid)
	fmt.Printf("Number of splits: %d\n", splits)

	fmt.Println("\n\nPart 2:")
	timelines := simulateBeamsManyWorlds(grid)
	fmt.Printf("Active timelines: %d\n", timelines)
}

func simulateBeams(grid []string) int {
	// Find S position
	startRow, startCol := 0, 0
	for r, line := range grid {
		for c, ch := range line {
			if ch == 'S' {
				startRow = r
				startCol = c
			}
		}
	}

	// Track active beam positions as a set
	beams := map[int]bool{startCol: true}
	totalSplits := 0

	// Process each row below S
	for r := startRow + 1; r < len(grid); r++ {
		newBeams := make(map[int]bool)

		for col := range beams {
			if col >= 0 && col < len(grid[r]) && grid[r][col] == '^' {
				// Beam hits a splitter - split!
				totalSplits++
				// Two new beams on either side
				if col-1 >= 0 {
					newBeams[col-1] = true
				}
				if col+1 < len(grid[r]) {
					newBeams[col+1] = true
				}
			} else {
				// Beam continues downward
				newBeams[col] = true
			}
		}

		beams = newBeams
	}

	return totalSplits
}

func simulateBeamsManyWorlds(grid []string) int64 {
	// Find S position
	startRow, startCol := 0, 0
	for r, line := range grid {
		for c, ch := range line {
			if ch == 'S' {
				startRow = r
				startCol = c
			}
		}
	}

	// Track beam counts at each position
	beams := map[int]int64{startCol: 1}

	for r := startRow + 1; r < len(grid); r++ {
		newBeams := make(map[int]int64)

		for col, count := range beams {
			if col >= 0 && col < len(grid[r]) && grid[r][col] == '^' {
				// Each timeline at this position splits into two
				if col-1 >= 0 {
					newBeams[col-1] += count
				}
				if col+1 < len(grid[r]) {
					newBeams[col+1] += count
				}
			} else {
				// Beams continue downward
				newBeams[col] += count
			}
		}

		beams = newBeams
	}

	// Sum all active timelines
	var total int64
	for _, count := range beams {
		total += count
	}

	return total
}
