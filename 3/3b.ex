defmodule AOCDay3b do

  def put_marks_on_neighbors(marks, i, j) do
    neighbors = for x <- [-1, 0, 1], y <- [-1, 0, 1], do: {i+x, j+y} 
    id = Enum.random(0..1_000_000_000)
    Enum.reduce(neighbors, marks, &Map.put(&2, &1, id))
  end

  def put_marks_for_field(marks, i, j, c) do
    cond do
      c == ?. -> marks 
      c >= ?0 and c <= ?9 -> marks 
      true -> put_marks_on_neighbors(marks, i, j)
    end
  end

  def put_marks_for_row(marks, i, row) do
    Enum.zip(1..String.length(row), String.to_charlist(row))
    |> Enum.reduce(marks, &put_marks_for_field(&2, i, elem(&1, 0), elem(&1, 1)))
  end

  def update_entry(mp, -1, _nv) do
    mp
  end

  def update_entry(mp, id, nv) do
    {ov, num} = Map.get(mp, id, {1, 0})
    Map.put(mp, id, {ov * nv, num+1})
  end

  def update_entries(mp, id, nv) do
    id = Enum.uniq(id)
    Enum.reduce(id, mp, &update_entry(&2, &1, nv))
  end

  def update_sum({mp, cur, id}, i, j, c, marks) do
    cond do
      c >= ?0 and c <= ?9 -> {mp, cur * 10 + c - ?0, [Map.get(marks, {i, j}, -1) | id] }
      id >= 0 -> {update_entries(mp, id, cur), 0, []}
      true -> {mp, 0, []}
    end
  end

  def gear_sum(mp, i, row, marks) do
    Enum.zip(1..String.length(row), String.to_charlist(row))
    |> Enum.reduce({mp, 0, []}, &update_sum(&2, i, elem(&1, 0), elem(&1, 1), marks))
    |> elem(0)
  end

  def solve(file) do
    input = File.read!(file)
    |> String.trim()
    |> String.split("\n")

    marks = Enum.reduce(Enum.zip(1..length(input), input), %{}, &put_marks_for_row(&2, elem(&1, 0), elem(&1, 1)))
    Enum.reduce(Enum.zip(1..length(input), input), %{}, &gear_sum(&2, elem(&1, 0), elem(&1, 1) <> ".", marks))
    |> Map.values()
    |> Enum.reduce(0, &(&2 + if elem(&1, 1) == 2, do: elem(&1, 0), else: 0))
  end
end
