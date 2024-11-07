package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type position struct {
	row, col int
}

type grid map[position]bool

func neighbors(p position) []position {
	var n []position
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			n = append(n, position{p.row + i, p.col + j})
		}
	}
	return n
}

func hasNoNeighbor(g grid, p position) bool {
	for _, np := range neighbors(p) {
		if g[np] {
			return false
		}
	}
	return true
}

func getDirNeighbors(p position, dir int) []position {
	switch dir {
	case 0:
		return []position{{p.row - 1, p.col - 1}, {p.row - 1, p.col}, {p.row - 1, p.col + 1}}
	case 1:
		return []position{{p.row + 1, p.col - 1}, {p.row + 1, p.col}, {p.row + 1, p.col + 1}}
	case 2:
		return []position{{p.row - 1, p.col - 1}, {p.row, p.col - 1}, {p.row + 1, p.col - 1}}
	case 3:
		return []position{{p.row - 1, p.col + 1}, {p.row, p.col + 1}, {p.row + 1, p.col + 1}}
	}
	log.Fatal("missing case get dir neighbors")
	return []position{}
}

func isEmpty(g grid, p position, dir int) bool {
	for _, np := range getDirNeighbors(p, dir) {
		if g[np] {
			return false
		}
	}
	return true
}

func updatePos(p position, dir int) position {
	switch dir {
	case 0:
		return position{p.row - 1, p.col}
	case 1:
		return position{p.row + 1, p.col}
	case 2:
		return position{p.row, p.col - 1}
	case 3:
		return position{p.row, p.col + 1}
	}
	log.Fatal("case missing update pos")
	return position{}
}

func newGrid(g grid, i int) grid {
	ng := make(grid)
	numMoves := 0
	// check for others
	var proposals []position
	for p := range g {
		if !hasNoNeighbor(g, p) {
			proposals = append(proposals, p)
		} else {
			ng[p] = true
		}
	}
	// make proposals
	proposedPositions := make(map[position]int)
	for _, p := range proposals {
		added := false
		for j := range 4 {
			dir := (i + j) % 4
			if isEmpty(g, p, dir) {
				added = true
				nextPosition := updatePos(p, dir)
				cnt, ok := proposedPositions[nextPosition]
				if ok {
					proposedPositions[nextPosition] = cnt + 1
				} else {
					proposedPositions[nextPosition] = 1
				}
				break
			}
		}
		if !added {
			ng[p] = true
		}
	}
	for _, p := range proposals {
		for j := range 4 {
			dir := (i + j) % 4
			if isEmpty(g, p, dir) {
				nextPosition := updatePos(p, dir)
				if proposedPositions[nextPosition] <= 1 {
					ng[nextPosition] = true
					numMoves += 1
				} else {
					ng[p] = true
				}
				break
			}
		}
	}
	return ng
}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	g := make(grid)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, c := range line {
			if c == '#' {
				g[position{row, col}] = true
			}
		}
		row += 1
	}
	// make moves
	rep := 10
	for i := range rep {
		g = newGrid(g, i)
	}
	var rowMin, rowMax, colMin, colMax int
	for p := range g {
		rowMin = min(rowMin, p.row)
		rowMax = max(rowMax, p.row)
		colMin = min(colMin, p.col)
		colMax = max(colMax, p.col)
	}
	println((rowMax-rowMin+1)*(colMax-colMin+1) - len(g))
}

func part2(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	g := make(grid)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, c := range line {
			if c == '#' {
				g[position{row, col}] = true
			}
		}
		row += 1
	}
	// make moves
	rep := 1000
	for i := range rep {
		ng := newGrid(g, i)
		isSame := true
		for p := range ng {
			if !g[p] {
				isSame = false
			}
		}
		if isSame {
			println(i + 1)
			break
		}
		g = ng
	}
}

func main() {
	if len(os.Args) < 3 {
		err := errors.New("received less than two arguments: expects (part) '1' or '2' as first and (input) 'filename' as second argument")
		log.Fatal(err)
	}

	file, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open input file %s", err))
	}
	defer file.Close()

	switch os.Args[1] {
	case "1":
		part1(file)
	case "2":
		part2(file)
	default:
		err := fmt.Errorf("invalid part specification: only '1' and '2' are accepted, received %s", os.Args[1])
		log.Fatal(err)
		return
	}
}
