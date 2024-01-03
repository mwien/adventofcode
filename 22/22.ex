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

  def add_brick(grid, brick) do
    # TODO 
  end

  def fall_down(grid, {[[_, _, 1], [_, _, _]], _} = brick), do: add_brick(grid, brick) 
  def fall_down(grid, {[[x1, y1, z1], [x2, y2, z2]], id} = brick) do
    if brick_pos(brick)
    |> Enum.map(fn [x, y, z] -> [x, y, z-1] end)
    |> Enum.all?(&!Map.has_key?(grid, &1)) do
      fall_down(grid, {[[x1, y1, z1-1], [x2, y2, z2-1]], id})    
    else
      add_brick(grid, brick) 
    end 
  end

  def part1(file) do
    bricks = parse(file)
      
    support_bricks = Enum.sort_by(bricks, fn [[_, _, z], _] -> z end)
    |> Enum.reduce(%{}, &fall_down(&2, &1))
    |> Map.to_list()
    |> Enum.reduce(%{}, &add_below(&2, &1))
    |> Map.values()
    |> Enum.filter(&(length(&1) == 1))
    |> Enum.uniq()
    |> length()
    
    bricks - support_bricks
  end

end







