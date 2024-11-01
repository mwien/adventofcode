package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func absDiff(a int, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

type position struct {
	x, y int
}

func manhattanDist(pos1 position, pos2 position) int {
	return absDiff(pos1.x, pos2.x) + absDiff(pos1.y, pos2.y)
}

type measurement struct {
	sensorPos, beaconPos position
}

func parseMeasurement(line string) (measurement, error) {
	var sensorX, sensorY, beaconX, beaconY int
	_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
	if err != nil {
		return measurement{}, err
	}
	sensorPos := position{sensorX, sensorY}
	beaconPos := position{beaconX, beaconY}
	measurement := measurement{sensorPos, beaconPos}
	return measurement, nil
}

func readMeasurements(reader io.Reader) []measurement {
	var measurements []measurement
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		measurement, err := parseMeasurement(line)
		if err != nil {
			log.Printf("skipping line \"%s\" because %s", line, err)
			continue
		}
		measurements = append(measurements, measurement)
	}
	return measurements
}

type interval struct {
	left, right int
}

func getIntervals(measurements []measurement, targetY int) []interval {
	var intervals []interval
	for _, measurement := range measurements {
		dist := manhattanDist(measurement.sensorPos, measurement.beaconPos)
		remainingDist := dist - absDiff(targetY, measurement.sensorPos.y)
		if remainingDist < 0 {
			continue
		}
		left := measurement.sensorPos.x - remainingDist
		right := measurement.sensorPos.x + remainingDist
		intervals = append(intervals, interval{left, right})
	}
	return intervals
}

func part1(reader io.Reader) {
	// set target row according to task (2_000_000 for main, 10 for sample)
	targetY := 2_000_000
	noBeacon := make(map[position]bool)
	knownBeacon := make(map[position]bool)
	measurements := readMeasurements(reader)
	for _, measurement := range measurements {
		if measurement.beaconPos.y == targetY {
			knownBeacon[measurement.beaconPos] = true
		}
	}
	intervals := getIntervals(measurements, targetY)
	for _, interval := range intervals {
		for i := interval.left; i <= interval.right; i++ {
			noBeacon[position{i, targetY}] = true
		}
	}
	println(len(noBeacon) - len(knownBeacon))
}

func part2(reader io.Reader) {
	measurements := readMeasurements(reader)
	numRows := 4_000_000
	for i := 0; i < numRows; i++ {
		intervals := getIntervals(measurements, i)
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].left < intervals[j].left
		})
		maxRight := 0
		for j := 1; j < len(intervals); j++ {
			maxRight = max(maxRight, intervals[j-1].right)
			if maxRight < intervals[j].left-1 {
				println(numRows*(intervals[j].left-1) + i)
				return
			}
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		err := errors.New("received less than two arguments: expects part '1' or '2' as first and input filename as second argument")
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
