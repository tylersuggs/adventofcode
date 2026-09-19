package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct{ r, c int }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	rows, cols := len(grid), len(grid[0])

	// Group antenna positions by frequency
	antennas := map[byte][]point{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ch := grid[r][c]
			if ch != '.' {
				antennas[ch] = append(antennas[ch], point{r, c})
			}
		}
	}

	inBounds := func(p point) bool {
		return p.r >= 0 && p.r < rows && p.c >= 0 && p.c < cols
	}

	// Part 1: for each pair, two antinodes at 2A-B and 2B-A
	anti1 := map[point]bool{}
	for _, pts := range antennas {
		for i := 0; i < len(pts); i++ {
			for j := i + 1; j < len(pts); j++ {
				a, b := pts[i], pts[j]
				p1 := point{2*a.r - b.r, 2*a.c - b.c}
				p2 := point{2*b.r - a.r, 2*b.c - a.c}
				if inBounds(p1) {
					anti1[p1] = true
				}
				if inBounds(p2) {
					anti1[p2] = true
				}
			}
		}
	}

	// Part 2: all collinear grid points along the line through each pair
	anti2 := map[point]bool{}
	for _, pts := range antennas {
		for i := 0; i < len(pts); i++ {
			for j := i + 1; j < len(pts); j++ {
				a, b := pts[i], pts[j]
				dr, dc := b.r-a.r, b.c-a.c
				// Walk both directions from a
				for _, dir := range []int{1, -1} {
					p := a
					for inBounds(p) {
						anti2[p] = true
						p = point{p.r + dir*dr, p.c + dir*dc}
					}
				}
			}
		}
	}

	fmt.Println("Part 1:", len(anti1))
	fmt.Println("Part 2:", len(anti2))
}
