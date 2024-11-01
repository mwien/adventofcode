package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type grid [][]byte

type position struct {
	row, col, dir int
}

var dirs = []string{"right", "down", "left", "up"}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	readGrid := true
	var g grid
	var moves string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readGrid = false
		}
		if readGrid {
			g = append(g, []byte(line))
		} else {
			moves = line
		}
	}
	var firstNonEmpty []int
	for i := range g {
		for j := range g[i] {
			if g[i][j] != ' ' {
				firstNonEmpty = append(firstNonEmpty, j)
				break
			}
		}
	}
	pos := position{0, firstNonEmpty[0], 0}
	i := 0
	for i < len(moves) {
		if i == 'R' {
			pos.dir = (pos.dir + 1) % 4
		} else if i == 'L' {
			pos.dir = (pos.dir + 3) % 4
		} else {
			steps := 0
			for j := i; j < len(moves); j++ {
				if j == 'R' || j == 'L' {
					break
				}
				steps *= 10
				steps += int(moves[j] - '0')
			}
			for _ = range steps {
				if pos.col+1 > len(g[pos.row]) {
					pos.col = firstNonEmpty[pos.row]
				} else if pos.col-1 < 0 {
					pos.col = len(g[pos.row]) - 1
				} else {
					pos.col = 
				}
			}
		}
	}
}

func part2(reader io.Reader) {

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
