package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func snafuToDecimal(s string) int {
	result := 0
	multiplier := 1
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '2':
			result += multiplier * 2
		case '1':
			result += multiplier * 1
		case '-':
			result -= multiplier * 1
		case '=':
			result -= multiplier * 2
		}
		multiplier *= 5
	}
	return result
}

func decimalToSnafu(x int) string {
	if x == 0 {
		return ""
	}
	s := ""
	switch x % 5 {
	case 0:
		s += "0"
	case 1:
		s += "1"
		x -= 1
	case 2:
		s += "2"
		x -= 2
	case 3:
		s += "="
		x += 2
	case 4:
		s += "-"
		x += 1
	}
	return decimalToSnafu(x/5) + s
}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += snafuToDecimal(line)
	}
	println(sum)
	println(decimalToSnafu(sum))
}

func part2(reader io.Reader) {
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
