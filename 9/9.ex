defmodule AOCDay8 do

  def parse(file) do
    File.read!(file)
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(fn sequence -> 
      String.split(sequence)
      |> Enum.map(&String.to_integer(&1))
    end) 
  end

  def differences([_]) do
    []
  end

  def differences([a, b | tail]) do
    [b - a | differences([b | tail])]
  end

  def predict_next(sequence) do
    if Enum.all?(sequence, &(&1 == 0)) do
      0
    else 
      List.last(sequence) + predict_next(differences(sequence))
    end
  end

  def predict_before(sequence) do
    if Enum.all?(sequence, &(&1 == 0)) do
      0
    else 
      List.first(sequence) - predict_before(differences(sequence))
    end  
  end

  def part1(file) do
    parse(file)
    |> Enum.map(&predict_next(&1))
    |> Enum.sum()
  end

  def part2(file) do
    parse(file)
    |> Enum.map(&predict_before(&1))
    |> Enum.sum()
  end
end

