import sys

UP = (-1, 0)
DOWN = (1, 0)
LEFT = (0, -1)
RIGHT = (0, 1)


class Grid:
    all_directions = [UP, DOWN, LEFT, RIGHT]
    directions = {
        "|": [UP, DOWN],
        "-": [LEFT, RIGHT],
        "L": [UP, RIGHT],
        "J": [UP, LEFT],
        "7": [DOWN, LEFT],
        "F": [DOWN, RIGHT],
        "S": [UP, DOWN, LEFT, RIGHT],
    }

    def __init__(self, input: str) -> None:
        self.xy = []
        self.start = (0, 0)
        self.visited = {}

        lines = [line.strip() for line in input.split("\n")]
        for i, line in enumerate(lines):
            self.xy.append([c for c in line.strip()])
            if "S" in self.xy[-1]:
                self.start = (i, self.xy[-1].index("S"))

    def solve_one(self) -> int:
        queue = [(self.start, 0)]
        while 0 < len(queue):
            current, distance = queue.pop(0)
            if current in self.visited:
                continue
            self.visited[current] = distance
            dir = self.directions[self.xy[current[0]][current[1]]]
            for next in self.find_moves(dir, current, distance):
                queue.append(next)

        return max(self.visited.values())

    def solve_two(self) -> int:
        reachable = self.find_moves(self.all_directions, self.start)
        for pipe in self.directions:
            if len(reachable) != len(self.directions[pipe]):
                continue
            if all([d in self.directions[pipe] for d in reachable]):
                self.xy[self.start[0]][self.start[1]] = pipe

        sum = 0
        for i in range(len(self.xy)):
            enclosed = 0
            for j in range(len(self.xy[i])):
                if (i, j) in self.visited:
                    enclosed += UP in self.directions[self.xy[i][j]]
                    continue
                if enclosed % 2 == 1:
                    self.xy[i][j] = "I"
                sum += self.xy[i][j] == "I"

        return sum

    def find_moves(self, directions: list, current: tuple, distance: int = -1) -> list:
        part_two = distance < 0
        i, j = current
        moves = []
        for di, dj in directions:
            if i + di < 0 or len(self.xy) <= i + di:
                continue
            if j + dj < 0 or len(self.xy[i + di]) <= j + dj:
                continue
            target = self.xy[i + di][j + dj]
            if target not in self.directions:
                continue
            opposite = (di * -1, dj * -1)
            if opposite in self.directions[target]:
                if part_two:
                    moves.append((di, dj))
                else:
                    moves.append(((i + di, j + dj), distance + 1))

        return moves


with open(sys.argv[1]) as f:
    g = Grid(f.read())


print("[part 1] distance:", g.solve_one())
print("[part 2] enclosed:", g.solve_two())
