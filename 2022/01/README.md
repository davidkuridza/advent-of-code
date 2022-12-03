# Advent of Code 2022 - Day 1: Calorie Counting

Puzzle details are available at <https://adventofcode.com/2022/day/1>.

This puzzle is solved with [Go](https://go.dev).

## tldr: Solutions

- [part one](./part_one.go)
- [part two](./part_two.go)
- [part one & two combined](./main.go)

## Abstract

- We have a list of Elves carrying items with a certain amount of Calories.
- Each Elf may carry one or more items.
- Calories for each item are presented on a single line.
- Elves' items (inventory) separated by a new line.
- The sum of items in Elf's inventory defines the number of Calories they carry.

An example list:

```text
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
```

## Part One - Find the Elf carrying the most Calories

> tldr: [solution](./part_one.go)

To find the Elf carrying the most Calories, we need to figure out how many Calories worth of items each Elf carries.
We know the following:

- each line contains the Calories for a single item,
- each Elf's inventory separated by a new line.

This means we can read the list line by line and search for a blank line to determine the inventory per Elf.

### Sum of all Calories

Let's start by calculating the number of all Calories on the list:

1. Set the initial value of Calories to `0` before reading the lines.
2. Add the value of each line to the total sum of Calories.
3. Print the final value.

```go
calories := 0
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
    cal, _ := strconv.Atoi(line)
    calories += cal
}
fmt.Println(calories)
```

### Maximum Calories

To find the maximum Calories an Elf carries we can keep track of the current highest value.
Alternatively, we can track the Calories each Elf carries in a slice and find the maximum value.
To keep it simple, we will use the first option.

*Spoiler alert, we will use the second option in [part two](#part-two---find-the-calories-carried-by-the-top-three-elves).*

The steps needed:

1. Keep track of maximum Calories in `maxCalories`.
2. Keep track of current Elf's Calories in `elfCalories`.
3. When the line is empty:
   1. Set `maxCalories` with the current Elf's Calories if the value is higher.
   2. Reset current Elf's Calories.
4. When the line is not empty, add the value to the current Elf's calorie count.

```go
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
fmt.Println(maxCalories)
```

## Part Two - Find the Calories carried by the top three Elves

> tldr: [solution](./part_two.go)

Part two of the puzzle requires us to find the number of Calories carried by the top three Elves.
We can find them by keeping track of the highest amount of Calories the same way as we did in [Part One](#part-one---find-the-elf-carrying-the-most-calories) using three variables.
Since this approach is cumbersome, we will track them in a list.

The steps needed:

1. Keep track of Calories per Elf in the `calories` slice of integers.
2. Keep track of current Elf's Calories in `elfCalories`.
3. When the line is empty:
   1. Append the current Elf's Calories to the slice.
   2. Reset current Elf's Calories.
4. When the line is not empty, add the value to the current Elf's calorie count.

```go
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
```

### Top three Elves

The above code gives us an unordered slice with sums of Calories for each Elf.
We can sort the slice and take the last three elements or reverse-sort it and take the first three elements.
We will do the latter.

```go
sort.Slice(calories, func(i, j int) bool {
    return calories[i] > calories[j]
})
fmt.Println("calories:", calories[0]+calories[1]+calories[2])
```

## Part One and Two - Combined solution

> tldr: [solution](./main.go)

We can now go ahead and combine both parts by modifying the second solution.
We need to add another line at the end to output the first element in the sorted slice.

```go
fmt.Println("calories:", calories[0])
```

To make it clear which Calories are which, we can modify the output to something like this:

```go
fmt.Println("[part 1] calories:", calories[0])
fmt.Println("[part 2] calories:", calories[0]+calories[1]+calories[2])
```
