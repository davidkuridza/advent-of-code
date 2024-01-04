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

	// Rock     = A = X = 1
	// Paper    = B = Y = 2
	// Scissors = C = Z = 3

	// lose = 0
	// draw = 3
	// win  = 6
	outcomePartOne := map[string]int{
		"A X": 4, // Rock: Rock + draw = 1+3
		"A Y": 8, // Rock: Paper + win  = 2+6
		"A Z": 3, // Rock: Scissors + lose = 3+0
		"B X": 1, // Paper: Rock + lose = 1+0
		"B Y": 5, // Paper: Paper + draw  = 2+3
		"B Z": 9, // Paper: Scissors + win = 3+6
		"C X": 7, // Scissors: Rock + win = 1+6
		"C Y": 2, // Scissors: Paper + lose  = 2+0
		"C Z": 6, // Scissors: Scissors + draw = 3+3
	}
	// X = lose
	// Y = draw
	// Z = win
	outcomePartTwo := map[string]int{
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

	scorePartOne := 0
	scorePartTwo := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		scorePartOne += outcomePartOne[line]
		scorePartTwo += outcomePartTwo[line]
	}

	fmt.Println("[part 1] score:", scorePartOne)
	fmt.Println("[part 2] score:", scorePartTwo)
}
