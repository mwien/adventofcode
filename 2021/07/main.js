import * as fs from "fs";

//const input = fs.readFileSync("sample.in").toString();
const input = fs.readFileSync("main.in").toString();

part1(input);
part2(input);

function part1(input) {
  const positions = input.split(",").map(Number);
  positions.sort((a, b) => a - b);
  const median = positions[positions.length / 2];
  let res = 0;
  for (const pos of positions) {
    res += Math.abs(median - pos);
  }
  console.log(res);
}

function part2(input) {
  const positions = input.split(",").map(Number);
  positions.sort((a, b) => a - b);
  const max = positions[positions.length - 1];
  let res = 1e9;
  for (let i = 0; i <= max; i++) {
    const totalFuel = getFuel(positions, i);
    res = Math.min(res, totalFuel);
  }
  console.log(res);
}

function getFuel(positions, x) {
  let sum = 0;
  for (const pos of positions) {
    const d = Math.abs(x - pos);
    sum += ((d + 1) * d) / 2;
  }
  return sum;
}
