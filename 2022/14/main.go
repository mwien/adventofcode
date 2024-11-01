package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func readPaths(filename string) [][]position {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var paths [][]position
	for scanner.Scan() {
		line := scanner.Text()
		positions := strings.Split(line, " -> ")
		var path []position
		for _, posString := range positions {
			posStringFields := strings.Split(posString, ",")
			x, _ := strconv.Atoi(posStringFields[0])
			y, _ := strconv.Atoi(posStringFields[1])
			path = append(path, position{x, y})
		}
		paths = append(paths, path)
	}
	return paths
}

func swapIfNotOrdered(a int, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func initGrid(paths [][]position) [][]byte {
	n := 0
	m := 0
	for _, path := range paths {
		for _, pos := range path {
			n = max(n, pos.x+1)
			m = max(m, pos.y+1)
		}
	}

	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, m)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			grid[i][j] = '.'
		}
	}

	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			if path[i-1].x != path[i].x {
				from, to := swapIfNotOrdered(path[i-1].x, path[i].x)
				for j := from; j <= to; j++ {
					grid[j][path[i].y] = '#'
				}
			}
			if path[i-1].y != path[i].y {
				from, to := swapIfNotOrdered(path[i-1].y, path[i].y)
				for j := from; j <= to; j++ {
					grid[path[i].x][j] = '#'
				}
			}
		}
	}

	return grid
}

func updatePos(grid [][]byte, pos position) position {
	switch {
	case pos.y+1 == len(grid[pos.x]) || grid[pos.x][pos.y+1] == '.': 
		return position{pos.x, pos.y+1} 
	case pos.x == 0 || grid[pos.x-1][pos.y+1] == '.':
		return position{pos.x-1, pos.y+1}
	case pos.x+1 == len(grid) || grid[pos.x+1][pos.y+1] == '.':
		return position{pos.x+1, pos.y+1}
	default:
		return pos
	}
}

func part1(filename string) {
	paths := readPaths(filename)
	grid := initGrid(paths)
	cnt := 0
	for {
		currentPos := position{500, 0}
		done := false
		for {
			newPos := updatePos(grid, currentPos)
			if newPos.x < 0 || newPos.x >= len(grid) || newPos.y >= len(grid[newPos.x]) {
				done = true
				break
			}
			if newPos == currentPos {
				grid[newPos.x][newPos.y] = 'o'
				cnt += 1
				break
			}
			currentPos = newPos
		}
		if done {
			break
		}
	}
	println(cnt)
}

func part2(filename string) {
	paths := readPaths(filename)
	maxY := 0
	for _, path := range paths {
		for _, pos := range path {
			maxY = max(maxY, pos.y)
		}
	}
	extraPath := []position{{500-maxY-2, maxY+2}, {500+maxY+2, maxY+2}}
	paths = append(paths, extraPath)
	grid := initGrid(paths)
	cnt := 0
	for {
		startPos := position{500, 0}
		currentPos := startPos
		done := false
		for {
			newPos := updatePos(grid, currentPos)
			if newPos == currentPos {
				grid[newPos.x][newPos.y] = 'o'
				cnt += 1
				if newPos == startPos {
					done = true
				}
				break
			}
			currentPos = newPos
		}
		if done {
			break
		}
	}
	println(cnt)
}

func main() {
	if len(os.Args) < 3 {
		println("Error: Received less than two arguments. Expects part (1 or 2) as first and input (sample or main) as second argument.")
		return
	}

	var filename string
	switch os.Args[2] {
	case "sample":
		filename = "sample.in"
	case "main":
		filename = "main.in"
	default:
		println("Error: Did not receive \"sample\" or \"main\" as second argument.")
		return
	}
	switch os.Args[1] {
	case "1":
		part1(filename)
	case "2":
		part2(filename)
	default:
		println("Error: Did not receive \"1\" or \"2\" as first argument.")
		return
	}
}
