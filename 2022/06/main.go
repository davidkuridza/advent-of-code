package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[part  1] marker:", solve(string(b), 4))
	fmt.Println("[part  2] marker:", solve(string(b), 14))
}

func solve(line string, marker int) int {
	chars := ""
	for i, r := range line {
		if i := strings.IndexRune(chars, r); i > -1 {
			chars = chars[i+1:]
		}
		chars += string(r)
		if len(chars) == marker {
			return i + 1
		}
	}
	return -1
}
