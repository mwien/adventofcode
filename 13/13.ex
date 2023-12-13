defmodule AOCDay13 do

  def to_index_map(s) do
    String.split(s, "\n", trim: true)
    |> Enum.with_index()
    |> Enum.reduce(%{}, fn {ln, i}, mp -> 
      Map.merge(mp, 
        String.to_charlist(ln)
        |> Enum.with_index()
        |> Map.new(fn {c, j} -> {{j,i}, c} end)
      )
    end)
  end

  def check_line(i, j, m, grid, :horizontal) do
    Enum.all?(0..m, &(Map.get(grid, {i, &1}) == Map.get(grid, {j, &1})))
  end
  
  def check_line(i, j, m, grid, :vertical) do
    Enum.all?(0..m, &(Map.get(grid, {&1, i}) == Map.get(grid, {&1, j})))
  end

  def check_reflection(i, n, m, grid, dir) do
    cond do
      Enum.all?(0..min(i, n-i-1), &check_line(i - &1, &1 + 1 + i, m, grid, dir)) -> i + 1
      true -> 0
    end
  end

  def find_reflection(grid) do
    {n, m} = Enum.max(Map.keys(grid))
    horizontal = Enum.map(0..n, &check_reflection(&1, n, m, grid, :horizontal))
    vertical = Enum.map(0..m, &check_reflection(&1, m, n, grid, :vertical))
    |> Enum.map(&(&1 * 100))
    horizontal ++ vertical
  end

  def part1(file) do
    File.read!(file)
    |> String.split("\n\n", trim: true)
    |> Enum.map(&to_index_map(&1))
    |> Enum.map(&find_reflection(&1))
    |> List.flatten()
    |> Enum.sum()
  end

  def smudge_reflection(grid) do
    old = find_reflection(grid)
      |> Enum.sum()
    Map.keys(grid)
    |> Enum.map(fn pos -> 
      c = grid[pos]
      find_reflection(Map.put(grid, pos, (if c == ?#, do: ?., else: ?#)))
    end)
    |> List.flatten()
    |> Enum.filter(&(&1 != old))
    |> Enum.uniq()
    |> Enum.sum()
  end

  def part2(file) do
    File.read!(file)
    |> String.split("\n\n", trim: true)
    |> Enum.map(&to_index_map(&1))
    |> Enum.map(&smudge_reflection(&1))
    |> Enum.sum()
  end
end

