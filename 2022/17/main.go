package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

type position struct {
	x, y int
}

type rock struct {
	height, width int
	layout        [][]byte
}

func (r rock) listPositions() []position {
	var positions []position
	for i := range r.width {
		for j := range r.height {
			newPos := position{i, j}
			if r.getPos(newPos) == '#' {
				positions = append(positions, newPos)
			}
		}
	}
	return positions
}

func (r rock) getPos(pos position) byte {
	return r.layout[pos.y][pos.x]
}

func readRocks() []rock {
	file, err := os.Open("rocks.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rocks []rock
	var rockLayout [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			slices.Reverse(rockLayout)
			newRock := rock{len(rockLayout), len(rockLayout[0]), rockLayout}
			rocks = append(rocks, newRock)
			rockLayout = [][]byte{}
			continue
		}
		rockLayout = append(rockLayout, line)
	}
	return rocks
}

type grid [][7]byte

func (g grid) getPos(pos position) byte {
	return g[pos.y][pos.x]
}

func (g grid) setPos(pos position, val byte) {
	g[pos.y][pos.x] = val
}

func (g *grid) addRow() {
	var newRow [7]byte
	for j := range newRow {
		newRow[j] = '.'
	}
	*g = append(*g, newRow)
}

func isPlaceable(grid grid, rock rock, tryPos position) bool {
	if tryPos.x < 0 || tryPos.x+rock.width > 7 || tryPos.y < 0 {
		return false
	}
	for _, pos := range rock.listPositions() {
		if grid.getPos(position{pos.x + tryPos.x, pos.y + tryPos.y}) != '.' {
			return false
		}
	}
	return true
}

func applyJet(rockPos position, jet byte) position {
	y := rockPos.y
	x := rockPos.x
	switch jet {
	case '>':
		x += 1
	case '<':
		x -= 1
	}
	return position{x, y}
}

func placeRock(rocks []rock, pattern string, grid *grid, highestRock *int, indexPattern *int, rockNum int) {
	rock := rocks[rockNum%5]
	rockPos := position{2, *highestRock + 4} // lower left corner
	for len(*grid) < rockPos.y+rock.height {
		grid.addRow()
	}
	for {
		jetPos := applyJet(rockPos, pattern[*indexPattern%len(pattern)])
		if isPlaceable(*grid, rock, jetPos) {
			rockPos = jetPos
		}
		*indexPattern += 1

		fallPos := position{rockPos.x, rockPos.y - 1}
		if isPlaceable(*grid, rock, fallPos) {
			rockPos = fallPos
		} else {
			break
		}
	}
	for _, pos := range rock.listPositions() {
		grid.setPos(position{pos.x + rockPos.x, pos.y + rockPos.y}, rock.getPos(pos))
	}
	*highestRock = max(*highestRock, rockPos.y+rock.height-1)
}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	pattern := scanner.Text()
	rocks := readRocks()
	highestRock := -1
	indexPattern := 0
	numPlacements := 2022
	var grid grid
	for i := range numPlacements {
		placeRock(rocks, pattern, &grid, &highestRock, &indexPattern, i)
	}
	println(highestRock + 1)
}

type state struct {
	topRows            string
	rock, indexPattern int
}

type value struct {
	height, timestamp int
}

func part2(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	pattern := scanner.Text()
	rocks := readRocks()
	highestRock := -1
	indexPattern := 0
	numPlacements := 1_000_000_000_000
	var grid grid
	i := 0
	storeHeight := make(map[state]value)
	extraHeight := 0

	for i < numPlacements {
		placeRock(rocks, pattern, &grid, &highestRock, &indexPattern, i)

		if len(grid) < 30 || extraHeight != 0 {
			i += 1
			continue
		}

		var topRowsBuilder strings.Builder
		for i := range 30 {
			topRowsBuilder.WriteString(string(grid[len(grid)-i-1][:]))
		}
		topRows := topRowsBuilder.String()

		val, ok := storeHeight[state{topRows, i % 5, indexPattern % len(pattern)}]
		if ok {
			diffTime := i - val.timestamp
			cycles := (numPlacements - i - 1) / diffTime
			extraHeight = (highestRock - val.height) * cycles
			i += diffTime * cycles
		} else {
			storeHeight[state{topRows, i % 5, indexPattern % len(pattern)}] = value{highestRock, i}
		}
		i += 1
	}
	println(highestRock + extraHeight + 1)
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
