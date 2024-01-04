import re


with open("input.txt") as f:
    lines = f.readlines()
    lines = [line.strip() for line in lines]


def bag(line: str) -> dict:
    bag = {"red": 0, "green": 0, "blue": 0}
    matches = re.findall(r"((\d+) (\w+))", line)
    for _, n, color in matches:
        bag[color] = max(bag[color], int(n))
    return bag


sum_one = 0
sum_two = 0

for line in lines:
    game = bag(line)

    # part two - power of the set of cubes
    sum_two += game["red"] * game["green"] * game["blue"]

    # part one - skip lines with too many cubes
    if 12 < game["red"] or 13 < game["blue"] or 14 < game["green"]:
        continue

    game = re.findall("Game (\d+):", line)
    sum_one += int(game[0])

print("[part 1] sum:", sum_one)
print("[part 2] sum:", sum_two)
