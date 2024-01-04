package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	calories := []int{}
	elfCalories := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			calories = append(calories, elfCalories)
			elfCalories = 0
		} else {
			cal, _ := strconv.Atoi(line)
			elfCalories += cal
		}
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	fmt.Println("[part 1] calories:", calories[0])
	fmt.Println("[part 2] calories:", calories[0]+calories[1]+calories[2])
}
