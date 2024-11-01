package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type robotSpec struct {
	product string
	costs   map[string]int
}

type blueprint struct {
	id    int
	specs map[string]robotSpec
}

func lastByte(s string) byte {
	return s[len(s)-1]
}

func parseBlueprint(line string) blueprint {
	var id int
	_, err := fmt.Sscanf(line, "Blueprint %d:", &id)
	if err != nil {
		log.Fatal(fmt.Errorf("no blueprint on line %s", err))
	}
	bp := blueprint{id, make(map[string]robotSpec)}

	tokens := strings.Fields(line)
	var robot robotSpec
	for i, s := range tokens {
		if s == "Each" {
			robot = robotSpec{tokens[i+1], make(map[string]int)}
		}
		if s == "costs" || s == "and" {
			cost, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				log.Fatal(fmt.Errorf("could not parse cost %s", err))
			}
			prod := strings.Trim(tokens[i+2], ".")
			robot.costs[prod] = cost
			if lastByte(tokens[i+2]) == '.' {
				bp.specs[robot.product] = robot
			}
		}
	}
	return bp
}

// this assumes the four products ore, clay, obsidion, geode to allow as map key
type state struct {
	minute        int
	productCounts [4]int
	robotCounts   [4]int
}

var intToProduct = [4]string{"ore", "clay", "obsidian", "geode"}
var productToInt = map[string]int{"ore": 0, "clay": 1, "obsidian": 2, "geode": 3}
var bestSolution = 0

func canBuild(specs robotSpec, productCounts [4]int) bool {
	for k, v := range specs.costs {
		if productCounts[productToInt[k]] < v {
			return false
		}
	}
	return true
}

func recMaxGeodes(bp blueprint, maxMinutes int, memo map[state]int, s state) int {
	if s.minute > maxMinutes {
		bestSolution = max(bestSolution, s.productCounts[3])
		return s.productCounts[3]
	}
	remTime := maxMinutes - s.minute
	if s.productCounts[3]+(remTime+1)*s.robotCounts[3]+remTime*(remTime+1)/2 <= bestSolution {
		return 0
	}
	val, ok := memo[s]
	if ok {
		return val
	}
	s.minute += 1
	mx := 0
	for i := 3; i >= 0; i-- {
		v := intToProduct[i]
		if !canBuild(bp.specs[v], s.productCounts) {
			continue
		}
		for j, v := range s.robotCounts {
			s.productCounts[j] += v
		}
		s.robotCounts[i] += 1
		for k, v := range bp.specs[v].costs {
			s.productCounts[productToInt[k]] -= v
		}
		mx = max(mx, recMaxGeodes(bp, maxMinutes, memo, s))
		for k, v := range bp.specs[v].costs {
			s.productCounts[productToInt[k]] += v
		}
		s.robotCounts[i] -= 1
		for j, v := range s.robotCounts {
			s.productCounts[j] -= v
		}

	}
	for i, v := range s.robotCounts {
		s.productCounts[i] += v
	}
	mx = max(mx, recMaxGeodes(bp, maxMinutes, memo, s))
	memo[s] = mx
	return mx
}

func maxGeodes(bp blueprint, maxMinutes int) int {
	bestSolution = 0
	return recMaxGeodes(bp, maxMinutes, make(map[state]int), state{1, [4]int{}, [4]int{1}})
}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		bp := parseBlueprint(line)
		result += bp.id * maxGeodes(bp, 24)
	}
	print(result)
}

func part2(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	result := 1
	for scanner.Scan() {
		line := scanner.Text()
		bp := parseBlueprint(line)
		if bp.id > 3 {
			continue
		}
		result *= maxGeodes(bp, 32)
	}
	print(result)
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
