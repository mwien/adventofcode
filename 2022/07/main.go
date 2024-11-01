package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	name string
	parent *directory
	children map[string]*directory // maybe dict 
	files map[string]int // map to size 
	size int
}

func newDirectory(name string, parent *directory) *directory {
	d := directory{name: name, parent: parent}
	d.children = make(map[string]*directory)
	d.files = make(map[string]int)
	return &d
}

func computeSizes(currentDir *directory) int {
	size := 0
	for _, childDir := range currentDir.children {
		size += computeSizes(childDir)
	}
	for _, filesize := range currentDir.files {
		size += filesize
	}
	currentDir.size = size
	return size
}

func computeTotal(currentDir *directory) int {
	total := 0
	for _, childDir := range currentDir.children {
		total += computeTotal(childDir)
	}
	if currentDir.size <= 100000 {
		total += currentDir.size
	}
	return total
}

func smallestOver(currentDir *directory, threshold int) int {
	result := math.MaxInt
	for _, childDir := range currentDir.children {
		childResult := smallestOver(childDir, threshold)
		if childResult >= threshold {
			result = min(result, childResult)
		}
	}
	if currentDir.size >= threshold {
		result = min(result, currentDir.size)
	}
	if result == math.MaxInt {
		return 0
	} else {
		return result
	}
}

func getFileTree(filename string) *directory {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	currentDir := newDirectory("/", nil)
	rootDir := currentDir // should not mutate when current dir is mutated

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			if line[2] == 'l' { // ls
				continue
			} 
			// cd 
			targetDir := strings.Split(line, " ")[2]
			//println(target_dir)
			switch targetDir {
			case "..":
				currentDir = currentDir.parent
				if currentDir == nil {
					// there is no parent of root
					currentDir = rootDir
				}
			case "/":
				currentDir = rootDir
			default: 
				nextDir, ok := currentDir.children[targetDir]
				if ok {
					currentDir = nextDir
				}
			}
		} else {
			splitLine := strings.Split(line, " ")
			objectName := splitLine[1]
			if line[0] == 'd' {
				newDir := newDirectory(objectName, currentDir)
				currentDir.children[objectName] = newDir
			} else {
				sz, _ := strconv.Atoi(splitLine[0])
				currentDir.files[objectName] = sz
			}
		}
	}

	computeSizes(rootDir)
	return rootDir
}

func part1(filename string) {
	rootDir := getFileTree(filename)
	println(computeTotal(rootDir))
}

func part2(filename string) {
	rootDir := getFileTree(filename)
	totalUsed := rootDir.size
	needToFree := totalUsed - 40000000
	println(smallestOver(rootDir, needToFree))
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

