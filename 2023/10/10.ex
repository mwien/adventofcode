defmodule AOCDay10 do
  def connections(c) do
    case c do
      ?| -> [{0, -1}, {0, 1}]
      ?- -> [{-1, 0}, {1, 0}]
      ?L -> [{0, -1}, {1, 0}]
      ?J -> [{-1, 0}, {0, -1}]
      ?7 -> [{0, 1}, {-1, 0}]
      ?F -> [{0, 1}, {1, 0}]
      _ -> []
    end
  end

  def neighbors({i, j}) do
    for x <- [-1, 0, 1], y <- [-1, 0, 1], rem(x + y, 2) != 0, do: {i + x, j + y}
  end

  def connection_fromto({a1, b1}, {a2, b2}, c) do
    connections(c)
    |> Enum.map(fn {x, y} -> {a1 + x, b1 + y} end)
    |> Enum.member?({a2, b2})
  end

  def connection_between(grid, p1, p2) do
    cond do
      !connection_fromto(p1, p2, grid[p1]) -> false
      !connection_fromto(p2, p1, grid[p2]) -> false
      true -> true
    end
  end

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

  def find_tile(grid, pos) do
    [?|, ?-, ?L, ?J, ?7, ?F]
    |> Enum.filter(fn c ->
      neighbors(pos)
      |> Enum.filter(&connection_between(Map.put(grid, pos, c), pos, &1))
      |> length() == 2
    end)
    |> List.first()
  end

  def traverse(circle = [last | tail], grid, start) do
    next =
      neighbors(last)
      |> Enum.filter(&Map.has_key?(grid, &1))
      |> Enum.filter(&connection_between(grid, last, &1))
      |> Enum.filter(&(List.first(tail) != &1))
      |> List.first()

    cond do
      next == start -> circle
      true -> traverse([next | circle], grid, start)
    end
  end

  def get_cycle(file) do
    grid =
      File.read!(file)
      |> to_index_map()

    pos =
      grid
      |> Enum.find(fn {_, v} -> v == ?S end)
      |> elem(0)

    c = find_tile(grid, pos)
    grid = Map.put(grid, pos, c)

    traverse([pos], grid, pos)
  end

  def part1(file) do
    get_cycle(file)
    |> length()
    |> div(2)
  end

  def shoelace([_]), do: 0

  def shoelace([{x1, y1}, {x2, y2} | tail]) do
    (y1 + y2) * (x1 - x2) + shoelace([{x2, y2} | tail])
  end

  def picks(a, b) do
    a - div(b, 2) + 1
  end

  def part2(file) do
    cycle =
      get_cycle(file)
      |> then(&[List.last(&1) | &1])

    picks(div(shoelace(cycle), 2), length(cycle) - 1)
  end
end
