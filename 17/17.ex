defmodule AOCDay17 do

  def to_index_map(s) do
    String.trim(s)
    |> String.split("\n")
    |> Enum.with_index()
    |> Enum.reduce(%{}, fn {ln, i}, mp -> 
      Map.merge(mp, 
        String.to_charlist(ln)
        |> Enum.with_index()
        |> Map.new(fn {c, j} -> {{j,i}, c-?0} end)
      )
    end)
  end
  
  def get_min(dists) do
    Map.to_list(dists)
    |> Enum.min_by(fn {_, {v, vis}} -> if vis, do: 10**9, else: v end)
  end

  def get_between({i, j}, {ni, nj}) do
    cond do
      i == ni -> for k <- j..nj, do: {i, k}
      j == nj -> for k <- i..ni, do: {k, j} 
    end
    |> Enum.filter(&(&1 != {i,j}))
  end

  def postprocess(ls, grid, {i,j,_}) do
    List.flatten(ls)
    |> Enum.filter(&Map.has_key?(grid, &1))
    |> Enum.map(fn {ni, nj} -> 
      {ni, nj, get_between({i, j}, {ni, nj})
      |> Enum.map(&(Map.get(grid, &1)))
      |> Enum.sum()}
    end)
  end

  def neighbors(grid, {i, j, l}, :part1) do
    cond do
      l == 0 -> for delta <- 1..3, do: [{i, j-delta}, {i, j+delta}]
      l == 1 -> for delta <- 1..3, do: [{i-delta, j}, {i+delta, j}]
    end
    |> postprocess(grid, {i, j, l})
  end

  def neighbors(grid, {i, j, l}, :part2) do
    cond do
      l == 0 -> for delta <- 4..10, do: [{i, j-delta}, {i, j+delta}]
      l == 1 -> for delta <- 4..10, do: [{i-delta, j}, {i+delta, j}]
    end
    |> postprocess(grid, {i, j, l})
  end

  def relaxation(grid, dists, {i,j,l}, d, part) do
    neighbors(grid, {i,j,l}, part)
    |> Enum.reduce(dists, 
      fn {ni, nj, cost}, ndists -> 
        {odist, ovis} = Map.get(ndists, {ni, nj, 1-l}, {10**9, false})
        ndist = d + cost
        if ndist < odist, do: Map.put(ndists, {ni, nj, 1-l}, {ndist, ovis}), else: ndists
      end)
  end

  def dijkstra(dists, grid, part) do
    case get_min(dists) do
      {_, {_, true}} -> dists
      {pos, {d, _}} -> 
        relaxation(grid, Map.put(dists, pos, {d, true}), pos, d, part)
        |> dijkstra(grid, part)
    end
  end

  def solve(file, part) do
    grid = File.read!(file)
    |> to_index_map()
    {n, m} = Enum.max(Map.keys(grid))

    dists = Enum.map([0,1], &({ {0,0,&1}, {0,false} }))
    |> Map.new()
    |> dijkstra(grid, part)

    min(Map.get(dists, {n,m,0}, 10**9), Map.get(dists, {n,m,1}, 10**9))
  end
  
end

