# Advent of Code 2022 - Day 2: Rock Paper Scissors

Puzzle details are available at <https://adventofcode.com/2022/day/2>.

This puzzle is solved with [Go](https://go.dev).

## tldr: Solutions

- [part one](./part_one.go)
- [part two](./part_two.go)
- [part one & two combined](./main.go)

## Abstract

- We have an encrypted strategy guide list for the game Rock Paper Scissors.
- Each line contains the choices made by the opponent and the player.
- The first column is the move made by the opponent and the second by the player.
- Each move has a score.
- Each round outcome has a score.
- The player with the total highest score is the winner.

An example list:

```text
A Y
B X
C Z
```

## Part One - Total score

> tldr: [solution](./part_one.go)

Since the strategy list is encrypted, we need to decrypt it using the following key:

| Shape    | Keys |
| -------- | ---- |
| Rock     | A, X |
| Paper    | B, Y |
| Scissors | C, Z |

There are two scores, one for the shape played, and the other for the outcome:

| Shape    | Score |
| -------- | ----- |
| Rock     | 1     |
| Paper    | 2     |
| Scissors | 3     |

| Outcome    | Score |
| ---------- | ----- |
| Win        | 6     |
| Draw (tie) | 3     |
| Lose       | 0     |

### Open the file and read the lines

We will start by opening the file containing the data and assume the file is called `input.txt`.

```go
f, err := os.Open("input.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

The following happens above:

1. The file is opened for reading.
2. If there is an error opening it (e.g. the file does not exist), an error is logged, and the program terminates.
3. The file is closed once it is no longer needed.

To read the lines, we will use `bufio.NewScanner()` like this:

```go
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
}
```

### Player and opponent moves

We can then split each line using the space as a separator to get the opponent and the player.
One way of doing it is like this:

```go
line := scanner.Text()
columns := strings.Split(line, " ")
opponent := columns[0] // first column
player := columns[1] // second column
```

### Scoring the moves

To score the move and the outcome, we start by setting the score outside the loop first.

```go
score := 0
for scanner.Scan() {
}
```

After extracting the moves, we need to score them using the shape score as defined above.

```go
if player == "X" { // Rock
    score += 1
} else if player == "Y" { // Paper
    score += 2
} else if player == "Z" { // Scissors
    score += 3
}
```

Here comes the tricky part, we need to figure out the outcome of the round.
To do that, we need to compare the player's move to the opponent's using the following rules of the game:

| opponent     | player       | outcome | score |
| ------------ | ------------ | ------- | ----- |
| A (Rock)     | X (Rock)     | draw    | 3     |
| A (Rock)     | Y (Paper)    | win     | 6     |
| A (Rock)     | Z (Scissors) | lose    | 0     |
| B (Paper)    | X (Rock)     | lose    | 0     |
| B (Paper)    | Y (Paper)    | draw    | 3     |
| B (Paper)    | Z (Scissors) | win     | 6     |
| C (Scissors) | X (Rock)     | win     | 6     |
| C (Scissors) | Y (Paper)    | lose    | 0     |
| C (Scissors) | Z (Scissors) | draw    | 3     |

Translating this into conditions in the code is cumbersome and error-prone.
We could simplify it by checking whether `player == opponent` and adding `3` to the total score, but that still leaves us with 6 additional conditions.

We can create an outcome table using the above values like this:

```go
outcome := map[string]map[string]int{
    "A": {
        "X": 3,
        "Y": 6,
        "Z": 0,
    },
    "B": {
        "X": 0,
        "Y": 3,
        "Z": 6,
    },
    "C": {
        "X": 6,
        "Y": 0,
        "Z": 3,
    },
}
```

The only thing left to do is to add the outcome to the total score:

```go
score += lookup[opponent][player]
```

### Simplify the code

Since there are only nine possible outcomes, we can simplify the code by calculating the outcome in advance.
We take the values from the above outcome table, flatten them and add the two values together.

For example, we know that:

- `A` (Rock) and played against `X=1` (Rock) yields a `draw=3`, resulting in the `score=4` (1+3).
- `A` (Rock) and played against `Y=2` (Paper) yields a `win=6`, resulting in the `score=8` (2+6).
- `A` (Rock) and played against `Z=3` (Scissors) yields a `lose=0`, resulting in the `score=3` (3+0).

We can go ahead and change the outcome lookup table to this:

```go
outcome := map[string]int{
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
```

With this change, we can go ahead and remove the lines where the player and opponent's moves are extracted, and the player's move is scored.

In other words, we now need to change the lines inside the loop to this:

```go
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
    score += outcome[line]
}
```

And we are done.

## Part Two - Total score

> tldr: [solution](./part_two.go)

The second part of the puzzle changes the value of the second column.
Instead of the shape, the second columns now represents the outcome of the round using the following rules:

| Value | Outcome |
| ----- | ------- |
| X     | lose    |
| Y     | draw    |
| Z     | win     |

For example, if the line reads `A Y`, the opponent chooses Rock (`A`) and the player needs to draw (`Y`), meaning the player should choose the Rock as well.

The scoring stays the same.

### Implementation

The implementation can be done following the same steps as in the first part.
However, since we already saw how it can be simplified, we can update our outcome lookup table and the rest of the code stays the same:

```go
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
```

## Part One and Two - Combined solution

Take a look at the [combined code](./main.go) for details.
