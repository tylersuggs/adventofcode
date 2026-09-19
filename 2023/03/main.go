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
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:")
	fmt.Println(part1(lines))

	fmt.Println("\nPart 2:")
	fmt.Println(part2(lines))
}

func part1(input []string) int {
	for _, ln := range input {

	}

	return 0
}

func part2(input []string) int {
	return 0
}

func isSymbol(b byte) bool {
	return b != '.' && (b < '0' || b > '9')
}
