package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseStacks(s string) [][]byte {
	lines := strings.Split(s, "\n")	
	
	stacks := make([][]byte, len(lines[len(lines)-1])/4+1)
	lines = lines[:len(lines)-1]

	slices.Reverse(lines)

	for _, line := range lines {
		cnt := 0
		for {
			if strings.TrimSpace(line) == "" {
				break
			}
			token := line[1]
			if token != ' ' {
				stacks[cnt] = append(stacks[cnt], token)
			}
			if len(line) <= 4 {
				break
			}
			line = line[4:]
			cnt++
		}
	}
	return stacks
}

func parseMove(line string) (int, int, int) {
	f := strings.Fields(line)
	k, _ := strconv.Atoi(f[1])
	from, _ := strconv.Atoi(f[3])
	to, _ := strconv.Atoi(f[5])
	return k, from-1, to-1
}

func cutAt(stack []byte, k int) (start []byte, end[]byte) {
	return stack[:len(stack)-k], stack[len(stack)-k:]
}

func part1(filename string) {
	filebytes, _ := os.ReadFile(filename)
	filestr := string(filebytes)

	s := strings.Split(filestr, "\n\n")

	stacks := parseStacks(s[0])

	for _, line := range strings.Split(s[1], "\n") {
		if strings.TrimSpace(line) == "" {
			break
		}
		k, from, to := parseMove(line)
		stack, mv := cutAt(stacks[from], k)
		slices.Reverse(mv)
		stacks[from] = stack 
		stacks[to] = append(stacks[to], mv...)
	}

	var result strings.Builder
	for _, stack := range stacks {
		result.WriteByte(stack[len(stack)-1])
	}
	fmt.Println(result.String())
}

func part2(filename string) {
	filebytes, _ := os.ReadFile(filename)
	filestr := string(filebytes)

	s := strings.Split(filestr, "\n\n")

	stacks := parseStacks(s[0])

	for _, line := range strings.Split(s[1], "\n") {
		if strings.TrimSpace(line) == "" {
			break
		}
		k, from, to := parseMove(line)
		stack, mv := cutAt(stacks[from], k)
		// slices.Reverse(mv)
		stacks[from] = stack 
		stacks[to] = append(stacks[to], mv...)
	}

	var result strings.Builder
	for _, stack := range stacks {
		result.WriteByte(stack[len(stack)-1])
	}
	fmt.Println(result.String())
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

