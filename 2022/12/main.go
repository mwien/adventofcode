package main

import (
	"bufio"
	"os"
)

const Inf = 1_000_000

type position struct {
	x, y int
}

func parseGrid(filename string) [][]byte {	
	var grid [][]byte
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		grid = append(grid, line)
	}
	return grid
}

func findPosWithValue(grid [][]byte, value byte) position {
	x := 0
	y := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == value {
				x = i 
				y = j
			}
		}
	}
	return position{x, y}
}

func neighbors(grid [][]byte, curPos position) []position {
	var neighbors []position
	if curPos.x + 1 < len(grid) {
		neighbors = append(neighbors, position{curPos.x + 1, curPos.y})
	}
	if curPos.x - 1 >= 0 {
		neighbors = append(neighbors, position{curPos.x - 1, curPos.y})
	}
	if curPos.y + 1 < len(grid[curPos.x]) {
		neighbors = append(neighbors, position{curPos.x, curPos.y + 1})
	}
	if curPos.y - 1 >= 0 {
		neighbors = append(neighbors, position{curPos.x, curPos.y - 1})
	}
	return neighbors
}

func getValue(grid [][]byte, pos position) byte {
	return grid[pos.x][pos.y]
}

func shortestPath(grid [][]byte, startPos position, endPos position) int {
	queue := []position{startPos}
	distances := make([][]int, len(grid))
	for i := range grid {
		distances[i] = make([]int, len(grid[i]))
		for j := range distances[i] {
			distances[i][j] = Inf
		}
	}
	distances[startPos.x][startPos.y] = 0
	for len(queue) > 0 {
		curPos := queue[0]
		if curPos == endPos {
			return distances[curPos.x][curPos.y]
		}
		queue = queue[1:]
		for _, nextPos := range neighbors(grid, curPos) {
			if distances[nextPos.x][nextPos.y] != Inf {
				continue
			}
			curVal := getValue(grid, curPos) 	
			nextVal := getValue(grid, nextPos)
			if curVal == 'S' || nextVal <= curVal + 1 {
				queue = append(queue, nextPos)
				distances[nextPos.x][nextPos.y] = distances[curPos.x][curPos.y] + 1
			}
		}
	}	
	return Inf
}

func part1(filename string) {
	grid := parseGrid(filename)
	startPos := findPosWithValue(grid, 'S')
	grid[startPos.x][startPos.y] = 'a'
	endPos := findPosWithValue(grid, 'E')
	grid[endPos.x][endPos.y] = 'z'
	println(shortestPath(grid, startPos, endPos))
}

func part2(filename string) {
	grid := parseGrid(filename)
	startPos := findPosWithValue(grid, 'S')
	grid[startPos.x][startPos.y] = 'a'
	endPos := findPosWithValue(grid, 'E')
	grid[endPos.x][endPos.y] = 'z'
	minDist := Inf
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'a' {
				dist := shortestPath(grid, position{i, j}, endPos)
				minDist = min(minDist, dist)
			}
		}
	}
	println(minDist)
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
