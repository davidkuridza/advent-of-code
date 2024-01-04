import math
import re
import sys


def solve(cursor: str, end_with: str) -> int:
    steps = 0
    while not cursor.endswith(end_with):
        instruction = instructions[steps % len(instructions)]
        i = 0 if instruction == "L" else 1
        cursor = network[cursor][i]
        steps += 1
    return steps


with open(sys.argv[1]) as f:
    lines = f.readlines()
    lines = [line.strip() for line in lines]

instructions = list(lines[0])
network = {}

for line in lines[2:]:
    node, left, right = re.findall("\w+", line)
    network[node] = (left, right)

# part one
steps = solve("AAA", "ZZZ")
print("[part 1] steps:", steps)

# part two
steps = []
for cursor in network:
    if cursor.endswith("A"):
        steps.append(solve(cursor, "Z"))

steps = math.lcm(*steps)
print("[part 2] steps:", steps)
