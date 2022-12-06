package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func parseInput(test bool) []string {
	var data []byte
	if test {
		data = []byte(``)
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
