package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func updateOutput(output *int, register int, cycle int) {
	if (cycle + 20) % 40 == 0 {
		*output += register * cycle
	}
}

func part1(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	register := 1
	cycle := 0
	output := 0
	for scanner.Scan() {
		lineArray := strings.Split(scanner.Text(), " ")
		if len(lineArray) == 1 {
			cycle += 1	
			updateOutput(&output, register, cycle)
		} else {
			cycle += 1
			updateOutput(&output, register, cycle)
			cycle += 1
			updateOutput(&output, register, cycle)
			addVal, _ := strconv.Atoi(lineArray[1])
			register += addVal	
		}
	}
	println(output)
}

func appendOutput(output *string, register int, cycle int) {
	spriteMiddle := (cycle-1) % 40
	var newChar string
	if register >= spriteMiddle - 1 && register <= spriteMiddle + 1 {
		newChar = "#"
	} else {
		newChar = "."
	}
	*output += newChar
	if cycle % 40 == 0 {
		println(*output)
		*output = ""
	}
}

func part2(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	register := 1
	cycle := 0
	output := ""
	for scanner.Scan() {
		lineArray := strings.Split(scanner.Text(), " ")
		if len(lineArray) == 1 {
			cycle += 1
			appendOutput(&output, register, cycle)
		} else {
			cycle += 1
			appendOutput(&output, register, cycle)
			cycle += 1
			appendOutput(&output, register, cycle)
			addVal, _ := strconv.Atoi(lineArray[1])
			register += addVal	
		}
	}
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

