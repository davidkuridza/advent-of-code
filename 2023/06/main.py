import re
import sys


def solve(time: int, distance: int) -> int:
    ways = 0
    for hold in range(time):
        speed = hold
        travel = speed * (time - hold)
        if distance < travel:
            ways += 1

    return ways


with open(sys.argv[1]) as f:
    lines = f.read().split("\n")
    lines = [line.strip() for line in lines]


times = [int(n) for n in re.findall("\d+", lines[0])]
distances = [int(n) for n in re.findall("\d+", lines[1])]

part_one = 1
for time, distance in zip(times, distances):
    part_one *= solve(time, distance)

print("[part 1] solution:", part_one)

time = int("".join(map(str, times)))
distance = int("".join(map(str, distances)))

part_two = solve(time, distance)
print("[part 1] solution:", part_two)
