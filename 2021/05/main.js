import * as fs from "fs";

//const input = fs.readFileSync("sample.in").toString();
const input = fs.readFileSync("main.in").toString();

part1(input);
part2(input);

function part1(input) {
  input = input.split("\n");
  const grid = new Map();
  for (const line of input) {
    if (line == "") {
      continue;
    }
    const [p1, p2] = parseLine(line);
    if (p1[0] != p2[0] && p1[1] != p2[1]) {
      continue;
    }
    for (const row of between(p1[0], p2[0])) {
      for (const col of between(p1[1], p2[1])) {
        updateGrid(grid, row, col);
      }
    }
  }
  console.log(countIntersections(grid));
}

function part2(input) {
  input = input.split("\n");
  const grid = new Map();
  for (const line of input) {
    if (line == "") {
      continue;
    }
    const [p1, p2] = parseLine(line);
    const rows = between(p1[0], p2[0]);
    const cols = between(p1[1], p2[1]);
    const n = Math.max(rows.length, cols.length);
    for (let i = 0; i < n; i++) {
      updateGrid(grid, rows[i % rows.length], cols[i % cols.length]);
    }
  }
  console.log(countIntersections(grid));
}

function parseLine(line) {
  return line.split("->").map((tuple) => {
    return tuple.trim().split(",").map(Number);
  });
}

function between(a, b) {
  const res = [];
  if (a <= b) {
    for (let i = a; i <= b; i++) {
      res.push(i);
    }
  } else {
    for (let i = a; i >= b; i--) {
      res.push(i);
    }
  }
  return res;
}

function updateGrid(grid, row, col) {
  const key = row + "," + col;
  const val = grid.get(key);
  if (val === undefined) {
    grid.set(key, 1);
  } else {
    grid.set(key, val + 1);
  }
}

function countIntersections(grid) {
  let count = 0;
  for (const [_, val] of grid) {
    if (val > 1) {
      count += 1;
    }
  }
  return count;
}
