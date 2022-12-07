package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	directories := map[string]int{}
	pwd := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		// next lines are far from optimal
		// but good enough to solve the puzzle :)
		if line[1] == "cd" && line[2] == ".." {
			pwd = pwd[:len(pwd)-1]
		} else if line[1] == "cd" {
			pwd = append(pwd, line[2])
		}
		if size, err := strconv.Atoi(line[0]); err == nil {
			path := ""
			for _, dir := range pwd {
				path += dir + "/"
				directories[path[1:]] += size
			}
		}
	}

	// part 1
	total := 0
	for _, size := range directories {
		if size < 100000 {
			total += size
		}
	}

	// part 2
	smallest := directories["/"]
	min := 30000000 - (70000000 - smallest)
	for _, size := range directories {
		if size >= min && size < smallest {
			smallest = size
		}
	}

	fmt.Println("[part 1] total:", total)
	fmt.Println("[part 2] smallest:", smallest)
}
