import sys
from enum import Enum
from typing import Self


class Command(Enum):
    N = "N"
    S = "S"
    E = "E"
    W = "W"
    L = "L"
    R = "R"
    F = "F"


class Instruction:
    def __init__(self, command: Command, value: int):
        self.command = command
        self.value = value

    @classmethod
    def from_str(cls, s: str) -> Self:
        command = Command(s[0])
        value = int(s[1:])
        return cls(command, value)


def perform_instruction(pos: list[int], ins: Instruction):
    match ins.command:
        case Command.N:
            pos[1] += ins.value
        case Command.S:
            pos[1] -= ins.value
        case Command.E:
            pos[0] += ins.value
        case Command.W:
            pos[0] -= ins.value
        case Command.L:
            pos[2] += ins.value
            pos[2] %= 360
        case Command.R:
            pos[2] -= ins.value
            pos[2] %= 360
        case Command.F:
            if pos[2] == 0:
                perform_instruction(pos, Instruction(Command.E, ins.value))
            elif pos[2] == 90:
                perform_instruction(pos, Instruction(Command.N, ins.value))
            elif pos[2] == 180:
                perform_instruction(pos, Instruction(Command.W, ins.value))
            elif pos[2] == 270:
                perform_instruction(pos, Instruction(Command.S, ins.value))


def perform_instruction2(ship: list[int], wp: list[int], ins: Instruction):
    match ins.command:
        case Command.N:
            wp[1] += ins.value
        case Command.S:
            wp[1] -= ins.value
        case Command.E:
            wp[0] += ins.value
        case Command.W:
            wp[0] -= ins.value
        case Command.L:
            # TODO
            if ins.value == 90:
                wp[0], wp[1] = wp[1], wp[0]
            elif ins.value == 180:
                wp[0], wp[1] = -wp[0], -wp[1]
            elif ins.value == 270:
                # TODO
                pass
        case Command.R:
            perform_instruction2(ship, wp, Instruction(Command.L, -ins.value % 360))
        case Command.F:
            pass


def navigate(instructions: list[Instruction]) -> list[int]:
    pos = [0, 0, 0]
    for instruction in instructions:
        perform_instruction(pos, instruction)
    return pos


instructions = [Instruction.from_str(line) for line in sys.stdin]
pos = navigate(instructions)
print(abs(pos[0]) + abs(pos[1]))
