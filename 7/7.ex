defmodule AOCDay7 do

  def card_rank(c) do
    case c do
      ?T -> 10 
      ?J -> 11 
      ?Q -> 12
      ?K -> 13
      ?A -> 14
      _ -> c - ?0
    end 
  end

  def preprocess_hand(h) do
    h = String.to_charlist(h)
    |> Enum.map(&card_rank(&1))
    freq = Enum.frequencies(h)
    |> Map.values()
    |> Enum.sort(:desc)
    {freq, h}
  end

  def part1(file) do
    File.read!(file)
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&String.split(&1))
    |> Enum.sort_by(&preprocess_hand(List.first(&1)), :asc)
    |> Enum.map(&(Enum.at(&1, 1)))
    |> Enum.map(&String.to_integer(&1))
    |> Enum.with_index(1)
    |> Enum.map(&Tuple.product(&1))
    |> Enum.sum()
  end
  
  def new_card_rank(c) do
    case c do
      ?T -> 10 
      ?J -> 1 
      ?Q -> 12
      ?K -> 13
      ?A -> 14
      _ -> c - ?0
    end 
  end
  
  def add_jokers([], j) do
    [j]
  end

  def add_jokers([h | t], j) do
    [h+j | t]
  end

  def new_preprocess_hand(h) do
    h = String.to_charlist(h)
    |> Enum.map(&new_card_rank(&1))
    {jokers, freq} = Enum.frequencies(h)
    |> Map.pop(1, 0)
    freq = Map.values(freq)
    |> Enum.sort(:desc)
    |> add_jokers(jokers)
    {freq, h}
  end

  def part2(file) do
    File.read!(file)
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&String.split(&1))
    |> Enum.sort_by(&new_preprocess_hand(List.first(&1)), :asc)
    |> Enum.map(&(Enum.at(&1, 1)))
    |> Enum.map(&String.to_integer(&1))
    |> Enum.with_index(1)
    |> Enum.map(&Tuple.product(&1))
    |> Enum.sum()
  end

end
