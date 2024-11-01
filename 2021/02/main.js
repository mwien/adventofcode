import * as fs from "fs";

const input = fs.readFileSync("main.in").toString();

part1(input);
part2(input);

function part1(input) {
    input = input.split("\n");
    let [x, y] = [0, 0];
    for (const line of input) {
        const [command, stepSizeString] = line.split(" ");
        const stepSize = Number(stepSizeString);
        switch (command) {
            case "forward":
                x += stepSize;
                break;
            case "down":
                y += stepSize;
                break;
            case "up":
                y -= stepSize;
                break;
            default:
                break;
        }
    }
    console.log(x * y);
}

function part2(input) {
    input = input.split("\n");
    let [x, y, aim] = [0, 0, 0];
    for (let line of input) {
        let [command, stepSizeString] = line.split(" ");
        let stepSize = Number(stepSizeString);
        switch (command) {
            case "forward":
                x += stepSize;
                y += aim * stepSize;
                break;
            case "down":
                aim += stepSize;
                break;
            case "up":
                aim -= stepSize;
                break;
            default:
                break;
        }
    }
    console.log(x * y);
}
