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

type grid [][]byte

type state struct {
	pos position
	t   int
}

func parseGrid(reader io.Reader) grid {
	var g grid
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		g = append(g, scanner.Bytes())
	}
	return g
}

func firstDot(s []byte) int {
	for i, c := range s {
		if c == '.' {
			return i
		}
	}
	return -1
}

func gridModulo(i int, add int, mod int) int {
	return 1 + ((((i - 1 + add) % mod) + mod) % mod)
}

func (g grid) getPos(p position) byte {
	if p.row >= len(g) || p.row < 0 || p.col >= len(g[p.row]) || p.col < 0 {
		return '#'
	}
	return g[p.row][p.col]
}

func gridForTime(g grid, t int) grid {
	var ng grid
	for i := range g {
		ng = append(ng, []byte{})
		for j := range g[0] {
			if g[i][j] == '#' {
				ng[i] = append(ng[i], '#')
			} else {
				ng[i] = append(ng[i], '.')
			}
		}
	}
	n := len(g) - 2
	m := len(g[0]) - 2
	for i := range g {
		for j := range g[0] {
			// calculate position of blizzard
			switch g[i][j] {
			case '>':
				ng[i][gridModulo(j, t, m)] = '>'
			case '<':
				ng[i][gridModulo(j, -t, m)] = '<'
			case '^':
				ng[gridModulo(i, -t, n)][j] = '^'
			case 'v':
				ng[gridModulo(i, t, n)][j] = 'v'
			}
		}
	}
	return ng
}

func neighbors(p position) []position {
	var ns []position
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i+j)%2 != 0 {
				ns = append(ns, position{p.row + i, p.col + j})
			}
		}
	}
	ns = append(ns, p)
	return ns
}

func solve(initialGrid grid, startPos position, endPos position, time int) int {
	gridByTime := make(map[int]grid)
	gridByTime[0] = initialGrid

	queue := make([]state, 0)
	visited := make(map[state]bool)

	initialState := state{startPos, time}
	queue = append(queue, initialState)
	visited[initialState] = true

	for len(queue) != 0 {
		x := queue[0]
		queue = queue[1:]
		_, ok := gridByTime[x.t+1]
		if !ok {
			nextGrid := gridForTime(initialGrid, x.t+1)
			gridByTime[x.t+1] = nextGrid
			delete(gridByTime, x.t)

		}
		for _, nextPos := range neighbors(x.pos) {
			nextState := state{nextPos, x.t + 1}
			if gridByTime[nextState.t].getPos(nextPos) == '.' && !visited[nextState] {
				if nextState.pos == endPos {
					return nextState.t
				}
				queue = append(queue, nextState)
				visited[nextState] = true
			}
		}
	}
	return -1
}

func part1(reader io.Reader) {
	grid := parseGrid(reader)
	n := len(grid)
	startPos := position{0, firstDot(grid[0])}
	endPos := position{n - 1, firstDot(grid[n-1])}
	println(solve(grid, startPos, endPos, 0))
}

func part2(reader io.Reader) {
	grid := parseGrid(reader)
	n := len(grid)
	startPos := position{0, firstDot(grid[0])}
	endPos := position{n - 1, firstDot(grid[n-1])}
	firstTime := solve(grid, startPos, endPos, 0)
	println(firstTime)
	secondTime := solve(grid, endPos, startPos, firstTime)
	println(secondTime)
	thirdTime := solve(grid, startPos, endPos, secondTime)
	println(thirdTime)

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
