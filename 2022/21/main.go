package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	name, leftChild, rightChild, operation string
	value                                  int
}

func eval(tree map[string]monkey, name string) int {
	m := tree[name]
	if m.value != -1 {
		return m.value
	}

	valueLeftChild := eval(tree, m.leftChild)
	valueRightChild := eval(tree, m.rightChild)

	var value int

	switch m.operation {
	case "+":
		value = valueLeftChild + valueRightChild
	case "-":
		value = valueLeftChild - valueRightChild
	case "*":
		value = valueLeftChild * valueRightChild
	case "/":
		value = valueLeftChild / valueRightChild
	}

	return value
}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	tree := make(map[string]monkey)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)
		if len(splitLine) == 2 {
			name := strings.Trim(splitLine[0], ":")
			value, err := strconv.Atoi(splitLine[1])
			if err != nil {
				log.Printf("skipping line %s because %s", line, err)
			} else {
				tree[name] = monkey{name: name, value: value}
			}
		} else {
			name := strings.Trim(splitLine[0], ":")
			leftChild := splitLine[1]
			operation := splitLine[2]
			rightChild := splitLine[3]
			tree[name] = monkey{name, leftChild, rightChild, operation, -1}
		}
	}
	println(eval(tree, "root"))

}

func part2(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	tree := make(map[string]monkey)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)
		if len(splitLine) == 2 {
			name := strings.Trim(splitLine[0], ":")
			value, err := strconv.Atoi(splitLine[1])
			if err != nil {
				log.Printf("skipping line %s because %s", line, err)
			} else {
				tree[name] = monkey{name: name, value: value}
			}
		} else {
			name := strings.Trim(splitLine[0], ":")
			leftChild := splitLine[1]
			operation := splitLine[2]
			if name == "root" {
				operation = "-"
			}
			rightChild := splitLine[3]
			tree[name] = monkey{name, leftChild, rightChild, operation, -1}
		}
	}
	// this assumes with higher humn value, the result goes down
	// it prints all values that work (can be more than one)
	lb := 0
	ub := math.MaxInt / 2
	mid := 0
	for {
		if lb > ub {
			break
		}
		mid = (lb + ub) / 2
		humn := tree["humn"]
		humn.value = mid
		tree["humn"] = humn
		res := eval(tree, "root")
		if res == 0 {
			println(mid)
			break
		}
		if res < 0 {
			ub = mid - 1
		} else {
			lb = mid + 1
		}
	}
	for mid >= 0 {
		mid -= 1
		humn := tree["humn"]
		humn.value = mid
		tree["humn"] = humn
		if eval(tree, "root") == 0 {
			println(mid)
		} else {
			break
		}
	}
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
