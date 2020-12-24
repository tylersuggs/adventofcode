package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Policy   *Policy
	Password string
}

type Policy struct {
	Min       int
	Max       int
	Character string
}

func main() {
	entries := readFile()
	var partOneValid int
	var partTwoValid int

	for _, e := range entries {
		if validatePasswordOne(e) {
			partOneValid++
		}

		if validatePasswordTwo(e) {
			partTwoValid++
		}
	}

	fmt.Println("Part 1:")
	fmt.Printf("There are %v valid passwords!\n\n", partOneValid)

	fmt.Println("Print 2:")
	fmt.Printf("There are %v valid passwords!\n\n", partTwoValid)

}

func validatePasswordOne(e *Entry) bool {
	count := strings.Count(e.Password, e.Policy.Character)
	return (count >= e.Policy.Min && count <= e.Policy.Max)
}

func validatePasswordTwo(e *Entry) bool {
	// return (strings.Index(e.Password, e.Policy.Character) == e.Policy.Min-1) || (strings.LastIndex(e.Password, e.Policy.Character) == e.Policy.Max-1)

	one := string(e.Password[e.Policy.Min-1])
	two := string(e.Password[e.Policy.Max-1])

	return (one == string(e.Policy.Character) && !(two == string(e.Policy.Character)) || !(one == string(e.Policy.Character)) && two == string(e.Policy.Character))
}

func readFile() []*Entry {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var entries []*Entry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entries = append(entries, parseEntry(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries
}

func parseEntry(e string) *Entry {
	split := strings.Split(e, ": ")

	return &Entry{
		Policy:   parsePolicy(split[0]),
		Password: split[1],
	}
}

func parsePolicy(p string) *Policy {
	first := strings.Split(p, " ")
	char := first[1]
	second := strings.Split(first[0], "-")

	min, _ := strconv.Atoi(second[0])
	max, _ := strconv.Atoi(second[1])

	return &Policy{
		Min:       min,
		Max:       max,
		Character: char,
	}
}
