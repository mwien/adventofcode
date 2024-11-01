package main

import (
	"bufio"
	"fmt"
	"os"
)

func putItem(comp *[52]int, char rune) {
	if char - 'a' < 26 && char - 'a' >= 0 {
		comp[char - 'a']++
	} else {
		comp[char - 'A' + 26]++
	}
}

func computeItems(rucksack string) [52]int {
	var items [52]int

	for _, char := range rucksack {
		putItem(&items, char)
	}

	return items
}

func itemPriority(rucksack string) int {
	n := len(rucksack)
	c := n / 2

	comp1 := computeItems(rucksack[0:c])
	comp2 := computeItems(rucksack[c:n])

	for i := 0; i < 52; i++ {
		if comp1[i] > 0 && comp2[i] > 0 {
			return i + 1
		}
	}

	return -1
}

func part1(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	sum := 0 
	for scanner.Scan() {
		line := scanner.Text()
		sum += itemPriority(line)
	}

	fmt.Println(sum)
}

func part2(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		itemsFirst := computeItems(first)
		itemsSecond := computeItems(second)
		itemsThird := computeItems(third)

		for i := 0; i < 52; i++ {
			if itemsFirst[i] > 0 && itemsSecond[i] > 0 && itemsThird[i] > 0 {
				sum += i + 1
			}
		}
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

