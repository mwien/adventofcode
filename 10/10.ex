defmodule AOCDay10 do

  def connections(c) do
    case c do
      ?| -> [{0,-1}, {0,1}]
      ?- -> [{-1,0}, {1,0}]
      ?L -> [{0,-1}, {1,0}]
      ?J -> [{-1,0}, {0,-1}]
      ?7 -> [{0,1}, {-1,0}]
      ?F -> [{0,1}, {1,0}]
      _ -> []
    end
  end

  def neighbors({i,j}) do
    for x <- [-1, 0, 1], y <- [-1, 0, 1], rem(x+y, 2) != 0, do: {i+x, j+y}
  end

  def connection_between({{a1, b1}, c1}, {{a2, b2}, c2}) do
    if Enum.member?(Enum.map(connections(c1), fn {x, y} -> {a1+x,b1+y} end),  {a2, b2}) and Enum.member?(Enum.map(connections(c2), fn {x, y} -> {a2+x,b2+y} end),  {a1, b1}) do
      true
    else 
      false
    end
  end

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

  def find_tile(grid, {i,j}) do
    [?|, ?-, ?L, ?J, ?7, ?F]
    |> Enum.filter(fn c -> 
      length(Enum.filter(neighbors({i,j}), fn {a,b} -> 
        connection_between({{i, j}, c}, {{a, b}, Map.get(grid, {a,b}, ?.)})
      end)) == 2 
    end)
    |> List.first()
  end

  def get_distances(grid, queue, dists) do
    case :queue.out(queue) do
      {{:value, {{i, j}, d}}, queue} -> 
        new_neighbors = Enum.filter(neighbors({i,j}), fn pos -> !Map.has_key?(dists, pos) and Map.has_key?(grid, pos) and connection_between({pos, Map.fetch!(grid, pos)}, {{i, j}, Map.fetch!(grid, {i,j})}) end)
        new_queue = Enum.reduce(new_neighbors, queue, &:queue.in({&1, d+1}, &2))
        new_dists = Enum.reduce(new_neighbors, dists, &Map.put(&2, &1, d+1))
        get_distances(grid, new_queue, new_dists)
      _ -> dists
    end
  end

  def solve(file) do
    grid = File.read!(file)
    |> to_index_map()
    |> IO.inspect()
    
    {i,j} = grid 
    |> Enum.find(fn {_, v} -> v == ?S end)
    |> elem(0)
    #|> IO.inspect()

    c = find_tile(grid, {i,j}) 
    queue = :queue.new()
    queue = :queue.in({{i,j}, 0}, queue)

    dists = Map.put(grid, {i,j}, c)
    |> get_distances(queue, Map.new([{{i,j}, 0}]))
    
    dists
    |> Map.values()
    |> Enum.max()
    |> IO.puts()

    Map.keys(grid)
    |> Enum.filter(&inside(dists, grid, &1))
    |> Enum.map(fn {x,y} -> {y+1, x+1} end)
    |> IO.inspect()
    |> Enum.count()

  end

end

