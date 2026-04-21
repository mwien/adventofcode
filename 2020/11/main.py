import sys
import copy


class Position:
    def __init__(self, i: int, j: int):
        self.i = i
        self.j = j

    def __repr__(self):
        return f"Position({self.i}, {self.j})"

    def __add__(self, other: "Position") -> "Position":
        return Position(self.i + other.i, self.j + other.j)

    def __iadd__(self, other: "Position") -> "Position":
        self.i += other.i
        self.j += other.j
        return self

    def in_bounds(self, xdim: int, ydim: int) -> bool:
        return 0 <= self.i < xdim and 0 <= self.j < ydim


class Grid:
    def __init__(self, grid: list[list[str]]):
        self.n = len(grid)
        self.m = len(grid[0])
        self.grid = grid
        self.next_seats = [[[] for _ in range(self.m)] for _ in range(self.n)]

    def __repr__(self):
        return f"Grid({self.n}, {self.m}, {self.grid})"

    def __getitem__(self, pos: Position):
        return self.grid[pos.i][pos.j]

    def __setitem__(self, pos, value):
        self.grid[pos.i][pos.j] = value

    def all_positions(self) -> list[Position]:
        return [Position(i, j) for i in range(self.n) for j in range(self.m)]

    def populate_next_seats(self, only_direct=False):
        deltas = [(x, y) for x in [-1, 0, 1] for y in [-1, 0, 1] if x != 0 or y != 0]
        for pos in self.all_positions():
            for x, y in deltas:
                k = 0
                new_pos = copy.copy(pos)
                while True:
                    if only_direct and k >= 1:
                        break
                    new_pos += Position(x, y)
                    if not new_pos.in_bounds(self.n, self.m):
                        break
                    if self[new_pos] == "#" or self[new_pos] == "L":
                        self.next_seats[pos.i][pos.j].append(new_pos)
                        break
                    k += 1

    def neighbors_occupied(self, pos: Position):
        cnt = 0
        for npos in self.next_seats[pos.i][pos.j]:
            if self[npos] == "#":
                cnt += 1
        return cnt

    def number_occupied(self):
        cnt = 0
        for pos in self.all_positions():
            if self[pos] == "#":
                cnt += 1
        return cnt

    def next_round(self, cutoff):
        new_grid = copy.deepcopy(self)
        changed = False
        for pos in self.all_positions():
            entry = self[pos]
            # print(f"{pos}: {self.neighbors_occupied(pos)}")
            if entry == "L" and not self.neighbors_occupied(pos):
                new_grid[pos] = "#"
                changed = True
            elif entry == "#" and self.neighbors_occupied(pos) >= cutoff:
                new_grid[pos] = "L"
                changed = True
        return new_grid, changed


def solve(grid: Grid, only_direct=False):
    grid = copy.deepcopy(grid)
    grid.populate_next_seats(only_direct=only_direct)
    # for i in range(grid.n):
    #     for j in range(grid.m):
    #         print(f"{i} {j}: {grid.next_seats[i][j]}")
    while True:
        # print(grid)
        grid, changed = grid.next_round(4 if only_direct else 5)
        if not changed:
            print(grid.number_occupied())
            break


grid = Grid([list(line.strip()) for line in sys.stdin])
solve(grid, only_direct=True)
solve(grid, only_direct=False)
