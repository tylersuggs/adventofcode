package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	graph := map[string][]string{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		name := strings.TrimSpace(parts[0])
		targets := strings.Fields(strings.TrimSpace(parts[1]))
		graph[name] = targets
	}

	// Count all paths from "you" to "out" using memoized DFS
	memo := map[string]int64{}
	var countPaths func(node string) int64
	countPaths = func(node string) int64 {
		if node == "out" {
			return 1
		}
		if v, ok := memo[node]; ok {
			return v
		}
		total := int64(0)
		for _, next := range graph[node] {
			total += countPaths(next)
		}
		memo[node] = total
		return total
	}

	fmt.Println("Part 1:", countPaths("you"))

	// Part 2: paths from "svr" to "out" visiting both "dac" and "fft"
	type state struct {
		node string
		mask int // bit 0 = visited dac, bit 1 = visited fft
	}
	memo2 := map[state]int64{}
	var countPaths2 func(node string, mask int) int64
	countPaths2 = func(node string, mask int) int64 {
		if node == "dac" {
			mask |= 1
		}
		if node == "fft" {
			mask |= 2
		}
		if node == "out" {
			if mask == 3 {
				return 1
			}
			return 0
		}
		s := state{node, mask}
		if v, ok := memo2[s]; ok {
			return v
		}
		total := int64(0)
		for _, next := range graph[node] {
			total += countPaths2(next, mask)
		}
		memo2[s] = total
		return total
	}

	fmt.Println("Part 2:", countPaths2("svr", 0))
}
