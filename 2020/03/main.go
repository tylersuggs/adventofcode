package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	//air  = "."
	tree = "#"
)

type Slope struct {
	Cursor map[string]int
	Trees  int
	Right  int
	Down   int
}

func main() {
	slope1 := NewSlope(1, 1)
	slope2 := NewSlope(3, 1)
	slope3 := NewSlope(5, 1)
	slope4 := NewSlope(7, 1)
	slope5 := NewSlope(1, 2)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		ln := scanner.Text()

		slope1.Do(&ln)
		slope2.Do(&ln)
		slope3.Do(&ln)
		slope4.Do(&ln)
		if count%2 == 0 {
			slope5.Do(&ln)
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:")
	fmt.Printf("There are %v trees!\n\n", slope2.Trees)

	fmt.Println("Part 2:")
	fmt.Printf("There are %v trees in Slope 1!\n", slope1.Trees)
	fmt.Printf("There are %v trees in Slope 2!\n", slope2.Trees)
	fmt.Printf("There are %v trees in Slope 3!\n", slope3.Trees)
	fmt.Printf("There are %v trees in Slope 4!\n", slope4.Trees)
	fmt.Printf("There are %v trees in Slope 5!\n", slope5.Trees)

	answer := slope1.Trees * slope2.Trees * slope3.Trees * slope4.Trees * slope5.Trees
	fmt.Printf("Answer: %v\n", answer)

}

func NewSlope(right int, down int) *Slope {
	return &Slope{
		Cursor: map[string]int{
			"x": 0,
			"y": 0,
		},
		Trees: 0,
		Right: right,
		Down:  down,
	}
}

func (s *Slope) Do(ln *string) {
	s.Cursor["y"]++

	// if (s.Down == 2 && s.Cursor["y"]%2 == 0) || s.Down == 1 {
	// 	if string((*ln)[s.Cursor["x"]]) == string(tree) {
	// 		s.Trees++
	// 	}
	// }
	if string((*ln)[s.Cursor["x"]]) == string(tree) {
		s.Trees++
	}
	s.StepCursorRight()
}

func (s *Slope) StepCursorRight() {
	next := s.Cursor["x"] + s.Right
	if next > 30 {
		s.Cursor["x"] = next - 31
	} else {
		s.Cursor["x"] = next
	}
}
