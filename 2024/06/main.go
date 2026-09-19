package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dirs = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down, left

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Find guard start position and direction
	var sr, sc, sdir int
	for r := range grid {
		for c := range grid[r] {
			switch grid[r][c] {
			case '^':
				sr, sc, sdir = r, c, 0
			case '>':
				sr, sc, sdir = r, c, 1
			case 'v':
				sr, sc, sdir = r, c, 2
			case '<':
				sr, sc, sdir = r, c, 3
			}
		}
	}
	grid[sr][sc] = '.'

	fmt.Println("Part 1:", part1(grid, sr, sc, sdir))
	fmt.Println("Part 2:", part2(grid, sr, sc, sdir))
}

// Part 1: Count distinct positions visited before the guard leaves the map.
func part1(grid [][]byte, sr, sc, sdir int) int {
	visited := patrol(grid, sr, sc, sdir)
	return len(visited)
}

func patrol(grid [][]byte, r, c, dir int) map[[2]int]bool {
	rows, cols := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	for {
		visited[[2]int{r, c}] = true
		nr, nc := r+dirs[dir][0], c+dirs[dir][1]
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			break
		}
		if grid[nr][nc] == '#' {
			dir = (dir + 1) % 4
		} else {
			r, c = nr, nc
		}
	}
	return visited
}

// Part 2: Count positions where adding one obstruction creates a loop.
func part2(grid [][]byte, sr, sc, sdir int) int {
	rows, cols := len(grid), len(grid[0])

	// Only test positions the guard actually visits
	candidates := patrol(grid, sr, sc, sdir)
	delete(candidates, [2]int{sr, sc})

	count := 0
	for pos := range candidates {
		or, oc := pos[0], pos[1]
		grid[or][oc] = '#'
		if loops(grid, rows, cols, sr, sc, sdir) {
			count++
		}
		grid[or][oc] = '.'
	}
	return count
}

func loops(grid [][]byte, rows, cols, r, c, dir int) bool {
	seen := make(map[[3]int]bool)
	for {
		state := [3]int{r, c, dir}
		if seen[state] {
			return true
		}
		seen[state] = true
		nr, nc := r+dirs[dir][0], c+dirs[dir][1]
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
			return false
		}
		if grid[nr][nc] == '#' {
			dir = (dir + 1) % 4
		} else {
			r, c = nr, nc
		}
	}
}
