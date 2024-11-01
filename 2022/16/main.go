package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const INF = 1_000_000_000

type node struct {
	id        string
	flowRate  int
	neighbors []string
}

func parseNode(line string) (node, error) {
	var id string
	var flowRate int
	// read id and flowrate
	_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &id, &flowRate)
	if err != nil {
		return node{}, err
	}
	// read neighboring valves
	neighbors := strings.Fields(line)[9:]
	for i, v := range neighbors {
		neighbors[i] = strings.Trim(v, ",")
	}
	// create node
	node := node{id, flowRate, neighbors}
	return node, nil

}

func readNodes(reader io.Reader) map[string]node {
	nodes := make(map[string]node)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		newNode, err := parseNode(line)
		if err != nil {
			log.Printf("skipping line \"%s\" because %s", line, err)
			continue
		}
		nodes[newNode.id] = newNode
	}
	return nodes
}

func getDistsFrom(id string, nodes map[string]node) map[string]int {
	var queue []string
	visited := make(map[string]bool)
	dists := make(map[string]int)
	queue = append(queue, id)
	visited[id] = true
	dists[id] = 0
	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range nodes[u].neighbors {
			if !visited[v] {
				queue = append(queue, v)
				visited[v] = true
				dists[v] = dists[u] + 1
			}
		}
	}
	return dists
}

func computePairwiseDistances(nonZeroNodes []string, nodes map[string]node) [][]int {
	n := len(nonZeroNodes)
	dists := make([][]int, n)
	for i := range dists {
		dists[i] = make([]int, n)
		for j := range n {
			dists[i][j] = INF
		}
	}
	for i, source := range nonZeroNodes {
		sourceDists := getDistsFrom(source, nodes)
		for j, target := range nonZeroNodes {
			dist, ok := sourceDists[target]
			if ok {
				dists[i][j] = dist
			}
		}
	}
	return dists
}

func maxPressureRelease(flowrates []int, dists [][]int, memo map[string]int, curIndex int, opened []bool, minute int) int {
	if minute >= 30 {
		return 0
	}
	idString := fmt.Sprintf("%d,", curIndex)
	for i, v := range opened {
		if v {
			idString += fmt.Sprintf("%d,", i)
		}
	}
	idString += fmt.Sprintf("%d", minute)
	res, ok := memo[idString]
	if ok {
		return res
	}

	bestSolution := 0

	if !opened[curIndex] && flowrates[curIndex] > 0 {
		opened[curIndex] = true
		bestSolution += (30 - minute - 1) * flowrates[curIndex]
		bestSolution += maxPressureRelease(flowrates, dists, memo, curIndex, opened, minute+1)
		opened[curIndex] = false
	} else {
		for i, v := range opened {
			if !v && i != curIndex {
				newSolution := maxPressureRelease(flowrates, dists, memo, i, opened, minute+dists[curIndex][i])
				bestSolution = max(bestSolution, newSolution)
			}
		}
	}

	memo[idString] = bestSolution
	return bestSolution
}

func part1(reader io.Reader) {
	nodes := readNodes(reader)
	// find relevant nodes (flow rate != 0) and their distances
	var nonZeroNodes []string
	for id, node := range nodes {
		if node.flowRate > 0 || node.id == "AA" {
			nonZeroNodes = append(nonZeroNodes, id)
		}
	}
	var flowrates []int
	for _, id := range nonZeroNodes {
		flowrates = append(flowrates, nodes[id].flowRate)
	}
	dists := computePairwiseDistances(nonZeroNodes, nodes)
	startIndex := 0
	for i, v := range nonZeroNodes {
		if v == "AA" {
			startIndex = i
		}
	}
	memo := make(map[string]int)
	opened := make([]bool, len(flowrates))
	result := maxPressureRelease(flowrates, dists, memo, startIndex, opened, 0)
	println(result)
}

func maxPressureRelease2(flowrates []int, dists [][]int, memo map[string]int, curIndices []int, opened []bool, minutes []int) int {
	turn := 0
	if minutes[1] < minutes[0] {
		turn = 1
	}
	if minutes[turn] >= 26 {
		return 0
	}

	var idStringBuilder strings.Builder
	idStringBuilder.WriteString(fmt.Sprintf("%d,%d,", curIndices[0], curIndices[1]))
	for i, v := range opened {
		if v {
			idStringBuilder.WriteString(fmt.Sprintf("%d,", i))
		}
	}
	idStringBuilder.WriteString(fmt.Sprintf("%d,%d", minutes[0], minutes[1]))
	idString := idStringBuilder.String()
	res, ok := memo[idString]
	if ok {
		return res
	}

	bestSolution := 0
	curIndex := curIndices[turn]
	minute := minutes[turn]

	if !opened[curIndex] && flowrates[curIndex] > 0 {
		opened[curIndex] = true
		minutes[turn] = minute + 1
		bestSolution += (26 - minute - 1) * flowrates[curIndex]
		bestSolution += maxPressureRelease2(flowrates, dists, memo, curIndices, opened, minutes)
		minutes[turn] = minute
		opened[curIndex] = false
	} else {
		moved := false
		for i, v := range opened {
			if !v && i != curIndex && i != curIndices[1-turn] && minute+dists[curIndex][i] < 26 {
				moved = true
				curIndices[turn] = i
				minutes[turn] = minute + dists[curIndex][i]
				newSolution := maxPressureRelease2(flowrates, dists, memo, curIndices, opened, minutes)
				minutes[turn] = minute
				curIndices[turn] = curIndex
				bestSolution = max(bestSolution, newSolution)
			}
		}
		if !moved {
			minutes[turn] = 30
			newSolution := maxPressureRelease2(flowrates, dists, memo, curIndices, opened, minutes)
			minutes[turn] = minute
			bestSolution = max(bestSolution, newSolution)
		}
	}

	memo[idString] = bestSolution
	return bestSolution
}

func part2(reader io.Reader) {
	nodes := readNodes(reader)
	// find relevant nodes (flow rate != 0) and their distances
	var nonZeroNodes []string
	for id, node := range nodes {
		if node.flowRate > 0 || node.id == "AA" {
			nonZeroNodes = append(nonZeroNodes, id)
		}
	}
	var flowrates []int
	for _, id := range nonZeroNodes {
		flowrates = append(flowrates, nodes[id].flowRate)
	}
	dists := computePairwiseDistances(nonZeroNodes, nodes)
	startIndex := 0
	for i, v := range nonZeroNodes {
		if v == "AA" {
			startIndex = i
		}
	}
	memo := make(map[string]int)
	opened := make([]bool, len(flowrates))
	result := maxPressureRelease2(flowrates, dists, memo, []int{startIndex, startIndex}, opened, []int{0, 0})
	println(result)
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
