package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	maxCalories := 0
	elfCalories := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if elfCalories > maxCalories {
				maxCalories = elfCalories
			}
			elfCalories = 0
		} else {
			cal, _ := strconv.Atoi(line)
			elfCalories += cal
		}
	}
	fmt.Println("calories:", maxCalories)
}
