package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sortedCaloriesPerElf(filename string) []int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	var caloriesPerElf []int

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// added newline at the end of the input files so that this works for the last elf as well
		if strings.TrimSpace(line) == "" {
			if sum != 0 {
				caloriesPerElf = append(caloriesPerElf, sum)
			}
			sum = 0
			continue
		}

		number, _ := strconv.Atoi(line)
		sum += number
	}

	slices.Sort(caloriesPerElf)

	return caloriesPerElf
}

func part1(filename string) {
	caloriesPerElf := sortedCaloriesPerElf(filename)

	slices.Reverse(caloriesPerElf)

	fmt.Println(caloriesPerElf[0])
}

func part2(filename string) {
	caloriesPerElf := sortedCaloriesPerElf(filename)

	slices.Reverse(caloriesPerElf)

	fmt.Println(caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2])
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
