import * as fs from "fs";

let input = fs.readFileSync("main.in").toString().split("\n").map(Number);

part1(input);
part2(input);

function part1(input) {
    let cnt = 0;
    for (let i = 1; i < input.length; i++) {
        if (input[i - 1] < input[i]) {
            cnt++;
        }
    }
    console.log("part1: " + cnt);
}

function part2(input) {
    let cnt = 0;
    for (let i = 3; i < input.length; i++) {
        if (input[i - 3] < input[i]) {
            cnt++;
        }
    }
    console.log("part2: " + cnt);
}
