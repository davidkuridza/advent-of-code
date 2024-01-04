package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sumPartOne := 0
	sumPartTwo := 0
	set := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// part one
		middle := len(line) / 2
		sumPartOne += priority(
			duplicates(
				line[:middle], // first half
				line[middle:], // second half
			),
		)
		// part two
		set = append(set, line)
		if len(set) == 3 {
			sumPartTwo += priority(
				duplicates(
					string(duplicates(set[0], set[1])),
					set[2],
				),
			)
			set = []string{}
		}
	}

	fmt.Println("[part 1] sum:", sumPartOne)
	fmt.Println("[part 2] sum:", sumPartTwo)
}

func duplicates(a, b string) []rune {
	set := []rune{}
	hash := make(map[rune]struct{})
	unique := make(map[rune]struct{})

	for _, r := range a {
		hash[r] = struct{}{}
	}

	for _, r := range b {
		if _, exists := hash[r]; exists {
			if _, visited := unique[r]; !visited {
				set = append(set, r)
				unique[r] = struct{}{}
			}
		}
	}
	return set
}

func priority(items []rune) int {
	sum := 0
	for _, r := range items {
		p := int(r)
		if p >= int('a') && p <= int('z') {
			sum += p - 96
		} else {
			sum += p - 38
		}
	}
	return sum
}
