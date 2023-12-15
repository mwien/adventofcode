defmodule AOCDay15 do

  def reindeer_hash(s) do
    String.to_charlist(s)
    |> Enum.reduce(0, fn c, curr -> 
        rem((curr + c) * 17, 256)
      end)
  end

  def part1(file) do
    File.read!(file)
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&reindeer_hash(&1))
    |> Enum.sum()
  end

  def perform_operation(s) do
    
  end

  def part2(file) do
    File.read!(file)
    |> String.trim()
    |> String.split(",")
    |> Enum.reduce(%{}, &perform_operation(&1))
    |> focusing_power()
  end

end

