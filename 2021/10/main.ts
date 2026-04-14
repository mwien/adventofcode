import { readFileSync } from "fs";

const input = readFileSync(0, "utf-8").trim().split("\n");

const partner: { [key: string]: string } = {
  ")": "(",
  "]": "[",
  "}": "{",
  ">": "<",
};

const closing: { [key: string]: string } = {
  "(": ")",
  "[": "]",
  "{": "}",
  "<": ">",
};

part1();
part2();

function part1() {
  let res = 0;
  const points: { [key: string]: number } = {
    ")": 3,
    "]": 57,
    "}": 1197,
    ">": 25137,
  };

  for (const line of input) {
    const c = firstIllegal(line);
    if (c === null) {
      continue;
    }
    res += points[c]!;
  }

  console.log(res);
}

function part2() {
  let scores = [];
  for (const line of input) {
    if (firstIllegal(line) !== null) {
      continue;
    }
    const c = completion(line);
    console.log(c);
    scores.push(computeScore(c));
    console.log(scores.slice(-1));
  }
  scores = scores.sort((a, b) => {
    return a - b;
  });
  console.log(scores[Math.floor(scores.length / 2)]);
}

function firstIllegal(s: string): string | null {
  const stack = [];
  for (let i = 0; i < s.length; i++) {
    let c = s.charAt(i);
    if (c === "{" || c === "(" || c === "[" || c === "<") {
      stack.push(c);
    } else {
      if (partner[c] !== stack.pop()) {
        return c;
      }
    }
  }
  return null;
}

function completion(s: string): string {
  const stack = [];
  for (let i = 0; i < s.length; i++) {
    let c = s.charAt(i);
    if (c === "{" || c === "(" || c === "[" || c === "<") {
      stack.push(c);
    } else {
      stack.pop();
    }
  }
  let res = "";
  while (stack.length > 0) {
    res += closing[stack.pop()!];
  }
  return res;
}

function computeScore(s: string): number {
  const points: { [key: string]: number } = {
    ")": 1,
    "]": 2,
    "}": 3,
    ">": 4,
  };
  let score = 0;
  for (let i = 0; i < s.length; i++) {
    score = score * 5 + points[s.charAt(i)]!;
  }
  return score;
}
