import sys
import re
from collections import defaultdict

with open(sys.argv[1]) as f:
    lines = f.read().split("\n")
    lines = [line.strip() for line in lines]

seeds = [int(n) for n in re.findall("\d+", lines[0])]
maps = []

for line in lines[2:]:
    if "map" in line:
        maps.append([])
        continue

    if line == "":
        continue

    # destination, source, range
    maps[len(maps) - 1].append([int(n) for n in line.split()])

print(maps)


def solve_one(seed, map):
    for destination, source, count in map:
        if source <= seed < source + count:
            return seed + destination - source
    return seed


def solve_two(range, map):
    range_in = []
    for destination, source, count in map:
        range_new = []
        while range:
            (start, end) = range.pop()
            before = (start, min(end, source))
            inner = (max(start, source), min(source + count, end))
            after = (max(source + count, start), end)
            if inner[0] < inner[1]:
                r = (inner[0] - source + destination, inner[1] - source + destination)
                range_in.append(r)
            if before[0] < before[1]:
                range_new.append(before)
            if after[0] < after[1]:
                range_new.append(after)
        range = range_new
    return range_in + range


locations = []
for seed in seeds:
    for map in maps:
        seed = solve_one(seed, map)
    locations.append(seed)

print("[part 1] location:", min(locations))

locations = []
pairs = list(zip(seeds[::2], seeds[1::2]))
for start, end in pairs:
    range = [(start, start + end)]
    for map in maps:
        range = solve_two(range, map)
    locations.append(min(range)[0])

print("[part 2] location:", min(locations))
