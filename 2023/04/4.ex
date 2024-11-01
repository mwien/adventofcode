defmodule AOCDay4 do
  import Bitwise

  def matches(card) do
    [first, second] = String.split(card, "|")
    [_, _ | winnums] = String.split(first)
    winnumset = MapSet.new(winnums, &Integer.parse(&1))
    cardnums = Enum.map(String.split(second), &Integer.parse(&1))
    Enum.reduce(cardnums, 0, &(&2 + if(MapSet.member?(winnumset, &1), do: 1, else: 0)))
  end

  def score(matches) do
    if matches == 0, do: 0, else: 1 <<< (matches - 1)
  end

  def points(card) do
    matches(card)
    |> score()
  end

  def parta(file) do
    input =
      File.read!(file)
      |> String.trim()
      |> String.split("\n")

    Enum.reduce(input, 0, &(&2 + points(&1)))
  end

  def add_copies(copies, i, j) do
    cur = Map.fetch!(copies, i)

    case Map.fetch(copies, i + j) do
      {:ok, cp} -> Map.put(copies, i + j, cp + cur)
      _ -> copies
    end
  end

  def update_copies(copies, i, card) do
    m = matches(card)
    if m == 0, do: copies, else: Enum.reduce(1..m, copies, &add_copies(&2, i, &1))
  end

  def partb(file) do
    input =
      File.read!(file)
      |> String.trim()
      |> String.split("\n")

    Enum.reduce(
      Enum.zip(input, 1..length(input)),
      Map.new(1..length(input), &{&1, 1}),
      &update_copies(&2, elem(&1, 1), elem(&1, 0))
    )
    |> Map.values()
    |> Enum.sum()
  end
end
