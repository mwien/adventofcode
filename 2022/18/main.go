package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type cube struct {
	x, y, z int
}

func parseCube(s string) cube {
	var x, y, z int
	_, err := fmt.Sscanf(s, "%d,%d,%d", &x, &y, &z)
	if err != nil {
		log.Printf("skipping line %s because %s", s, err)
	}
	return cube{x, y, z}
}

func neighbors(c cube) []cube {
	x := c.x
	y := c.y
	z := c.z
	return []cube{
		{x - 1, y, z},
		{x + 1, y, z},
		{x, y - 1, z},
		{x, y + 1, z},
		{x, y, z - 1},
		{x, y, z + 1},
	}
}

func getSurface(cubes map[cube]bool) int {
	coveredSurface := 0
	for cube := range cubes {
		for _, neighbor := range neighbors(cube) {
			if cubes[neighbor] {
				coveredSurface += 1
			}
		}
	}
	return 6*len(cubes) - coveredSurface
}

func part1(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	cubes := make(map[cube]bool)
	for scanner.Scan() {
		line := scanner.Text()
		cube := parseCube(line)
		cubes[cube] = true
	}
	println(getSurface(cubes))
}

func dfs(nonCubes map[cube]int, cubes map[cube]bool, minCube cube, maxCube cube, comp int, c cube) bool {
	nonCubes[c] = comp
	inner := true
	if c.x < minCube.x || c.x > maxCube.x || c.y < minCube.y || c.y > maxCube.y || c.z < minCube.z || c.z > maxCube.z {
		return false
	}
	for _, neighbor := range neighbors(c) {
		_, ok := nonCubes[neighbor]
		if !ok && !cubes[neighbor] {
			recInner := dfs(nonCubes, cubes, minCube, maxCube, comp, neighbor)
			if !recInner {
				inner = false
			}
		}
	}
	return inner
}

func part2(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	cubes := make(map[cube]bool)
	minCube := cube{1_000_000, 1_000_000, 1_000_000}
	maxCube := cube{0, 0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		cube := parseCube(line)
		cubes[cube] = true
		minCube.x = min(minCube.x, cube.x)
		minCube.y = min(minCube.y, cube.y)
		minCube.z = min(minCube.z, cube.z)
		maxCube.x = max(maxCube.x, cube.x)
		maxCube.y = max(maxCube.y, cube.y)
		maxCube.z = max(maxCube.z, cube.z)
	}

	nonCubes := make(map[cube]int)
	comp := 1
	compStatus := []bool{false} // offset by one
	for i := minCube.x; i <= maxCube.x; i++ {
		for j := minCube.y; j <= maxCube.y; j++ {
			for k := minCube.z; k <= maxCube.z; k++ {
				c := cube{i, j, k}
				if !cubes[c] && nonCubes[c] == 0 {
					inner := dfs(nonCubes, cubes, minCube, maxCube, comp, c)
					compStatus = append(compStatus, inner)
					comp += 1
				}
			}
		}
	}
	coveredSurface := 0
	for cube := range cubes {
		for _, neighbor := range neighbors(cube) {
			if cubes[neighbor] || compStatus[nonCubes[neighbor]] {
				coveredSurface += 1
			}
		}
	}
	println(6*len(cubes) - coveredSurface)

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
