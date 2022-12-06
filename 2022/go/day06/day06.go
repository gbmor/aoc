package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func parseInput(test bool) string {
	var data []byte
	if test {
		data = []byte(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`)
	} else {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		data = b
	}

	return string(data)
}

func uniqueSequence(length int, in string) int {
	for i := length; i < len(in); i++ {
		set := make(map[byte]struct{})
		for j := i - length; j < i; j++ {
			set[in[j]] = struct{}{}
		}
		if len(set) == length {
			return i
		}
	}

	return -1
}

func solve(in string) (int, int) {
	p1total := uniqueSequence(4, in)
	p2total := uniqueSequence(14, in)

	return p1total, p2total
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
