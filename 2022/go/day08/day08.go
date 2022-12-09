package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func parseInput(test bool) [][]int {
	var data []byte
	var grid [][]int
	if test {
		data = []byte(`30373
25512
65332
33549
35390`)
		grid = make([][]int, 5)
		for i := range grid {
			grid[i] = make([]int, 5)
		}
	} else {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		data = b
		grid = make([][]int, 99)
		for i := range grid {
			grid[i] = make([]int, 99)
		}
	}

	lines := bytes.Split(data, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			grid[i][j], _ = strconv.Atoi(string(lines[i][j]))
		}
	}

	return grid
}

func checkVisibility(x, y int, in [][]int) (bool, int) {
	point := in[y][x]
	top := true
	bottom := true
	right := true
	left := true
	hopsTop := 0
	hopsBottom := 0
	hopsRight := 0
	hopsLeft := 0
	for i := y - 1; i >= 0; i-- {
		hopsTop++
		if point <= in[i][x] {
			top = false
			break
		}
	}
	for i := y + 1; i < len(in[y]); i++ {
		hopsBottom++
		if point <= in[i][x] {
			bottom = false
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		hopsLeft++
		if point <= in[y][i] {
			left = false
			break
		}
	}
	for i := x + 1; i < len(in); i++ {
		hopsRight++
		if point <= in[y][i] {
			right = false
			break
		}
	}

	visibleFromEdge := top || bottom || left || right
	scenicScore := hopsTop * hopsBottom * hopsLeft * hopsRight

	return visibleFromEdge, scenicScore
}

func solve(in [][]int) (int, int) {
	p1total := 0
	p2total := 0

	p1total += len(in) * 2
	p1total += len(in[0]) * 2
	p1total -= 4

	for i := 1; i < len(in)-1; i++ {
		for j := 1; j < len(in[0])-1; j++ {
			visible, score := checkVisibility(j, i, in)
			if visible {
				p1total++
			}
			if score > p2total {
				p2total = score
			}
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
