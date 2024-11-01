import * as fs from "fs";

const input = fs.readFileSync("main.in").toString();

part1(input);

function part1(input) {
    input = input.split("\n");
    let [gamma, epsilon] = [0, 0];
    for (let i = 0; i < input[0].length; i++) {
        let [numZeros, numOnes] = [0, 0];
        for (let j = 0; j < input.length; j++) {
            if (input[j][i] == "0") {
                numZeros++;
            } else {
                numOnes++;
            }
        }
        if (numZeros > numOnes) {
            gamma = gamma * 2;
            epsilon = epsilon * 2 + 1;
        } else {
            gamma = gamma * 2 + 1;
            epsilon = epsilon * 2;
        }
    }
    console.log(gamma * epsilon);
}
