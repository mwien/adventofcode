from enum import Enum
import sys
from typing import Self
import copy


class CycleException(Exception):
    pass


class Operation(Enum):
    NOP = "nop"
    ACC = "acc"
    JMP = "jmp"


class Instruction:
    def __init__(self, op: Operation, arg: int) -> None:
        self.op = op
        self.arg = arg

    def __repr__(self) -> str:
        return f"Instruction({self.op}, {self.arg})"

    @classmethod
    def parse(cls, s: str) -> Self:
        tokens = s.split()
        op = Operation(tokens[0])
        arg = int(tokens[1])
        return cls(op, arg)


class Bytecode:
    def __init__(self, instructions: list[Instruction]) -> None:
        self.instructions = instructions
        self.index = 0
        self.acc = 0

    def __repr__(self) -> str:
        return f"Bytecode({self.instructions}, {self.index}, {self.acc})"

    def step_forward(self) -> None:
        """Performs the op at index and returns the new index."""
        instruction = self.instructions[self.index]
        if instruction.op == Operation.NOP:
            self.index += 1
        elif instruction.op == Operation.ACC:
            self.index += 1
            self.acc += instruction.arg
        else:
            self.index += instruction.arg

    def toggled(self, idx: int) -> Self:
        toggled_bytecode = copy.deepcopy(self)
        if self.instructions[idx].op == Operation.NOP:
            toggled_bytecode.instructions[idx].op = Operation.JMP
        elif self.instructions[idx].op == Operation.JMP:
            toggled_bytecode.instructions[idx].op = Operation.NOP
        else:
            pass
        return toggled_bytecode

    def simulate(self) -> int:
        indices = set()
        last_acc = 0
        while True:
            self.step_forward()
            index = self.index
            acc = self.acc
            if index in indices:
                self.acc = 0
                self.index = 0
                raise CycleException(self, last_acc)
            if index >= len(self.instructions):
                self.acc = 0
                self.index = 0
                return acc
            indices.add(index)
            last_acc = acc


def solve():
    instructions = []
    for line in sys.stdin:
        instructions.append(Instruction.parse(line))

    self = Bytecode(instructions)

    try:
        self.simulate()
    except CycleException as c:
        print(c.args[1])

    for i, _ in enumerate(self.instructions):
        toggled_bytecode = self.toggled(i)
        try:
            res = toggled_bytecode.simulate()
        except CycleException:
            pass
        else:
            print(res)


solve()
