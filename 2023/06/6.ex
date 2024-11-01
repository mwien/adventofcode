defmodule AOCDay6 do
  def preprocess(input, 1) do
    input
  end

  def preprocess(input, 2) do
    String.replace(input, " ", "")
  end

  def parse(file, task) do
    ["Time:" <> times, "Distance:" <> distances] =
      File.read!(file)
      |> preprocess(task)
      |> String.trim()
      |> String.split("\n")

    Enum.zip(
      Enum.map(times |> String.trim() |> String.split(), &String.to_integer(&1)),
      Enum.map(distances |> String.trim() |> String.split(), &String.to_integer(&1))
    )
  end

  def check_record(chargetime, time, distance) do
    if chargetime * (time - chargetime) > distance, do: 1, else: 0
  end

  def winning_speeds({time, distance}) do
    Enum.reduce(0..time, 0, &(&2 + check_record(&1, time, distance)))
  end

  def part1(file) do
    parse(file, 1)
    |> Enum.reduce(1, &(&2 * winning_speeds(&1)))
  end

  def correction(solution, time, distance) do
    if solution * (time - solution) <= distance, do: 1, else: 0
  end

  def winning_speeds_fast([{time, distance}]) do
    first = round(time / 2 - (time ** 2 / 4 - distance) ** 0.5)
    first = first + correction(first, time, distance)
    second = round(time / 2 + (time ** 2 / 4 - distance) ** 0.5)
    second = second - correction(second, time, distance)
    second - first + 1
  end

  def part2(file) do
    parse(file, 2)
    |> winning_speeds_fast()
  end
end
