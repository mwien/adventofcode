defmodule AOCDay1b do

  def update_firstlast(t, s) do
    c = [{"one", "1"}, {"two", "2"}, {"three", "3"}, {"four", "4"}, {"five", "5"}, {"six", "6"}, {"seven", "7"}, {"eight", "8"}, {"nine", "9"}]
      |> Enum.reduce(s, &String.replace_prefix(&2, elem(&1, 0), elem(&1, 1)))
      |> String.to_charlist()
      |> List.first()
    cond  do
      ?0 <= c and c <= ?9 -> case elem(t, 0) do
        -1 -> {c - ?0, c - ?0}
        _ -> put_elem(t, 1, c - ?0)
      end
      true -> t
    end
  end

  def get_calibration(s) do
    t = Enum.to_list(0..String.length(s)-1)
      |> Enum.reduce({-1, -1}, &update_firstlast(&2, String.slice(s, &1, 5)))
    elem(t, 0) * 10 + elem(t, 1)
  end

  def solve(file) do
    input = File.read!(file)
    |> String.split()
    Enum.reduce(input, 0, &(&2 + get_calibration(&1)))
  end
end


