package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var max uint64 = 0
	var seats []int

	for scanner.Scan() {
		ln := scanner.Text()

		seatId := calculateSeatId(ln)

		if seatId > max {
			max = seatId
		}

		seats = append(seats, int(seatId))
	}

	sort.Ints(seats)

	println("Part 1:")
	fmt.Printf("Max Seat Id: %v\n", max)

	mySeat := findMySeat(seats)
	println("Part 2:")
	if mySeat != 0 {
		fmt.Printf("My Seat: %v\n", mySeat)
	} else {
		fmt.Println("Seat not found!")
	}

}

func calculateSeatId(id string) uint64 {
	rowString := seatToBinary(id[:7], "F", "B")
	colString := seatToBinary(id[7:], "L", "R")

	row, _ := strconv.ParseUint(rowString, 2, 16)
	col, _ := strconv.ParseUint(colString, 2, 16)
	return row*8 + col
}

func seatToBinary(s string, zero string, one string) string {
	s = strings.ReplaceAll(s, zero, "0")
	s = strings.ReplaceAll(s, one, "1")
	return s
}

func findMySeat(seats []int) int {
	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1]-seats[i] == 2 {
			return seats[i] + 1
		}
	}
	return 0
}
