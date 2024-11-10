import * as fs from "fs";

//const input = fs.readFileSync("sample.in").toString();
const input = fs.readFileSync("main.in").toString();

part1(input);
part2(input);

function part1(input) {
  input = input.split("\n\n");
  const numbers = input[0].split(",").map(Number);
  const boards = input.slice(1).map(parseBoard);
  for (const x of numbers) {
    for (let i = 0; i < boards.length; i++) {
      const positions = findNumber(boards[i], x);
      if (!positions.length) {
        continue;
      }
      for (const [row, col] of positions) {
        boards[i][row][col] = -1;
      }
      for (const [row, col] of positions) {
        if (winningRow(boards[i], row) || winningCol(boards[i], col)) {
          console.log(x * boardSum(boards[i]));
          return;
        }
      }
    }
  }
}

function part2(input) {
  input = input.split("\n\n");
  const numbers = input[0].split(",").map(Number);
  const boards = input.slice(1).map(parseBoard);
  const finishedBoards = new Array(boards.length).fill(false);
  let lastValue = 0;
  for (const x of numbers) {
    for (let i = 0; i < boards.length; i++) {
      if (finishedBoards[i]) {
        continue;
      }
      const positions = findNumber(boards[i], x);
      if (!positions.length) {
        continue;
      }
      for (const [row, col] of positions) {
        boards[i][row][col] = -1;
      }
      for (const [row, col] of positions) {
        if (winningRow(boards[i], row) || winningCol(boards[i], col)) {
          finishedBoards[i] = true;
          lastValue = x * boardSum(boards[i]);
        }
      }
    }
  }
  console.log(lastValue);
}

function parseBoard(s) {
  return s.split("\n").map(parseRow);
}

function parseRow(s) {
  return s.trim().split(/\s+/).map(Number);
}

function findNumber(board, x) {
  const positions = [];
  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      if (board[i][j] == x) {
        positions.push([i, j]);
      }
    }
  }
  return positions;
}

function winningRow(board, row) {
  for (let i = 0; i < board[row].length; i++) {
    if (board[row][i] != -1) {
      return false;
    }
  }
  return true;
}

function winningCol(board, col) {
  for (let i = 0; i < board.length; i++) {
    if (board[i][col] != -1) {
      return false;
    }
  }
  return true;
}

function boardSum(board) {
  let sum = 0;
  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      if (board[i][j] != -1) {
        sum += board[i][j];
      }
    }
  }
  return sum;
}
