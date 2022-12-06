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

func solve(in string) (int, int) {
	p1total := 0
	p2total := 0

	for i := 4; i < len(in); i++ {
		set := make(map[byte]struct{})
		for j := i - 4; j < i; j++ {
			set[in[j]] = struct{}{}
		}
		if len(set) == 4 {
			p1total = i
			break
		}
	}

	for i := 14; i < len(in); i++ {
		set := make(map[byte]struct{})
		for j := i - 14; j < i; j++ {
			set[in[j]] = struct{}{}
		}
		if len(set) == 14 {
			p2total = i
			break
		}
	}

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
