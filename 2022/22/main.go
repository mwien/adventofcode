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

func walkSteps1(g grid, pos position, steps int) position {
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

func parseMoves(line string) []move {
	var m []move
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
	return m
}

func parseInput(reader io.Reader) (grid, []move) {
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
			m = parseMoves(line)
		}
	}
	return g, m
}

func initPos(g grid) position {
	for i, v := range g[0] {
		if v != ' ' {
			return position{0, i, 0}
		}
	}
	log.Fatal("no initial pos")
	return position{}
}

func part1(reader io.Reader) {
	g, m := parseInput(reader)
	pos := initPos(g)
	for _, mv := range m {
		pos = walkSteps1(g, pos, mv.steps)
		pos = doTurn(pos, mv.turn)
	}
	println(1000*(pos.row+1) + 4*(pos.col+1) + pos.dir)
}

func walkSteps2(g grid, sideLength int, pos position, steps int) position {
	for range steps {
		nextPos := goStep(pos)
		if g.getPos(nextPos) == ' ' {
			var cubeNum int
			switch {
			case pos.row < sideLength:
				cubeNum = pos.col / sideLength
			case pos.row < 2*sideLength:
				cubeNum = 3
			case pos.row < 3*sideLength:
				cubeNum = 5 - pos.col/sideLength
			default:
				cubeNum = 6
			}
			switch {
			case cubeNum == 1 && pos.dir == 2:
				// cubeNum = 5
				nextPos.row = 3*sideLength - 1 - pos.row
				nextPos.col = 0
				nextPos.dir = 0
			case cubeNum == 1 && pos.dir == 3:
				// cubeNum = 6
				nextPos.row = 2*sideLength + pos.col
				nextPos.col = 0
				nextPos.dir = 0
			case cubeNum == 2 && pos.dir == 0:
				// cubeNum = 4
				nextPos.row = 3*sideLength - 1 - pos.row
				nextPos.col = 2*sideLength - 1
				nextPos.dir = 2
			case cubeNum == 2 && pos.dir == 1:
				// cubeNum = 3
				nextPos.row = pos.col - sideLength
				nextPos.col = 2*sideLength - 1
				nextPos.dir = 2
			case cubeNum == 2 && pos.dir == 3:
				// cubeNum = 6
				nextPos.row = 4*sideLength - 1
				nextPos.col = pos.col - 2*sideLength
				nextPos.dir = 3
			case cubeNum == 3 && pos.dir == 0:
				// cubeNum = 2
				nextPos.row = sideLength - 1
				nextPos.col = pos.row + sideLength
				nextPos.dir = 3
			case cubeNum == 3 && pos.dir == 2:
				// cubeNum = 5
				nextPos.row = 2 * sideLength
				nextPos.col = pos.row - sideLength
				nextPos.dir = 1
			case cubeNum == 4 && pos.dir == 0:
				// cubeNum = 2
				nextPos.row = 3*sideLength - 1 - pos.row
				nextPos.col = 3*sideLength - 1
				nextPos.dir = 2
			case cubeNum == 4 && pos.dir == 1:
				// cubeNum = 6
				nextPos.row = pos.col + 2*sideLength
				nextPos.col = sideLength - 1
				nextPos.dir = 2
			case cubeNum == 5 && pos.dir == 2:
				// cubeNum = 1
				nextPos.row = 3*sideLength - 1 - pos.row
				nextPos.col = sideLength
				nextPos.dir = 0
			case cubeNum == 5 && pos.dir == 3:
				// cubeNum = 3
				nextPos.row = sideLength + pos.col
				nextPos.col = sideLength
				nextPos.dir = 0
			case cubeNum == 6 && pos.dir == 0:
				// cubeNum = 4
				nextPos.row = 3*sideLength - 1
				nextPos.col = pos.row - 2*sideLength
				nextPos.dir = 3
			case cubeNum == 6 && pos.dir == 1:
				// cubeNum = 2
				nextPos.row = 0
				nextPos.col = pos.col + 2*sideLength
				nextPos.dir = 1
			case cubeNum == 6 && pos.dir == 2:
				// cubeNum = 1
				nextPos.row = 0
				nextPos.col = pos.row - 2*sideLength
				nextPos.dir = 1
			default:
				log.Fatal("unknown jump")
			}
		}
		if g.getPos(nextPos) == '#' {
			return pos
		}
		pos = nextPos
	}
	return pos
}

func part2(reader io.Reader) {
	g, m := parseInput(reader)
	pos := initPos(g)
	sideLength := pos.col
	for _, mv := range m {
		pos = walkSteps2(g, sideLength, pos, mv.steps)
		pos = doTurn(pos, mv.turn)
	}
	println(1000*(pos.row+1) + 4*(pos.col+1) + pos.dir)
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
