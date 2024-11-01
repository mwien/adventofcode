defmodule AOCDay3a do
  def put_marks_on_neighbors(marks, i, j) do
    neighbors = for x <- [-1, 0, 1], y <- [-1, 0, 1], do: {i + x, j + y}
    Enum.reduce(neighbors, marks, &Map.put(&2, &1, true))
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

  def update_sum({sum, cur, sym}, i, j, c, marks) do
    cond do
      c >= ?0 and c <= ?9 -> {sum, cur * 10 + c - ?0, sym or Map.get(marks, {i, j}, false)}
      sym -> {sum + cur, 0, false}
      true -> {sum, 0, false}
    end
  end

  def compute_sum(i, row, marks) do
    Enum.zip(1..String.length(row), String.to_charlist(row))
    |> Enum.reduce({0, 0, false}, &update_sum(&2, i, elem(&1, 0), elem(&1, 1), marks))
    |> elem(0)
  end

  def solve(file) do
    input =
      File.read!(file)
      |> String.trim()
      |> String.split("\n")

    marks =
      Enum.zip(1..length(input), input)
      |> Enum.reduce(%{}, &put_marks_for_row(&2, elem(&1, 0), elem(&1, 1)))

    Enum.zip(1..length(input), input)
    |> Enum.reduce(0, &(&2 + compute_sum(elem(&1, 0), elem(&1, 1) <> ".", marks)))
  end
end
