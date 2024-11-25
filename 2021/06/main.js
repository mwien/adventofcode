import * as fs from "fs";

//const input = fs.readFileSync("sample.in").toString();
const input = fs.readFileSync("main.in").toString();

part1(input);
part2(input);

function part1(input) {
  const population = parsePopulation(input);
  const newPopulation = simulatePopulation(population, 80);
  console.log(sum(newPopulation));
}

function part2(input) {
  const population = parsePopulation(input);
  const newPopulation = simulatePopulation(population, 256);
  console.log(sum(newPopulation));
}

function simulatePopulation(population, days) {
  for (let i = 0; i < days; i++) {
    const newPopulation = new Array(9).fill(0);
    newPopulation[6] = population[0];
    newPopulation[8] = population[0];
    for (let j = 1; j <= 8; j++) {
      newPopulation[j - 1] += population[j];
    }
    population = newPopulation;
  }
  return population;
}

function parsePopulation(input) {
  const initialTimers = input.split(",").map(Number);
  const population = new Array(9).fill(0);
  for (let i = 0; i < initialTimers.length; i++) {
    population[initialTimers[i]] += 1;
  }
  return population;
}

function sum(arr) {
  let sum = 0;
  for (let i = 0; i < arr.length; i++) {
    sum += arr[i];
  }
  return sum;
}
