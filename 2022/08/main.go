package main

import (
	"bufio"
	"os"
)

func readGrid(filename string) [][]int8 {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	grid := make([][]int8, 0)
	for scanner.Scan() {
		line := scanner.Text()
		newRow := make([]int8, 0)
		for _, c := range line {
			newRow = append(newRow, int8(c - '0'))
		}
		grid = append(grid, newRow)
	}
	return grid
}

func updateVisible(mx *int8, i int, j int, isVisible *[][]bool, grid *[][]int8) {
	if (*grid)[i][j] > *mx {
		(*isVisible)[i][j] = true	
	}
	*mx = max(*mx, (*grid)[i][j])
	
}

func part1(filename string) {
	grid := readGrid(filename)
	n := len(grid)
	m := len(grid[0])
	isVisible := make([][]bool, n)
	for i := 0; i < n; i++ {
		isVisible[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		mx := int8(-1)
		for j := 0; j < m; j++ {
			updateVisible(&mx, i, j, &isVisible, &grid)
		}
		mx = int8(-1)
		for j := m-1; j >= 0; j-- {
			updateVisible(&mx, i, j, &isVisible, &grid)
		}
	}
	for j := 0; j < m; j++ {
		mx := int8(-1)
		for i := 0; i < n; i++ {
			updateVisible(&mx, i, j, &isVisible, &grid)
		}
		mx = int8(-1)
		for i := n-1; i >= 0; i-- {
			updateVisible(&mx, i, j, &isVisible, &grid)
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if isVisible[i][j] {
				cnt += 1
			}
		}
	}
	println(cnt)
}

func computeScenicScore(i int, j int, n int, m int, grid *[][]int8) int {
	if i == 0 || i == n-1 || j == 0 || j == m-1 {
		return 0;
	}
	total := 1
	for k := 1; i+k < n; k++ {
		if (*grid)[i+k][j] >= (*grid)[i][j] {
			total *= k
			break;
		}
		if i+k == n-1 {
			total *= k
		}
	}
	for k := 1; i-k >= 0; k++ {
		if (*grid)[i-k][j] >= (*grid)[i][j] {
			total *= k
			break;
		}
		if i-k == 0 {
			total *= k
		}
	}
	for k := 1; j+k < m; k++ {
		if (*grid)[i][j+k] >= (*grid)[i][j] {
			total *= k
			break;
		}
		if j+k == m-1 {
			total *= k
		}
	}
	for k := 1; j-k >= 0; k++ {
		if (*grid)[i][j-k] >= (*grid)[i][j] {
			total *= k
			break;
		}
		if j-k == 0 {
			total *= k
		}
	}
	return total
}

func part2(filename string) {
	grid := readGrid(filename)
	n := len(grid)
	m := len(grid[0])
	mx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mx = max(mx, computeScenicScore(i, j, n, m, &grid))
		}
	}
	println(mx)
}

func main() {
	filename := "sample.in"
	if os.Args[2] == "full" {
		filename = "main.in"
	}

	if os.Args[1] == "1" {
		part1(filename)	
	} else {
		part2(filename)
	}
}

