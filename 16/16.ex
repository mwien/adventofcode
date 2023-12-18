defmodule AOCDay16 do

  def to_index_map(s) do
    String.trim(s)
    |> String.split("\n")
    |> Enum.with_index()
    |> Enum.reduce(%{}, fn {ln, i}, mp -> 
      Map.merge(mp, 
        String.to_charlist(ln)
        |> Enum.with_index()
        |> Map.new(fn {c, j} -> {{j,i}, c} end)
      )
    end)
  end

  def neighbors(grid, {x, y} = pos, {dx, dy} = dir) do
    case grid[pos] do
      ?. -> [ {{x+dx, y+dy}, dir} ]
      ?| -> case dir do
        {0, _} -> [ {{x, y+dy}, dir} ]
        {_, 0} -> [ {{x, y-dy}, {0, -1}}, {{x, y+dy}, {0, 1}} ]
      end
      ?- -> case dir do
        {0, _} -> [ {{x-dx, y}, {-1, 0}}, {{x+dx, y}, {1, 0}} ]
        {_, 0} -> [ {{x+dx, y}, dir} ]
      end
      ?/ -> [ {{x-dy, y-dx}, {-dy, -dx}} ]
      ?\\ -> [ {{x+dy, y+dx}, {dy, dx}} ]
    end
    |> Enum.filter(fn {pos, _} -> Map.has_key?(grid, pos) end)
  end

  def dfs(grid, pos, dir, vis) do
    Enum.reduce(neighbors(grid, pos, dir), MapSet.put(vis, pos), fn {npos, ndir}, nvis -> 
      dfs(grid, npos, ndir, nvis) end) 
    # TODO: check vis
  end

  def part1(file) do
    grid = File.read!(file)
    |> to_index_map()

    dfs(grid, {0,0}, {1,0}, MapSet.new()) 
    |> MapSet.size()
  end

end

