defmodule AOCDay11 do
  def to_index_map(s) do
    String.trim(s)
    |> String.split("\n")
    |> Enum.with_index()
    |> Enum.reduce(%{}, fn {ln, i}, mp ->
      Map.merge(
        mp,
        String.to_charlist(ln)
        |> Enum.with_index()
        |> Map.new(fn {c, j} -> {{j, i}, c} end)
      )
    end)
  end

  def is_empty(i, dim, grid, dir, part) do
    Enum.map(0..dim, &if(dir == :row, do: {i, &1}, else: {&1, i}))
    |> Enum.count(&(grid[&1] == ?.))
    |> div(dim + 1)
    |> Kernel.*(if part == 2, do: 1_000_000, else: 2)
    |> max(1)
  end

  def presum(dim1, dim2, dir, grid, part) do
    List.duplicate(1, dim1 + 1)
    |> Enum.with_index()
    |> Enum.map(fn {x, i} ->
      {i, x * is_empty(i, dim2, grid, dir, part)}
    end)
    |> Enum.scan({-1, 0}, fn {j, x}, {_, y} -> {j, x + y} end)
    |> Map.new()
  end

  def solve(file, part) do
    grid =
      File.read!(file)
      |> to_index_map()

    {n, m} =
      grid
      |> Map.keys()
      |> Enum.max()

    prows = presum(n, m, :row, grid, part)
    pcols = presum(m, n, :col, grid, part)

    pos =
      Map.keys(grid)
      |> Enum.filter(&(grid[&1] == ?#))

    for {x1, y1} <- pos, {x2, y2} <- pos do
      cond do
        {x1, y1} == {x2, y2} -> 0
        true -> abs(prows[x2] - prows[x1]) + abs(pcols[y2] - pcols[y1])
      end
    end
    |> Enum.sum()
    |> div(2)
  end
end
