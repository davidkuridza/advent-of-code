package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	overlapOne := 0
	overlapTwo := 0

	re := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		d, _ := strconv.Atoi(match[4])

		// part one: full overlap
		if (in(a, b, c) && in(a, b, d)) || (in(c, d, a) && in(c, d, b)) {
			overlapOne += 1
		}
		// part two: general overlap
		if in(a, b, c) || in(a, b, d) || in(c, d, a) || in(c, d, b) {
			overlapTwo += 1
		}
	}

	fmt.Println("[part 1] overlaps:", overlapOne)
	fmt.Println("[part 2] overlaps:", overlapTwo)
}

func in(from, to, candidate int) bool {
	return from <= candidate && candidate <= to
}
