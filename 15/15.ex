defmodule AOCDay15 do
  def reindeer_hash(s) do
    String.to_charlist(s)
    |> Enum.reduce(0, fn c, curr ->
      rem((curr + c) * 17, 256)
    end)
  end

  def part1(file) do
    File.read!(file)
    |> String.trim()
    |> String.split(",")
    |> Enum.map(&reindeer_hash(&1))
    |> Enum.sum()
  end

  def perform_remove([s], boxes) do
    b = reindeer_hash(s)

    if Map.has_key?(boxes, b) do
      Map.update!(boxes, b, fn l ->
        Enum.filter(l, fn {_, l} -> l != s end)
      end)
    else
      boxes
    end
  end

  def perform_add([s, i], boxes) do
    b = reindeer_hash(s)

    if Map.has_key?(boxes, b) do
      l = boxes[b]

      case Enum.find_index(l, fn {_, l} -> l == s end) do
        nil -> Map.put(boxes, b, List.insert_at(l, -1, {i, s}))
        j -> Map.put(boxes, b, List.replace_at(l, j, {i, s}))
      end
    else
      Map.put(boxes, b, [{i, s}])
    end
  end

  def perform_operation(s, boxes) do
    if String.contains?(s, "-") do
      perform_remove(String.split(s, "-", trim: true), boxes)
    else
      perform_add(String.split(s, "="), boxes)
    end
  end

  def focusing_power(boxes) do
    Map.to_list(boxes)
    |> Enum.map(fn {b, l} ->
      Enum.with_index(l)
      |> Enum.map(fn {{i, _}, j} ->
        (b + 1) * (j + 1) * String.to_integer(i)
      end)
      |> Enum.sum()
    end)
    |> Enum.sum()
  end

  def part2(file) do
    File.read!(file)
    |> String.trim()
    |> String.split(",")
    |> Enum.reduce(%{}, &perform_operation(&1, &2))
    |> focusing_power()
  end
end
