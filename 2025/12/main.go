package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Parse shapes: compute area of each shape
	idx := 0
	var shapeAreas []int
	for idx < len(lines) {
		line := strings.TrimSpace(lines[idx])
		// Shape header like "0:"
		if len(line) > 0 && line[len(line)-1] == ':' {
			idx++
			area := 0
			for idx < len(lines) && strings.TrimSpace(lines[idx]) != "" {
				for _, ch := range strings.TrimSpace(lines[idx]) {
					if ch == '#' {
						area++
					}
				}
				idx++
			}
			shapeAreas = append(shapeAreas, area)
			continue
		}
		if line == "" {
			idx++
			continue
		}
		break // reached tree section
	}

	// Parse trees: "WxH: c0 c1 c2 ..."
	canFit := 0
	for idx < len(lines) {
		line := strings.TrimSpace(lines[idx])
		idx++
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		dims := strings.Split(strings.TrimSpace(parts[0]), "x")
		w, _ := strconv.Atoi(dims[0])
		h, _ := strconv.Atoi(dims[1])
		regionArea := w * h

		counts := strings.Fields(strings.TrimSpace(parts[1]))
		totalPresents := 0
		for i, cs := range counts {
			c, _ := strconv.Atoi(cs)
			totalPresents += c * shapeAreas[i]
		}

		if totalPresents <= regionArea {
			canFit++
		}
	}

	fmt.Println("Part 1:", canFit)
	fmt.Println("Part 2:")
}
