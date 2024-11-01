package main

import (
	"fmt"
	"os"
)

func isDistinct(s []byte) bool {
	for i := range s {
		for j := range s {
			if i != j && s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func part1(filename string) {
	filebytes, _ := os.ReadFile(filename)
	for i := range filebytes {
		if i < 3 {
			continue 
		}
		if isDistinct([]byte{filebytes[i-3], filebytes[i-2], filebytes[i-1], filebytes[i]}) {
			fmt.Println(i+1)
			return
		}
	}
	fmt.Println("Did not find signal")
}

func part2(filename string) {
	filebytes, _ := os.ReadFile(filename)
	for i := range filebytes {
		if i < 13 {
			continue 
		}
		s := []byte{}
		for j := 0; j < 14; j++ {
			s = append(s, filebytes[i-j])
		}

		if isDistinct(s) {
			fmt.Println(i+1)
			return
		}
	}
	fmt.Println("Did not find signal")
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

