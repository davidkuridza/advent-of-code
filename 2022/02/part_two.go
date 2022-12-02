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

	outcome := map[string]int{
		"A X": 3, // Rock: Scissors + lose = 3+0
		"A Y": 4, // Rock: Rock + draw  = 1+3
		"A Z": 8, // Rock: Paper + win = 2+6
		"B X": 1, // Paper: Rock + lose = 1+0
		"B Y": 5, // Paper: Paper + draw  = 2+3
		"B Z": 9, // Paper: Scissors + win = 3+6
		"C X": 2, // Scissors: Paper + lose = 2+0
		"C Y": 6, // Scissors: Scissors + draw  = 3+3
		"C Z": 7, // Scissors: Rock + win = 1+6
	}
	score := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		score += outcome[line]
	}

	fmt.Println("score:", score)
}
