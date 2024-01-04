import re
from typing import Tuple, Union

with open("input.txt") as f:
    grid = f.readlines()
    grid = [line.strip() for line in grid]

directions = [
    (-1, -1),
    (-1, 0),
    (-1, 1),
    (0, -1),
    (0, 0),
    (0, 1),
    (1, -1),
    (1, 0),
    (1, 1),
]
grid_y_len = len(grid)
grid_x_len = len(grid[0])


def is_adjacent(i: int, j: int) -> bool:
    return get_coordinates(i, j, False)


def get_coordinates(
    i: int, j: int, asterix: bool = True
) -> Union[Tuple[int, int], bool]:
    for di, dj in directions:
        ni, nj = i + di, j + dj
        if 0 <= ni < grid_y_len and 0 <= nj < grid_y_len:
            if asterix and grid[ni][nj] == "*":
                return ni, nj
            if not asterix and re.search("[^\d.]", grid[ni][nj]) is not None:
                return True

    return False


part_one = 0
part_two = 0
asterix = {}
for row_y, line in enumerate(grid):
    matches = re.finditer("\d+", line)
    for m in matches:
        number = int(m.group())
        row_x_start = m.start()
        row_x_end = m.end()

        # part one
        for row_x in range(row_x_start, row_x_end):
            if is_adjacent(row_y, row_x):
                part_one += number
                break

        # part two
        for row_x in range(row_x_start, row_x_end):
            xy = get_coordinates(row_y, row_x)
            if not xy:
                continue

            if xy not in asterix:
                asterix[xy] = {"occurrence": 1, "ratio": number}
            else:
                asterix[xy]["occurrence"] += 1
                asterix[xy]["ratio"] *= number
            break

    part_two = sum(v["ratio"] for v in asterix.values() if v["occurrence"] == 2)

print("[part 1] sum:", part_one)
print("[part 2] sum:", part_two)
