defmodule AOCDay18 do

  def convert_hex(hex) do
    << "(#", val::binary-size(6), ")" >> = hex

    {num, dir} = String.split_at(val, 5)
    {
      case dir do
        "0" -> "R"
        "1" -> "D"
        "2" -> "L"
        "3" -> "U"
      end, 
      String.to_integer(num, 16)}
  end

  def parse(file, :part1) do
    File.read!(file)
    |> String.split("\n", trim: true)
    |> Enum.map(fn s -> 
      String.split(s)
      |> Enum.take(2)
      end)
    |> Enum.map(fn [d, l] -> {d, String.to_integer(l)} end) 
  end
  
  def parse(file, :part2) do
    File.read!(file)
    |> String.split("\n", trim: true)
    |> Enum.map(fn s -> 
      String.split(s)
      |> then(fn [_, _, hex] -> hex end)
      |> convert_hex()
      end)
  end
  

  def shoelace([_]), do: 0

  def shoelace([{x1, y1}, {x2, y2} | tail]) do
    (y1 + y2) * (x1 - x2) + shoelace([{x2, y2} | tail])
  end

  def solve(file, part) do
    edges = parse(file, part)

    circ = Enum.map(edges, fn {_, l} -> l end)
      |> Enum.sum()
    area = Enum.scan(edges, {0,0}, fn {d, l}, {x,y} -> 
        case d do
          "R" -> {x + l, y}
          "L" -> {x - l, y}
          "U" -> {x, y + l}
          "D" -> {x, y - l}
        end
      end)
      |> then(fn l -> [{0,0} | l] end)
      |> shoelace()
      |> abs()
      |> div(2)
    circ + area - div(circ, 2) + 1
  end

  

end


