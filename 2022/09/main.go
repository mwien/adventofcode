package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func diffAbs(a int, b int) int {
	res := a - b
	if res < 0 {
		res = -res
	}
	return res
}

func part1(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	grid := make(map[point]bool)
	head := point {}
	tail := point {}
	for scanner.Scan() {
		lineArray := strings.Split(scanner.Text(), " ")
		dir := lineArray[0]
		stepSize, _ := strconv.Atoi(lineArray[1])

		for range stepSize {
			// first move head 
			switch dir {
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			case "R":
				head.x += 1
			case "L":
				head.x -= 1
			}
			if diffAbs(tail.x, head.x) == 2 {
				tail.x += (head.x - tail.x) / 2
				if tail.y != head.y {
					tail.y = head.y
				}
			} 
			if diffAbs(tail.y, head.y) == 2 {
				tail.y += (head.y - tail.y) / 2
				if tail.x != head.x {
					tail.x = head.x
				}
			}
			grid[tail] = true
		}
	}
	println(len(grid))
}

func part2(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	grid := make(map[point]bool)
	var knots [10]point;
	for scanner.Scan() {
		lineArray := strings.Split(scanner.Text(), " ")
		dir := lineArray[0]
		stepSize, _ := strconv.Atoi(lineArray[1])

		for range stepSize {
			switch dir {
			case "U":
				knots[0].y += 1
			case "D":
				knots[0].y -= 1
			case "R":
				knots[0].x += 1
			case "L":
				knots[0].x -= 1
			}
			for i := 1; i < 10; i++ {
				if diffAbs(knots[i].x, knots[i-1].x) == 2 {
					knots[i].x += (knots[i-1].x - knots[i].x) / 2
					if knots[i].y != knots[i-1].y {
						knots[i].y += (knots[i-1].y - knots[i].y) / diffAbs(knots[i-1].y, knots[i].y)
					}
				} 
				if diffAbs(knots[i].y, knots[i-1].y) == 2 {
					knots[i].y += (knots[i-1].y - knots[i].y) / 2
					if knots[i].x != knots[i-1].x {
						knots[i].x += (knots[i-1].x - knots[i].x) / diffAbs(knots[i-1].x, knots[i].x)
					}
				}
			}
			grid[knots[9]] = true
		}
	}
	println(len(grid))
}

func main() {
	filename := "sample.in"
	if os.Args[2] == "full" {
		filename = "main.in"
	}
	if os.Args[2] == "extra" {
		filename = "extra.in"
	}

	if os.Args[1] == "1" {
		part1(filename)	
	} else {
		part2(filename)
	}
}

