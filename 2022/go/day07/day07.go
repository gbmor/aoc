package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	lsRgx     = regexp.MustCompile(`^\$ ls$`)
	lsDirRgx  = regexp.MustCompile(`^dir (.*)$`)
	lsFileRgx = regexp.MustCompile(`^(\d+) (.*)$`)
	cdRgx     = regexp.MustCompile(`^\$ cd (.*)$`)
)

type node struct {
	parent   *node
	children map[string]*node
	files    map[string]int
}

func (n *node) directories() []*node {
	dirs := []*node{n}
	for _, v := range n.children {
		subDirs := v.directories()
		dirs = append(dirs, subDirs...)
	}

	return dirs
}

func (n *node) size() int {
	total := 0
	for _, v := range n.children {
		total += v.size()
	}
	for _, v := range n.files {
		total += v
	}

	return total
}

func parseInput(test bool) []string {
	var data []byte
	if test {
		data = []byte(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)
	} else {
		data, _ = os.ReadFile("input.txt")
	}

	return strings.Split(string(data), "\n")
}

func walk(in []string) *node {
	root := &node{
		parent:   nil,
		files:    make(map[string]int),
		children: make(map[string]*node),
	}
	current := root
	for i := 0; i < len(in); i++ {
		matches := lsRgx.FindStringSubmatch(in[i])
		if len(matches) > 0 {
			continue
		}

		matches = cdRgx.FindStringSubmatch(in[i])
		if len(matches) > 0 {
			thisDir := matches[1]
			if thisDir == ".." {
				current = current.parent
				continue
			}
			if thisDir == "/" {
				current = root
				continue
			}

			if _, ok := current.children[thisDir]; !ok {
				current.children[thisDir] = &node{
					files:    make(map[string]int),
					parent:   current,
					children: make(map[string]*node),
				}
			}
			current = current.children[thisDir]
			continue
		}

		matches = lsDirRgx.FindStringSubmatch(in[i])
		if len(matches) > 0 {
			dirName := matches[1]
			if _, ok := current.children[dirName]; !ok {
				current.children[dirName] = &node{
					files:    make(map[string]int),
					parent:   current,
					children: make(map[string]*node),
				}
			}
			continue
		}

		matches = lsFileRgx.FindStringSubmatch(in[i])
		if len(matches) > 0 {
			file := matches[2]
			size, _ := strconv.Atoi(matches[1])
			current.files[file] = size
			continue
		}
	}

	return root
}

func solve(root *node) (int, int) {
	dirs := root.directories()
	sums := make([]int, 0, len(dirs))
	p1total := 0
	for _, d := range dirs {
		s := d.size()
		sums = append(sums, s)
		if s <= 100000 {
			p1total += s
		}
	}

	rsize := root.size()
	total := 70000000
	reqFree := 30000000
	curFree := total - rsize
	req := reqFree - curFree
	sort.Ints(sums)
	p2total := 0
	for _, s := range sums {
		if s >= req {
			p2total = s
			break
		}
	}

	return p1total, p2total
}

func main() {
	start := time.Now()
	data := parseInput(false)
	root := walk(data)
	solveStart := time.Now()
	p1, p2 := solve(root)
	solveTime := time.Since(solveStart)

	fmt.Printf("Part 1:\t%v\n", p1)
	fmt.Printf("Part 2:\t%v\n", p2)
	fmt.Printf("Solve Time:\t%s\n", solveTime)
	fmt.Printf("Run Time:\t%s\n", time.Since(start))
}
