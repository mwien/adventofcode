defmodule AOCDay14 do

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

  def accumulate_load(grid, i, j, {sum, open}, m) do
    case grid[{i,j}] do
      ?O -> {sum, open+1}
      ?. -> {sum, open}
      ?# -> {sum + (m-j+1) * open - div(open * (open + 1), 2), 0}
    end
  end

  def load_per_column(grid, i, m) do
    {sum, open} = Enum.reduce(m..0, {0,0}, &accumulate_load(grid, i, &1, &2, m))
    sum + (m+1) * open - div(open * (open - 1), 2)
  end

  def part1(file) do
    grid = File.read!(file)
    |> to_index_map()

    {n, m} = Map.keys(grid) 
    |> Enum.max()

    Enum.map(0..n, &load_per_column(grid, &1, m))
    |> IO.inspect()
    |> Enum.sum()
  end

end

