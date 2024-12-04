import sys

with open(sys.argv[1]) as f:
    grid = f.readlines()
    height, width = len(grid), len(grid[0]) - 1
    grid = {(x, y): grid[x][y] for x in range(height) for y in range(width)}


def solve_one(grid: dict[tuple[int, int], str]) -> int:
    sum = 0
    search = (list("XMAS"),)
    directions = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

    for x, y in grid.keys():
        for dx, dy in directions:
            candidate = [grid.get((x + dx * i, y + dy * i), "") for i in range(4)]
            sum += candidate in search

    return sum


def solve_two(grid: dict[tuple[int, int], str]) -> int:
    sum = 0
    search = list("MAS"), list("SAM")
    directions = [-1, 0, 1]

    for x, y in grid.keys():
        candidate1 = [grid.get((x + d, y + d)) for d in directions] in search
        candidate2 = [grid.get((x + d, y - d)) for d in directions] in search
        sum += candidate1 and candidate2

    return sum


print("[part 1]:", solve_one(grid))
print("[part 2]:", solve_two(grid))
