import * as fs from "fs";

const input = fs.readFileSync(0, "utf-8").trim();

part1(input);
part2(input);

function part1(input) {
  const grid = input.split("\n").map((line) => {
    return line.split("").map(Number);
  });
  const rows = grid.length;
  const cols = grid[0].length;
  let risk = 0;
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      const mn = neighbors(i, j, rows, cols).reduce(
        (mn, pos) => Math.min(grid[pos[0]][pos[1]], mn),
        10000,
      );
      if (mn > grid[i][j]) {
        risk += grid[i][j] + 1;
      }
    }
  }
  console.log(risk);
}

function part2(input) {
  const grid = input.split("\n").map((line) => {
    return line.split("").map(Number);
  });
  const rows = grid.length;
  const cols = grid[0].length;

  const vis = Array.from({ length: rows }, () => Array(cols).fill(false));

  const comps = [];
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (vis[i][j] || grid[i][j] == 9) {
        continue;
      }
      comps.push(componentSize(grid, vis, i, j, rows, cols));
    }
  }
  comps.sort((a, b) => b - a);
  console.log(comps[0] * comps[1] * comps[2]);
}

function componentSize(grid, vis, i, j, rows, cols) {
  let sz = 0;
  const stack = [[i, j]];
  vis[i][j] = true;
  while (stack.length > 0) {
    const [u, v] = stack.pop();
    sz += 1;
    for (const [x, y] of neighbors(u, v, rows, cols)) {
      if (!vis[x][y] && grid[x][y] != 9) {
        stack.push([x, y]);
        vis[x][y] = true;
      }
    }
  }
  return sz;
}

function neighbors(i, j, rows, cols) {
  const result = [];
  for (let x = -1; x <= 1; x++) {
    for (let y = -1; y <= 1; y++) {
      if (
        (Math.abs(x) + Math.abs(y)) % 2 == 1 &&
        i + x >= 0 &&
        i + x < rows &&
        j + y >= 0 &&
        j + y < cols
      ) {
        result.push([i + x, j + y]);
      }
    }
  }
  return result;
}
