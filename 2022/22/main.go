package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type grid [][]byte

type position struct {
	row, col, dir int
}

func (g grid) getPos(p position) byte {
	if p.row < 0 || p.row >= len(g) || p.col < 0 || p.col >= len(g[p.row]) {
		return ' '
	}
	return g[p.row][p.col]
}

type move struct {
	steps int
	turn  rune
}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	readGrid := true
	var g grid
	var m []move
	for scanner.Scan() {
		line := scanner.Text()
		// reading grid stops with empty line
		if line == "" {
			readGrid = false
			continue
		}
		if readGrid {
			g = append(g, []byte(line))
		} else {
			// parse moves
			line += "S" // last move has "S(tay)" turn
			stepsString := ""
			for _, v := range line {
				if v == 'L' || v == 'R' || v == 'S' {
					steps, err := strconv.Atoi(stepsString)
					if err != nil {
						log.Fatal(err)
					}
					m = append(m, move{steps, v})
					stepsString = ""
				} else {
					stepsString += string(v)
				}
			}
		}
	}
	var initialCol int
	for i, v := range g[0] {
		if v != ' ' {
			initialCol = i
		}
	}
	pos := position{0, initialCol, 0}
	for _, mv := range m {
		pos = walkSteps(g, pos, mv.steps)
		pos = doTurn(pos, mv.turn)
	}
	println(1000*(pos.row+1) + 4*(pos.col+1) + pos.dir)
}

func goStep(pos position) position {
	nextPos := pos
	switch pos.dir {
	case 0:
		nextPos.col += 1
	case 1:
		nextPos.row += 1
	case 2:
		nextPos.col -= 1
	case 3:
		nextPos.row -= 1
	}
	return nextPos
}

func walkSteps(g grid, pos position, steps int) position {
	for range steps {
		nextPos := goStep(pos)
		if g.getPos(nextPos) == ' ' {
			nextPos.dir = (nextPos.dir + 2) % 4
			for {
				nextPos = goStep(nextPos)
				if g.getPos(nextPos) == ' ' {
					nextPos.dir = (nextPos.dir + 2) % 4
					nextPos = goStep(nextPos)
					break
				}
			}
		}
		if g.getPos(nextPos) == '#' {
			return pos
		}
		pos = nextPos
	}
	return pos
}

func doTurn(pos position, t rune) position {
	if t == 'L' {
		pos.dir = (pos.dir + 3) % 4
	}
	if t == 'R' {
		pos.dir = (pos.dir + 1) % 4
	}
	return pos
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
