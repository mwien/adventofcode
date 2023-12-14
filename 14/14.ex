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

  def rotate(grid) do
    {_, m} = Map.keys(grid)
    |> Enum.max()
    Map.new(grid, fn {{i,j}, _} -> {{i,j}, grid[{j, m-i}]} end)
  end

  def roll_south(grid) do
    new_grid = Enum.map(Map.keys(grid), fn {i, j} ->
      if Map.get(grid, {i,j}) == ?O and Map.get(grid, {i,j+1}) == ?. do
        {i,j}
      else 
        {-1, -1}
      end
    end)  
    |> Enum.filter(fn {i, _} -> i >= 0 end)
    |> Enum.reduce(grid, fn {i,j}, grid -> 
        Map.put(grid, {i,j}, ?.)
        |> Map.put({i,j+1}, ?O)
      end)
    if new_grid == grid, do: new_grid, else: roll_south(new_grid)
  end

  def compute_score(grid) do
    Enum.map(Map.keys(grid), fn {i,j} -> 
      if grid[{i,j}] == ?O, do: j+1, else: 0  
    end)
    |> Enum.sum()
  end

  def part1(file) do
    grid = File.read!(file)
    |> to_index_map()
    
    grid 
    |> rotate()
    |> rotate()
    |> roll_south()
    |> compute_score()
  end

  def forward(grid) do
    Enum.reduce(1..4, grid, fn _, grid -> roll_south(grid) |> rotate() end)
  end

  def cycle(grid, past, reps) do
    if Map.has_key?(past, grid) do
      {grid, reps, reps - Map.get(past, grid)}
    else 
      past = Map.put(past, grid, reps)
      cycle(forward(grid), past, reps+1)
    end
  end

  def part2(file) do
    grid = File.read!(file)
    |> to_index_map()
   
    startgrid = grid 
    |> rotate()
    |> rotate()
    
    {cgrid, time, length} = cycle(startgrid, %{}, 0)
    rest = rem(10**9 - time, length)
    Enum.reduce(1..rest, cgrid, fn _, grid -> forward(grid) end)
    |> compute_score()
  end

end

