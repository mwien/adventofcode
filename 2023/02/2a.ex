defmodule AOCDay2a do
  def check_color(c, n) do
    {n, _} = Integer.parse(n)

    cond do
      String.starts_with?(c, "red") -> if n > 12, do: false, else: true
      String.starts_with?(c, "green") -> if n > 13, do: false, else: true
      String.starts_with?(c, "blue") -> if n > 14, do: false, else: true
      true -> true
    end
  end

  def rec_check_game(g, id) do
    case g do
      [num, color | tail] -> if check_color(color, num), do: rec_check_game(tail, id), else: 0
      _ -> id
    end
  end

  def check_game(s) do
    s_list = String.split(s)
    [_, id | game] = s_list
    {id, _} = Integer.parse(String.slice(id, 0..(String.length(id) - 2)))
    rec_check_game(game, id)
  end

  def solve(file) do
    input =
      File.read!(file)
      |> String.trim()
      |> String.split("\n")

    Enum.reduce(input, 0, &(&2 + check_game(&1)))
  end
end
