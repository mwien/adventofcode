defmodule AOCDay5 do

  def int_parse(s) do
    {res, _} = Integer.parse(s)
    res
  end

  def parse_map([]) do
    {[], []}
  end

  def parse_map(input) do
    [head | tail] = input 
    if head == "" do
      {[], tail}
    else 
      {map, tail} = parse_map(tail)
      {[Enum.map(String.split(head), &int_parse(&1)) | map], tail}
    end
  end

  def parse_maps([]) do
    []
  end

  def parse_maps(input) do
    [head | tail] = input
    if String.ends_with?(head, "map:") do
      {map, tail} = parse_map(tail)
      [map | parse_maps(tail)]
    else
      parse_maps(tail)
    end
  end

  def parse(input) do
    [seedinfo | rest] = input 
    [_ | seedinfo] = String.split(seedinfo)
    seedinfo = Enum.map(seedinfo, &int_parse(&1))
    maps = parse_maps(rest)
    {seedinfo, maps}
  end

  def check_transform({object, true}, _) do
    {object, true} 
  end

  def check_transform({object, done}, [t, s, l]) do
    if s <= object and object < s + l do
      {object - s + t, true}
    else 
      {object, done}
    end
  end

  def transform(object, map) do
    #IO.inspect(object)
    {object, _} = Enum.reduce(map, {object, false}, &check_transform(&2, &1)) 
    object
  end

  def location_for_seed(seed, maps) do
    Enum.reduce(maps, seed, &transform(&2, &1))
    #|> IO.inspect()
  end

  def parta(file) do
    input = File.read!(file)
    |> String.trim()
    |> String.split("\n")

    {seeds, maps} = parse(input)
    Enum.reduce(seeds, 10**18, &(min(&2, location_for_seed(&1, maps))))
  end

end
