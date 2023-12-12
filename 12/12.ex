defmodule AOCDay12 do
  
  def all_confs(springs) do
    cond do
      !String.contains?(springs, "?") -> [springs]
      true -> all_confs(String.replace(springs, "?", "#", global: false)) ++ all_confs(String.replace(springs, "?", ".", global: false))
    end
  end

  def check_conf(springs, seq) do
    l = String.split(springs, ".", trim: true)
    |> Enum.map(&String.length(&1))
    |> Enum.join(",")
    l == seq
  end

  def count_row([springs, seq]) do
    all_confs(springs)
    |> Enum.filter(&check_conf(&1, seq))
    |> length()
  end

  def part1(file) do
    File.read!(file)
    |> String.split("\n", trim: true)
    |> Enum.map(&count_row(String.split(&1)))
    |> Enum.sum()
  end

end

