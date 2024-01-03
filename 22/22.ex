defmodule AOCDay22 do

  def parse_line(line) do
    String.split(line, "~")
    |> Enum.map(fn pos -> 
      String.split(pos, ",")
      |> Enum.map(&String.to_integer(&1))
    end)
  end
    
  def parse(file) do
    File.read!(file)
    |> String.split("\n", trim: true)
    |> Enum.map(&parse_line(&1))
  end

  def brick_pos([[x1, y1, z1], [x2, y2, z2]]) do
    for x <- x1..x2, y <- y1..y2, z <- z1..z2, do: [x, y, z]
  end

  def below_pos([x, y, z]), do: [x, y, z-1]

  def add_brick(grid, {pos, id}) do
    brick_pos(pos)
    |> Enum.reduce(grid, &Map.put(&2, &1, id))
  end

  def fall_down(grid, {[[_, _, 1], [_, _, _]], _} = brick), do: add_brick(grid, brick) 
  def fall_down(grid, {[spos, epos], id} = brick) do
    if brick_pos([spos, epos])
    |> Enum.map(&below_pos(&1))
    |> Enum.all?(&!Map.has_key?(grid, &1)) do
      fall_down(grid, {[below_pos(spos), below_pos(epos)], id})    
    else
      add_brick(grid, brick) 
    end 
  end

  def add_below(below, {pos, id}, grid) do
    case Map.get(grid, below_pos(pos)) do
      nil -> below
      ^id -> below
      b_id -> Map.update(below, id, MapSet.new([b_id]), &MapSet.put(&1, b_id)) 
    end
  end

  def part1(file) do
    bricks = parse(file)
      
    grid = Enum.sort_by(bricks, fn [[_, _, z], _] -> z end)
    |> Enum.with_index()
    |> Enum.reduce(%{}, &fall_down(&2, &1))
    
    support_bricks = Map.to_list(grid)
    |> Enum.reduce(%{}, &add_below(&2, &1, grid))
    |> Map.values()
    |> Enum.filter(&(MapSet.size(&1) == 1))
    |> Enum.uniq()
    |> length()
    
    length(bricks) - support_bricks
  end 

  def remove_brick(below, id) do
    # TODO
  end

  chain_reaction(below) do
    # remove brick, here or before first call
  end

  def part2(file) do
    bricks = parse(file)
      
    grid = Enum.sort_by(bricks, fn [[_, _, z], _] -> z end)
    |> Enum.with_index()
    |> Enum.reduce(%{}, &fall_down(&2, &1))
    
    below = Map.to_list(grid)
    |> Enum.reduce(%{}, &add_below(&2, &1, grid))

    Enum.map(1..length(bricks), chain_reaction(below))
  end

end







