defmodule AOCDay23 do
  
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

  def neighbors(grid, {x, y}) do
    (for dx <- [-1,0,1], dy <- [-1,0,1], rem(dx+dy, 2) != 0, do: {x + dx, y + dy})
    |> Enum.filter(&Map.has_key?(grid, &1))
    |> Enum.filter(fn {nx, ny} = pos -> 
      case Map.get(grid, pos) do
        ?# -> false 
        ?> -> x + 1 == nx 
        ?< -> x - 1 == nx 
        ?v -> y + 1 == ny 
        ?^ -> y - 1 == ny 
        ?. -> true
      end
    end)
  end

  def top_sort(grid, u, vis, to) do
    neighbors(grid, u)
    |> Enum.filter(&(!MapSet.member?(vis, &1)))
    |> Enum.reduce({MapSet.put(vis, u), to}, fn v, {vis, to} -> 
      top_sort(grid, v, vis, to)
    end)
    |> then(fn {vis, to} -> {vis, [u | to]} end)
  end

  def longest_path(dists, [], _), do: dists
  def longest_path(dists, [head | tail], grid) do
    neighbors(grid, head)
    |> Enum.reduce(dists, fn ngh, dists -> 
      Map.update(dists, ngh, dists[head]+1, &max(&1, dists[head]+1))
    end)
    |> longest_path(tail, grid)
  end

  def part1(file) do
    grid = File.read!(file)
    |> to_index_map()

    {_, to} = top_sort(grid, {1,0}, MapSet.new(), [])

    longest_path(%{{1,0} => 0}, to, grid)
    |> Enum.to_list()
    |> Enum.max_by(fn {{_, y}, _} -> y end)
  end
  
  def allneighbors(grid, {x, y}) do
    (for dx <- [-1,0,1], dy <- [-1,0,1], rem(dx+dy, 2) != 0, do: {x + dx, y + dy})
    |> Enum.filter(&Map.has_key?(grid, &1))
    |> Enum.filter(fn pos -> 
      case Map.get(grid, pos) do
        ?# -> false 
        _ -> true
      end
    end)
  end

  def dfs(grid, u, vis, {n, m}) do
    if u == {n-1, m} do
      MapSet.size(vis)
    else 
      allneighbors(grid, u)
      |> Enum.filter(&(!MapSet.member?(vis, &1)))
      |> Enum.reduce(0, fn ngh, mx -> 
        max(mx, dfs(grid, ngh, MapSet.put(vis, ngh), {n, m})) 
      end)
    end
  end

  # too slow like this -> need to compress the input to a graph with edge weights
  def part2(file) do
    grid = File.read!(file)
    |> to_index_map()
   
    {n, m} = Map.keys(grid) |> 
      Enum.max()

    dfs(grid, {1,0}, MapSet.new([{1,0}]), {n,m})
    |> Kernel.-(1)
  end

end
