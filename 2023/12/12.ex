defmodule AOCDay12 do
  def rec_solve({i, j, k}, springs, gaps, dp) do
    cond do
      Map.has_key?(dp, {i, j, k}) ->
        {dp, dp[{i, j, k}]}

      i == 0 and j == 0 and k == 0 ->
        {dp, 1}

      i < 0 or j < 0 or k < 0 ->
        {dp, 0}

      k > 0 and springs[i] != ?. ->
        rec_solve({i - 1, j - 1, k - 1}, springs, gaps, dp)

      k == 0 and springs[i] != ?# ->
        {dp1, sol1} = rec_solve({i - 1, j, 0}, springs, gaps, dp)

        {dp2, sol2} =
          if Map.has_key?(gaps, j),
            do: rec_solve({i - 1, j, gaps[j]}, springs, gaps, dp1),
            else: {dp1, 0}

        {dp2, sol1 + sol2}

      true ->
        {dp, 0}
    end
    |> then(fn {dp, sol} -> {Map.put(dp, {i, j, k}, sol), sol} end)
  end

  def count_row([springs, seq]) do
    springs =
      (springs <> ".")
      |> String.to_charlist()
      |> Enum.with_index(fn c, i -> {i + 1, c} end)
      |> Map.new()

    seq =
      String.split(seq, ",")
      |> Enum.map(&String.to_integer(&1))

    gaps =
      seq
      |> Enum.scan(&(&1 + &2))
      |> Enum.zip(seq)
      |> Map.new()

    n = Enum.max(Map.keys(springs))
    k = Enum.max(Map.keys(gaps))
    {_, res} = rec_solve({n, k, 0}, springs, gaps, %{})
    res
  end

  def part1(file) do
    File.read!(file)
    |> String.split("\n", trim: true)
    |> Enum.map(&count_row(String.split(&1)))
    |> Enum.sum()
  end

  def modify_row(row) do
    [springs, seq] = String.split(row)

    new_springs =
      List.duplicate(springs, 5)
      |> Enum.join("?")

    new_seq =
      List.duplicate(seq, 5)
      |> Enum.join(",")

    [new_springs, new_seq]
  end

  def part2(file) do
    File.read!(file)
    |> String.split("\n", trim: true)
    |> Enum.map(&modify_row(&1))
    |> Enum.map(&count_row(&1))
    |> Enum.sum()
  end
end
