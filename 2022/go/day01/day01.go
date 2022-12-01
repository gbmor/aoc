package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func parseInput(test bool) [][]int {
	var data []byte
	if test {
		data = []byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
	} else {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		data = b
	}

	packs := strings.Split(string(data), "\n\n")
	out := make([][]int, 0, len(packs))

	for _, pack := range packs {
		packSplit := strings.Split(pack, "\n")
		items := make([]int, 0, len(packSplit))
		for _, item := range packSplit {
			n, _ := strconv.Atoi(item)
			items = append(items, n)
		}
		out = append(out, items)
	}

	return out
}

func solve(in [][]int) (int, int) {
	max := 0
	totals := make([]int, 0, len(in))

	for _, pack := range in {
		total := 0
		for _, item := range pack {
			total += item
		}
		if total > max {
			max = total
		}
		totals = append(totals, total)
	}

	sort.Ints(totals)
	topThree := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]

	return max, topThree
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
