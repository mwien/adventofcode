defmodule AOCDay1a do
  def update_firstlast(t, c) do
    cond do
      ?0 <= c and c <= ?9 ->
        case elem(t, 0) do
          -1 -> {c - ?0, c - ?0}
          _ -> put_elem(t, 1, c - ?0)
        end

      true ->
        t
    end
  end

  def get_calibration(s) do
    t = Enum.reduce(s, {-1, -1}, &update_firstlast(&2, &1))
    elem(t, 0) * 10 + elem(t, 1)
  end

  def solve(file) do
    input =
      File.read!(file)
      |> String.split()

    Enum.reduce(input, 0, &(&2 + get_calibration(String.to_charlist(&1))))
  end
end
