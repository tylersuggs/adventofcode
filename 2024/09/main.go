package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)
	_ = input

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input string) int {
	// TODO
	fid := 0
	var blockById []int
	checksum := 0
	for i := 0; i < len(input); i++ {
		block := int(input[i] - '0')
		for j := 1; j <= block; j++ {
			if i%2 == 0 {
				blockById = append(blockById, fid)
			} else {
				blockById = append(blockById, -1)
			}
		}
		if i%2 == 0 {
			fid++
		}
	}

	for l := 0; l < len(blockById); l++ {
		if blockById[l] != -1 {
			continue
		}

		for r := len(blockById) - 1; r >= l; r-- {
			if blockById[r] != -1 {
				blockById[l] = blockById[r]
				blockById[r] = -1

				break
			}
		}
	}

	for i, v := range blockById {
		if v != -1 {
			checksum += i * v
		}
	}

	return checksum
}

func part2(input string) int {
	// TODO
	fid := 0
	var blockById []int
	checksum := 0
	for i := 0; i < len(input); i++ {
		block := int(input[i] - '0')
		for j := 1; j <= block; j++ {
			if i%2 == 0 {
				blockById = append(blockById, fid)
			} else {
				blockById = append(blockById, -1)
			}
		}
		if i%2 == 0 {
			fid++
		}
	}

	type group struct {
		id     int
		start  int
		length int
	}
	var groups []group
	var groupStart = 0
	var cgc = -2

	for k, c := range blockById {
		if c != cgc {
			// the OLD group ran from groupStart to k-1
			groups = append(groups, group{cgc, groupStart, k - groupStart})
			cgc = c
			groupStart = k // the NEW group starts here
		}
	}
	groups = append(groups, group{cgc, groupStart, len(blockById) - groupStart})

	// Here

	return 0
}
