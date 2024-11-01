package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	sum := 0 
	for scanner.Scan() {
		round := scanner.Text()

		points := map[string]int{
			"A X": 1 + 3,
			"A Y": 2 + 6,
			"A Z": 3 + 0,
			"B X": 1 + 0,
			"B Y": 2 + 3,
			"B Z": 3 + 6,
			"C X": 1 + 6,
			"C Y": 2 + 0,
			"C Z": 3 + 3,
		}

		sum += points[round]
	}

	fmt.Println(sum)
}

func part2(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	sum := 0 
	for scanner.Scan() {
		round := scanner.Text()

		points := map[string]int{
			"A X": 3 + 0, // need to play scissors
			"A Y": 1 + 3, // need to play rock
			"A Z": 2 + 6, // need to play paper
			"B X": 1 + 0, // need to play rock
			"B Y": 2 + 3, // need to play paper
			"B Z": 3 + 6, // need to play scissors
			"C X": 2 + 0, // need to play paper
			"C Y": 3 + 3, // need to play rock
			"C Z": 1 + 6, // need to play scissors
		}

		sum += points[round]
	}

	fmt.Println(sum)
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

