defmodule AOCDay16 do
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

  def neighbors(grid, {x, y} = pos, {dx, dy} = dir) do
    case grid[pos] do
      ?. ->
        [{{x + dx, y + dy}, dir}]

      ?| ->
        case dir do
          {0, _} -> [{{x, y + dy}, dir}]
          {_, 0} -> [{{x, y - dy}, {0, -1}}, {{x, y + dy}, {0, 1}}]
        end

      ?- ->
        case dir do
          {0, _} -> [{{x - dx, y}, {-1, 0}}, {{x + dx, y}, {1, 0}}]
          {_, 0} -> [{{x + dx, y}, dir}]
        end

      ?/ ->
        [{{x - dy, y - dx}, {-dy, -dx}}]

      ?\\ ->
        [{{x + dy, y + dx}, {dy, dx}}]
    end
    |> Enum.filter(fn {pos, _} -> Map.has_key?(grid, pos) end)
  end

  def dfs(grid, pos, dir, vis) do
    Enum.reduce(neighbors(grid, pos, dir), MapSet.put(vis, {pos, dir}), fn {npos, ndir}, nvis ->
      if MapSet.member?(nvis, {npos, ndir}), do: nvis, else: dfs(grid, npos, ndir, nvis)
    end)
  end

  def count_energized(grid, {pos, dir} = _init) do
    dfs(grid, pos, dir, MapSet.new())
    |> Enum.map(fn {pos, _} -> pos end)
    |> MapSet.new()
    |> MapSet.size()
  end

  def part1(file) do
    File.read!(file)
    |> to_index_map()
    |> count_energized({{0, 0}, {1, 0}})
  end

  def get_init_pos({n, m}) do
    (Enum.map(0..n, &[{{&1, 0}, {0, 1}}, {{&1, m}, {0, -1}}]) ++
       Enum.map(0..m, &[{{0, &1}, {1, 0}}, {{n, &1}, {-1, 0}}]))
    |> List.flatten()
  end

  def part2(file) do
    grid =
      File.read!(file)
      |> to_index_map()

    Enum.max(Map.keys(grid))
    |> get_init_pos()
    |> Enum.map(&count_energized(grid, &1))
    |> Enum.max()
  end
end
