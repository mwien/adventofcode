import sys
from enum import Enum
from typing import Dict, Self


class EyeColor(Enum):
    AMBER = "amb"
    BLUE = "blu"
    BROWN = "brn"
    GREY = "gry"
    GREEN = "grn"
    HAZEL = "hzl"
    OTHER = "oth"


class HeightUnit(Enum):
    CM = "cm"
    INCH = "in"


class Height:
    def __init__(self, value: float, unit: HeightUnit):
        self.value = value
        self.unit = unit

    def __repr__(self):
        return f"{self.value} {self.unit.value}"


class Passport:
    def __init__(
        self,
        byr: int,
        iyr: int,
        eyr: int,
        hgt: Height,
        hcl: str,
        ecl: EyeColor,
        pid: str,
        cid: str,
    ):
        self.byr = byr
        self.iyr = iyr
        self.eyr = eyr
        self.hgt = hgt
        self.hcl = hcl
        self.ecl = ecl
        self.pid = pid
        self.cid = cid

    @staticmethod
    def parse_and_validate(field: str, val: str):
        if field == "byr":
            byr = int(val)
            if byr < 1920 or byr > 2002:
                raise ValueError(f"birth year '{byr}' unplausible")
            return byr
        elif field == "iyr":
            iyr = int(val)
            if iyr < 2010 or iyr > 2020:
                raise ValueError(f"issue year '{iyr}' incorrect")
            return iyr
        elif field == "eyr":
            eyr = int(val)
            if eyr < 2020 or eyr > 2030:
                raise ValueError(f"expiration date '{eyr}' incorrect")
            return eyr
        elif field == "hgt":
            height = int(val[:-2])
            unit = val[-2:]
            if unit == "cm":
                if height < 150 or height > 193:
                    raise ValueError(f"height '{height}' cm unplausible")
                return Height(height, HeightUnit.CM)
            elif unit == "in":
                if height < 59 or height > 76:
                    raise ValueError(f"height '{height}' in unplausible")
                return Height(height, HeightUnit.INCH)
            else:
                raise ValueError(f"unknown unit '{unit}'")
        elif field == "hcl":
            if val[0] != "#":
                raise ValueError(f"hair color starts with '{val[0]}' instead of '#'")
            int(val[1:], 16)
            return val
        elif field == "ecl":
            return EyeColor(val)
        elif field == "pid":
            if len(val) != 9 or not val.isdigit():
                raise ValueError(f"pid '{val}' is not a nine digit number")
            return val
        elif field == "cid":
            return val
        else:
            # unknown field is ok
            return val

    @classmethod
    def from_fields(cls, fields: Dict[str, str], validate=True) -> Self:
        required_fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
        values = []
        for r in required_fields:
            if r not in fields:
                raise ValueError(f"Required field {r} is missing")
            if validate:
                values.append(Passport.parse_and_validate(r, fields[r]))
            else:
                values.append(fields[r])
        if "cid" in fields:
            values.append(Passport.parse_and_validate("cid", fields["cid"]))
        else:
            values.append("")
        return cls(*values)


cnt1 = 0
cnt2 = 0
passport = {}
for line in sys.stdin:
    line = line.strip()
    if not line:
        try:
            Passport.from_fields(passport, validate=False)
            cnt1 += 1
        except ValueError as e:
            print(e)
        try:
            Passport.from_fields(passport)
            cnt2 += 1
        except ValueError as e:
            print(e)
        passport = {}
    else:
        for field in line.split():
            k, v = field.split(":")
            passport[k] = v

print(cnt1)
print(cnt2)
