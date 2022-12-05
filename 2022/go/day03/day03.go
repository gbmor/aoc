package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func getPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item - 38)
	}
	return int(item - 96)
}

func parseInput(test bool) []string {
	var data []byte
	if test {
		data = []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)
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
	for _, sack := range in {
		comp1 := sack[:len(sack)/2]
		comp2 := sack[len(sack)/2:]
		items := make(map[rune]int)
		for _, k := range comp1 {
			items[k] = 1
		}
		for _, k := range comp2 {
			if items[k] > 0 {
				items[k] = 0
				p1total += getPriority(k)
			}
		}
	}

	for i := 0; i < len(in); i += 3 {
		items1 := make(map[rune]int)
		items2 := make(map[rune]int)
		badge := '0'
		for _, k := range in[i] {
			items1[k]++
		}
		for _, k := range in[i+1] {
			items2[k]++
		}
		for _, k := range in[i+2] {
			if items1[k] > 0 && items2[k] > 0 {
				badge = k
				break
			}
		}

		p2total += getPriority(badge)
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
