package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var instructionRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parseInput(test bool) ([]string, [][]string) {
	var data string
	if test {
		data = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	} else {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		data = string(b)
	}

	dataSplit := strings.Split(data, "\n\n")
	stacksRaw := dataSplit[0]
	stacksStrings := strings.Split(stacksRaw, "\n")
	instructions := dataSplit[1]
	numColsRaw := stacksStrings[len(stacksStrings)-1]
	numCols, _ := strconv.Atoi(string(numColsRaw[len(numColsRaw)-1]))
	stacks := make([][]string, numCols)
	for i := len(stacksStrings) - 1; i >= 0; i-- {
		row := stacksStrings[i]
		count := 0
		for j := 1; j < len(row); j += 4 {
			val := string(row[j])
			if val == " " {
				count++
				continue
			}

			if len(stacks[count]) == 0 {
				stacks[count] = make([]string, 0)
			}
			stacks[count] = append(stacks[count], val)
			count++
		}
	}

	return strings.Split(instructions, "\n"), stacks
}

func part2(in []string, stacks [][]string) string {
	p2 := strings.Builder{}
	for _, line := range in {
		s := instructionRegex.FindStringSubmatch(line)
		count, _ := strconv.Atoi(s[1])
		src, _ := strconv.Atoi(s[2])
		dest, _ := strconv.Atoi(s[3])

		src -= 1
		dest -= 1

		srclen := len(stacks[src])

		stacks[dest] = append(stacks[dest], stacks[src][srclen-count:srclen]...)
		newsrc := make([]string, len(stacks[src])-count)
		copy(newsrc, stacks[src][:srclen-count])
		stacks[src] = newsrc
	}

	for _, stack := range stacks {
		p2.WriteString(stack[len(stack)-1])
	}

	return p2.String()
}

func part1(in []string, stacks [][]string) string {
	p1 := strings.Builder{}
	for _, line := range in {
		s := instructionRegex.FindStringSubmatch(line)
		if len(s) < 4 {
			continue
		}
		count, _ := strconv.Atoi(s[1])
		src, _ := strconv.Atoi(s[2])
		dest, _ := strconv.Atoi(s[3])

		src -= 1
		dest -= 1

		srclen := len(stacks[src])

		for j := 1; j <= count; j++ {
			got := stacks[src][srclen-j]
			stacks[dest] = append(stacks[dest], got)
		}

		newsrc := make([]string, len(stacks[src])-count)
		copy(newsrc, stacks[src][:srclen-count])
		stacks[src] = newsrc
	}

	for _, stack := range stacks {
		p1.WriteString(stack[len(stack)-1])
	}

	return p1.String()
}

func main() {
	start := time.Now()
	data, stacks := parseInput(false)
	p1Start := time.Now()
	p1 := part1(data, stacks)
	p1Done := time.Since(p1Start)

	data, stacks = parseInput(false)
	p2Start := time.Now()
	p2 := part2(data, stacks)
	p2Done := time.Since(p2Start)

	fmt.Printf("Part 1:\t%v\t(%s)\n", p1, p1Done)
	fmt.Printf("Part 2:\t%v\t(%s)\n", p2, p2Done)
	fmt.Printf("Run Time:\t\t(%s)\n", time.Since(start))
}
