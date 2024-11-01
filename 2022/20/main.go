package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func modAdd(i int, j int, mod int) int {
	res := i + j
	return (res%mod + mod) % mod
}

type pair struct {
	value, origIdx int
}

type mixlist struct {
	length int
	list   []pair
	find   map[pair]int
}

func mixlistFromSlice(s []int) mixlist {
	var list []pair
	find := make(map[pair]int)
	for i, v := range s {
		newElement := pair{v, i}
		list = append(list, newElement)
		find[newElement] = i
	}
	return mixlist{len(list), list, find}
}

func (m mixlist) swap(i int, j int) {
	m.list[i], m.list[j] = m.list[j], m.list[i]
	m.find[m.list[i]] = i
	m.find[m.list[j]] = j
}

func (m mixlist) move(p pair) {
	idx := m.find[p]
	moveVal := p.value % (m.length - 1)
	for i := range max(moveVal, -moveVal) {
		if p.value > 0 {
			m.swap(modAdd(idx, i, m.length), modAdd(idx, i+1, m.length))
		} else {
			m.swap(modAdd(idx, -i, m.length), modAdd(idx, -i-1, m.length))
		}
	}
}

func readMixlist(reader io.Reader, part int) mixlist {
	scanner := bufio.NewScanner(reader)
	var s []int
	for scanner.Scan() {
		line := scanner.Text()
		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(fmt.Errorf("could not read list element %s", err))
		}
		if part == 2 {
			x *= 811589153
		}
		s = append(s, x)
	}
	return mixlistFromSlice(s)
}

func part1(reader io.Reader) {
	m := readMixlist(reader, 1)
	list := make([]pair, len(m.list))
	copy(list, m.list)
	for _, v := range list {
		m.move(v)
	}
	res := 0
	zeroIdx := 0
	for i, v := range m.list {
		if v.value == 0 {
			zeroIdx = i
		}
	}
	for _, v := range []int{1000, 2000, 3000} {
		res += m.list[(zeroIdx+v)%len(m.list)].value
	}
	println(res)
}

func part2(reader io.Reader) {
	m := readMixlist(reader, 2)
	list := make([]pair, len(m.list))
	copy(list, m.list)
	for range 10 {
		for _, v := range list {
			m.move(v)
		}
	}
	res := 0
	zeroIdx := 0
	for i, v := range m.list {
		if v.value == 0 {
			zeroIdx = i
		}
	}
	for _, v := range []int{1000, 2000, 3000} {
		res += m.list[(zeroIdx+v)%len(m.list)].value
	}
	println(res)
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
