package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var points = map[string]int{
	"A": 1,
	"X": 1,
	"B": 2,
	"Y": 2,
	"C": 3,
	"Z": 3,
}

var beats = map[string]string{
	"X": "C",
	"Y": "A",
	"Z": "B",
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var loses = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var draw = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func outcome(round string) (int, int) {
	fields := strings.Fields(round)
	opp := fields[0]
	me := fields[1]

	p1 := 0
	p2 := 0

	// part 1
	if me == draw[opp] {
		p1 = 3 + points[me]
	} else if me == beats[opp] {
		p1 = 6 + points[me]
	} else {
		p1 = points[me]
	}

	// part 2
	switch me {
	case "X":
		p2 = points[loses[opp]]
	case "Y":
		p2 = 3 + points[opp]
	case "Z":
		p2 = 6 + points[beats[opp]]
	}

	return p1, p2
}

func parseInput(test bool) []string {
	var data []byte
	if test {
		data = []byte(`A Y
B X
C Z`)
	} else {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		data = b
	}

	return strings.Split(string(data), "\n")
}

func solve(rounds []string) (int, int) {
	p1score := 0
	p2score := 0
	for _, round := range rounds {
		p1, p2 := outcome(round)
		p1score += p1
		p2score += p2
	}

	return p1score, p2score
}

func main() {
	start := time.Now()
	data := parseInput(false)
	solveStart := time.Now()
	p1, p2 := solve(data)
	solveTime := time.Since(solveStart)

	fmt.Printf("Part 1:\t%v\n", p1)
	fmt.Printf("Part 2:\t%v\n", p2)
	fmt.Printf("Solve Time:\t%s\n", solveTime)
	fmt.Printf("Run Time:\t%s\n", time.Since(start))
}
