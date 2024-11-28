import * as fs from "fs";

//const input = fs.readFileSync("sample.in").toString().trim();
const input = fs.readFileSync("main.in").toString().trim();

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
  console.log(res);
}

function reconstructLine(pattern, digits) {
  const mapping = findMapping(pattern);
  let output = 0;
  for (const digit of digits) {
    output *= 10;
    output += decodeDigit(digit, mapping);
  }
  return output;
}

function findMapping(pattern) {
  // untangle by hand
  const [one] = getWithLength(pattern, 2);
  const [seven] = getWithLength(pattern, 3);
  const [a] = setDiff(seven, one);
  const [four] = getWithLength(pattern, 4);
  const zerosixnine = getWithLength(pattern, 6);
  const cde = notInAll(zerosixnine);
  const [c] = setIntersection(cde, one);
  const [d] = setIntersection(setDiff(cde, c), four);
  const [e] = setDiff(cde, c + d);
  const [f] = setDiff(one, c);
  const [b] = setDiff(four, c + d + f);
  const [eight] = getWithLength(pattern, 7);
  const [g] = setDiff(eight, a + b + c + d + e + f);
  // assign to map for decoding
  const mapping = new Map();
  mapping.set(a, "a");
  mapping.set(b, "b");
  mapping.set(c, "c");
  mapping.set(d, "d");
  mapping.set(e, "e");
  mapping.set(f, "f");
  mapping.set(g, "g");
  return mapping;
}

function decodeDigit(digit, mapping) {
  const decodedDigit = sortString([...digit].map((c) => mapping.get(c)));
  switch (decodedDigit) {
    case "abcefg":
      return 0;
    case "cf":
      return 1;
    case "acdeg":
      return 2;
    case "acdfg":
      return 3;
    case "bcdf":
      return 4;
    case "abdfg":
      return 5;
    case "abdefg":
      return 6;
    case "acf":
      return 7;
    case "abcdefg":
      return 8;
    case "abcdfg":
      return 9;
  }
}

function getWithLength(pattern, l) {
  return pattern.filter((digit) => digit.length == l);
}

function setIntersection(s1, s2) {
  const res = [];
  for (const c of s1) {
    if (s2.includes(c)) {
      res.push(c);
    }
  }
  return res;
}

function notInAll(arr) {
  const res = [];
  for (let i = 0; i < arr.length; i++) {
    for (const c of arr[i]) {
      let isIn = 1;
      for (let j = 0; j < arr.length; j++) {
        if (i == j) {
          continue;
        }
        if (arr[j].includes(c)) {
          isIn += 1;
        }
      }
      if (isIn < 3 && !res.includes(c)) {
        res.push(c);
      }
    }
  }
  return res;
}

function setDiff(s1, s2) {
  const res = [];
  for (const c of s1) {
    if (!s2.includes(c)) {
      res.push(c);
    }
  }
  return res;
}

function sortString(s) {
  return [...s].sort().join("");
}
