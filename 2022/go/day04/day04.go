package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput(test bool) []string {
	var data []byte
	if test {
		data = []byte(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)
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

func solve(in []string) (int, int) {
	p1total := 0
	p2total := 0

	for _, pair := range in {
		sections := strings.Split(pair, ",")
		left := strings.Split(sections[0], "-")
		right := strings.Split(sections[1], "-")
		ll, _ := strconv.Atoi(left[0])
		lu, _ := strconv.Atoi(left[1])
		rl, _ := strconv.Atoi(right[0])
		ru, _ := strconv.Atoi(right[1])

		if (ll <= rl && lu >= ru) || (rl <= ll && ru >= lu) {
			p1total++
		}
		if (ll <= rl && lu >= rl) || (rl <= ll && ru >= ll) {
			p2total++
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
