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
	accessible := countAccessibleRolls(grid)
	fmt.Printf("Accessible rolls: %d\n", accessible)

	fmt.Println("\n\nPart 2:")
	removed := removeAllAccessibleRolls(grid)
	fmt.Printf("Total rolls removed: %d\n", removed)
}

func countAccessibleRolls(grid []string) int {
	rows := len(grid)
	count := 0

	for r := 0; r < rows; r++ {
		cols := len(grid[r])
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}

			// Count adjacent rolls in all 8 directions
			adjacentRolls := 0
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					nr, nc := r+dr, c+dc
					if nr >= 0 && nr < rows && nc >= 0 && nc < len(grid[nr]) && grid[nr][nc] == '@' {
						adjacentRolls++
					}
				}
			}

			if adjacentRolls < 4 {
				count++
			}
		}
	}

	return count
}

func removeAllAccessibleRolls(grid []string) int {
	// Make a mutable copy of the grid
	rows := len(grid)
	g := make([][]byte, rows)
	for r := 0; r < rows; r++ {
		g[r] = []byte(grid[r])
	}

	totalRemoved := 0

	for {
		// Find all accessible rolls this round
		var toRemove [][2]int

		for r := 0; r < rows; r++ {
			cols := len(g[r])
			for c := 0; c < cols; c++ {
				if g[r][c] != '@' {
					continue
				}

				adjacentRolls := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						nr, nc := r+dr, c+dc
						if nr >= 0 && nr < rows && nc >= 0 && nc < len(g[nr]) && g[nr][nc] == '@' {
							adjacentRolls++
						}
					}
				}

				if adjacentRolls < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		// Remove all accessible rolls simultaneously
		for _, pos := range toRemove {
			g[pos[0]][pos[1]] = '.'
		}
		totalRemoved += len(toRemove)
	}

	return totalRemoved
}
