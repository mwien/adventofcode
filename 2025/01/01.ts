import * as fs from "fs";

let input = fs.readFileSync("main.in").toString().split("\n");

part1(input);
part2(input);

function part1(input) {
  let pos = 50;
  let pwd = 0;
  for (const line of input) {
    let rot = Number(line.substring(1));
    if (line[0] === "L") {
      pos -= rot;
    } else {
      pos += rot;
    }
    if (pos % 100 === 0) {
      pwd += 1;
    }
  }
  console.log(pwd);
}

function part2(input) {
  let pos = 50;
  let pwd = 0;
  for (const line of input) {
    let rot = Number(line.substring(1));
    for (let i = 0; i < rot; i++) {
      if (line[0] === "L") {
        pos -= 1;
      } else {
        pos += 1;
      }
      if (pos % 100 === 0) {
        pwd += 1;
      }
    }
  }
  console.log(pwd);
}
