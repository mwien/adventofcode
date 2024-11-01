package main

import (
	"bufio"
	"os"
	"strconv"
)

func isList(a string) bool {
	if a[0] == '[' {
		return true
	} else {
		return false
	}
}

func popNextElement(list string) (string, string) {
	openBrackets := 0
	for i := range list {
		if openBrackets == 0 && list[i] == ',' {
			return list[:i], list[i+1:]
		}
		if list[i] == '[' {
			openBrackets += 1
		}
		if list[i] == ']' {
			openBrackets -= 1
		}
	}
	return list, ""
}

func compare(a string, b string) int {
	switch {
	case len(a) == 0 && len(b) == 0:
		return 0
	case len(a) == 0:
		return -1
	case len(b) == 0:
		return 1
	}

	aIsList := isList(a)
	bIsList := isList(b)

	if !aIsList && !bIsList {
		aNum, _ := strconv.Atoi(a)
		bNum, _ := strconv.Atoi(b)
		switch {
		case aNum < bNum:
			return -1
		case aNum == bNum:
			return 0
		case aNum > bNum:
			return 1
		}
	}

	if aIsList {
		a = a[1 : len(a)-1]
	}
	if bIsList {
		b = b[1 : len(b)-1]
	}

	for {
		var aElement, bElement string
		aElement, a = popNextElement(a)
		bElement, b = popNextElement(b)
		comp := compare(aElement, bElement)
		if comp != 0 {
			return comp
		}
		if len(a) == 0 && len(b) == 0 {
			return 0
		}
	}
}

func part1(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	index := 0
	result := 0
	for scanner.Scan() {
		index += 1
		packet1 := scanner.Text()
		scanner.Scan()
		packet2 := scanner.Text()
		scanner.Scan()
		if compare(packet1, packet2) < 1 {
			result += index
		}
	}
	println(result)
}

func part2(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	smallerThanDivider2 := 1
	smallerThanDivider6 := 2 // [[2]] is initially smaller
	for scanner.Scan() {
		packet := scanner.Text()
		if packet == "" {
			continue
		}
		if compare(packet, "[[2]]") == -1 {
			smallerThanDivider2 += 1
		}
		if compare(packet, "[[6]]") == -1 {
			smallerThanDivider6 += 1
		}
	}
	println(smallerThanDivider2 * smallerThanDivider6)
}

func main() {
	if len(os.Args) < 3 {
		println("Error: Received less than two arguments. Expects part (1 or 2) as first and input (sample or main) as second argument.")
		return
	}

	var filename string
	switch os.Args[2] {
	case "sample":
		filename = "sample.in"
	case "main":
		filename = "main.in"
	default:
		println("Error: Did not receive \"sample\" or \"main\" as second argument.")
		return
	}
	switch os.Args[1] {
	case "1":
		part1(filename)
	case "2":
		part2(filename)
	default:
		println("Error: Did not receive \"1\" or \"2\" as first argument.")
		return
	}
}
