package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 1568 low

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	block := []string{}
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()

		if strings.TrimSpace(ln) != "" {
			block = append(block, ln)
			continue
		}

		if len(block) != 0 {
			uniqueBlock := getUnique(strings.Join(block, ""))
			total += len(uniqueBlock)

			block = []string{}
		}
	}

	println("Part 1:")
	println(total)
}

func getUnique(b string) string {
	found := map[rune]bool{}
	var result string

	fmt.Println(b)

	for _, v := range b {
		if !found[v] {
			found[v] = true
			result = result + string(v)
		}
	}

	fmt.Printf("u: %v (%v)\n---\n", result, len(result))

	return result
}
