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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(grid))
}

// Part 1: Count all occurrences of "XMAS" in any direction.
func part1(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	word := "XMAS"
	dirs := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, d := range dirs {
				if matches(grid, r, c, d[0], d[1], word) {
					count++
				}
			}
		}
	}
	return count
}

func matches(grid []string, r, c, dr, dc int, word string) bool {
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < len(word); i++ {
		nr, nc := r+dr*i, c+dc*i
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			return false
		}
		if grid[nr][nc] != word[i] {
			return false
		}
	}
	return true
}

// Part 2: Count X-MAS patterns — two "MAS" crossing diagonally at their center 'A'.
func part2(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] != 'A' {
				continue
			}
			// Check both diagonals for MAS or SAM
			d1 := string([]byte{grid[r-1][c-1], grid[r][c], grid[r+1][c+1]})
			d2 := string([]byte{grid[r-1][c+1], grid[r][c], grid[r+1][c-1]})
			if (d1 == "MAS" || d1 == "SAM") && (d2 == "MAS" || d2 == "SAM") {
				count++
			}
		}
	}
	return count
}
