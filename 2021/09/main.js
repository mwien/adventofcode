import * as fs from "fs";

//const input = fs.readFileSync("sample.in").toString().trim();
const input = fs.readFileSync("main.in").toString().trim();

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

function part2(input) {}

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
