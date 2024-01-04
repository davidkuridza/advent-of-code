package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func top(stacks map[int]string) string {
	top := ""
	for i := 1; i <= len(stacks); i++ {
		top += stacks[i][0:1]
	}
	return top
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(b), "\n\n")
	stacksOne := map[int]string{}
	stacksTwo := map[int]string{}

	// stacks
	for _, s := range strings.Split(parts[0], "\n") {
		for i := 1; i < len(s); i += 4 {
			r := s[i]
			if r >= 'A' && r <= 'Z' {
				column := int(math.Ceil(float64(i) / 4))
				stacksOne[column] += string(r)
				stacksTwo[column] += string(r)
			}
		}
	}

	// moves
	re := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	for _, m := range re.FindAllStringSubmatch(parts[1], -1) {
		from, _ := strconv.Atoi(m[2])
		to, _ := strconv.Atoi(m[3])
		move, _ := strconv.Atoi(m[1])

		stacksOne[to] = reverse(stacksOne[from][0:move]) + stacksOne[to]
		stacksOne[from] = stacksOne[from][move:]

		stacksTwo[to] = stacksTwo[from][0:move] + stacksTwo[to]
		stacksTwo[from] = stacksTwo[from][move:]
	}

	fmt.Println("[part 1] top crates", top(stacksOne))
	fmt.Println("[part 2] top crates", top(stacksTwo))
}
