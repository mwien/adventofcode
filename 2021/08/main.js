import * as fs from "fs";

const input = fs.readFileSync("sample.in").toString().trim();
//const input = fs.readFileSync("main.in").toString().trim();

part1(input);
part2(input);

function part1(input) {
  const lines = input.split("\n");
  let res = 0;
  for (const line of lines) {
    const [_, digits] = line.split(" | ");
    digits.split(" ").forEach((x) => {
      if ([2, 3, 4, 7].includes(x.length)) {
        res += 1;
      }
    });
  }
  console.log(res);
}

function part2(input) {
  let res = 0;
  input.split("\n").forEach((line) => {
    const [pattern, digits] = line.split(" | ").map((x) => {
      return x.split(" ");
    });
    res += reconstructLine(pattern, digits);
  });
}

function reconstructLine(pattern, digits) {
  // find potential positions for each letter
  // do matching
}
