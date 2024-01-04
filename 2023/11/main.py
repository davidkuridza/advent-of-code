import sys


class Universe:
    def __init__(self, input: str) -> None:
        lines = [line.strip() for line in input.splitlines()]
        self.universe = [[0 if i == "." else 1 for i in line] for line in lines]

        self.rows = [i for i, row in enumerate(self.universe) if 0 == sum(row)]
        self.cols = [j for j, col in enumerate(zip(*self.universe)) if 0 == sum(col)]

        self.galaxies = []
        for x in range(len(self.universe[0])):
            for y in range(len(self.universe)):
                if self.universe[y][x] != 0:
                    self.galaxies.append((x, y))

    def solve(self, factor: int) -> int:
        distance = 0
        for i in range(len(self.galaxies) - 1):
            for j in range(i + 1, len(self.galaxies)):
                xs = sorted([self.galaxies[i][0], self.galaxies[j][0]])
                ys = sorted([self.galaxies[i][1], self.galaxies[j][1]])
                nx = self.expanded(self.cols, xs)
                ny = self.expanded(self.rows, ys)
                distance += xs[1] - xs[0] + ys[1] - ys[0] + (nx + ny) * factor

        return distance

    def expanded(self, input: list, xy: list) -> int:
        n = 0
        for i in input:
            if i > xy[0] and i < xy[1]:
                n += 1
        return n


with open(sys.argv[1]) as f:
    u = Universe(f.read())

print("[part 1] distance:", u.solve(1))
print("[part 2] distance:", u.solve(999999))
