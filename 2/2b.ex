defmodule AOCDay2b do
  def update_maxs(max_dict, [num, color]) do
    {num, _} = Integer.parse(num)

    color =
      String.trim_trailing(color, ",")
      |> String.trim_trailing(";")

    Map.put(max_dict, color, max(Map.fetch!(max_dict, color), num))
  end

  def game_power(s) do
    [_, _ | game] = String.split(s)

    Enum.chunk_every(game, 2)
    |> Enum.reduce(%{"red" => 0, "green" => 0, "blue" => 0}, &update_maxs(&2, &1))
    |> Map.values()
    |> Enum.reduce(&*/2)
  end

  def solve(file) do
    input =
      File.read!(file)
      |> String.trim()
      |> String.split("\n")

    IO.puts(Enum.reduce(input, 0, &(&2 + game_power(&1))))
  end
end
