package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type operation struct {
	operator string
	operand  int
}

func (op operation) applyOperation(x int) int {
	switch op.operator {
	case "pow": // only pow 2 occurs
		return x * x
	case "mult":
		return x * op.operand
	case "add":
		return x + op.operand
	default:
		return x // should not occur
	}
}

type monkey struct {
	items                       []int
	op                          operation
	modulo, nextTrue, nextFalse int
}

func (monkey *monkey) updateWorryLevels() {
	for i, item := range monkey.items {
		monkey.items[i] = monkey.op.applyOperation(item)
		monkey.items[i] /= 3
	}
}

func (monkey monkey) checkItem(item int) bool {
	return item%monkey.modulo == 0
}

func parseInitialMonkeys(filename string) []monkey {
	var monkeys []monkey
	numMonkeys := 0
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "Monkey"):
			monkeys = append(monkeys, monkey{})
			numMonkeys += 1
		case strings.HasPrefix(line, "Starting"):
			fields := strings.Fields(line)
			for i := 2; i < len(fields); i++ {
				item, _ := strconv.Atoi(strings.Trim(fields[i], ","))
				monkeys[numMonkeys-1].items = append(monkeys[numMonkeys-1].items, item)
			}
		case strings.HasPrefix(line, "Operation"):
			fields := strings.Fields(line)
			if fields[len(fields)-1] == "old" {
				monkeys[numMonkeys-1].op.operator = "pow"
				monkeys[numMonkeys-1].op.operand = 2
			} else {
				switch fields[len(fields)-2] {
				case "*":
					monkeys[numMonkeys-1].op.operator = "mult"
					monkeys[numMonkeys-1].op.operand, _ = strconv.Atoi(fields[len(fields)-1])
				case "+":
					monkeys[numMonkeys-1].op.operator = "add"
					monkeys[numMonkeys-1].op.operand, _ = strconv.Atoi(fields[len(fields)-1])
				}
			}
		case strings.HasPrefix(line, "Test"):
			fields := strings.Fields(line)
			modulo, _ := strconv.Atoi(fields[len(fields)-1])
			monkeys[numMonkeys-1].modulo = modulo
		case strings.HasPrefix(line, "If true"):
			fields := strings.Fields(line)
			nextMonkey, _ := strconv.Atoi(fields[len(fields)-1])
			monkeys[numMonkeys-1].nextTrue = nextMonkey
		case strings.HasPrefix(line, "If false"):
			fields := strings.Fields(line)
			nextMonkey, _ := strconv.Atoi(fields[len(fields)-1])
			monkeys[numMonkeys-1].nextFalse = nextMonkey
			// default just skips line
		}
	}
	return monkeys
}

func passItemsFrom(monkeys []monkey, idx int) {
	for _, item := range monkeys[idx].items {
		if monkeys[idx].checkItem(item) {
			monkeys[monkeys[idx].nextTrue].items = append(monkeys[monkeys[idx].nextTrue].items, item)
		} else {
			monkeys[monkeys[idx].nextFalse].items = append(monkeys[monkeys[idx].nextFalse].items, item)
		}
	}
}

func part1(filename string) {
	monkeys := parseInitialMonkeys(filename)
	inspectionCount := make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeys[j].updateWorryLevels()
			inspectionCount[j] += len(monkeys[j].items)
			passItemsFrom(monkeys, j)
			monkeys[j].items = make([]int, 0)
		}
	}
	slices.Sort(inspectionCount)
	n := len(inspectionCount)
	println(inspectionCount[n-1] * inspectionCount[n-2])
}

func (monkey *monkey) updateWorryLevelsPart2(bigmodulo int) {
	for i, item := range monkey.items {
		monkey.items[i] = monkey.op.applyOperation(item)
		monkey.items[i] %= bigmodulo
	}
}

func part2(filename string) {
	monkeys := parseInitialMonkeys(filename)
	bigmodulo := 1
	for i := range monkeys {
		bigmodulo *= monkeys[i].modulo
	}
	inspectionCount := make([]int, len(monkeys))
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeys[j].updateWorryLevelsPart2(bigmodulo)
			inspectionCount[j] += len(monkeys[j].items)
			passItemsFrom(monkeys, j)
			monkeys[j].items = make([]int, 0)
		}
	}
	slices.Sort(inspectionCount)
	n := len(inspectionCount)
	println(inspectionCount[n-1] * inspectionCount[n-2])
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
