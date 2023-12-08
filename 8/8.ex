defmodule AOCDay8 do
  def parse(file) do
    [sequence, rest] = File.read!(file)
      |> String.split("\n\n")
    rest = String.trim(rest)
      |> String.split("\n")
      |> Enum.map(fn ln ->
        [[node], [left], [right]] = Regex.scan(~r/[[:alnum:]]+/, ln)
        {node, {left, right}}
      end)
      |> Map.new()
    sequence = String.to_charlist(sequence)
      |> Enum.map(fn c ->
        if ?L == c, do: 0, else: 1
      end)
    {sequence, rest}
  end

  def forward_step(node, sequence, network, n, depth) do
    elem(network[node], sequence[rem(depth, n)])
  end

  def search(node, sequence, network, n, depth) do
    if String.ends_with?(node, "Z") do
      depth
    else
      search(forward_step(node, sequence, network, n, depth), sequence, network, n, depth + 1)
    end
  end

  def part1(file) do
    {sequence, network} = parse(file)
    search("AAA", Map.new(Enum.with_index(sequence, &({&2, &1}))), network, length(sequence), 0)
  end

  def lcm(a, b) do
    div(abs(a * b), Integer.gcd(a, b))
  end

  def part2(file) do
    {sequence, network} = parse(file)
    seq_map = Map.new(Enum.with_index(sequence, &({&2, &1})))
    Map.keys(network) 
    |> Enum.filter(&String.ends_with?(&1, "A"))
    |> Enum.map(&search(&1, seq_map, network, length(sequence), 0))
    |> Enum.reduce(&lcm(&1, &2))
  end
end
