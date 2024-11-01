package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end int
}

func contains(int1 interval, int2 interval) bool {
	return int1.start <= int2.start && int1.end >= int2.end
}

func overlaps(int1 interval, int2 interval) bool {
	return int1.end >= int2.start && int1.start <= int2.end
}

func parseInterval(elf string) interval {
	i := strings.Split(elf, "-")
	start, _ := strconv.Atoi(i[0])
	end, _ := strconv.Atoi(i[1])
	return interval{start, end}
}

func parseLine(line string) (interval, interval) {
	l := strings.Split(line, ",")
	return parseInterval(l[0]), parseInterval(l[1])
}

func part1(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	cnt := 0 
	for scanner.Scan() {
		line := scanner.Text()
		first, second := parseLine(line)
		if contains(first, second) || contains(second, first) {
			cnt++;
		}
	}

	fmt.Println(cnt)
}

func part2(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	cnt := 0 
	for scanner.Scan() {
		line := scanner.Text()
		first, second := parseLine(line)
		if overlaps(first, second) {
			cnt++;
		}
	}

	fmt.Println(cnt)
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

