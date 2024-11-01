defmodule AOCDay19 do
  def parse_workflow(s) do
    [name, rest] = String.split(s, "{")

    {name,
     String.trim_trailing(rest, "}")
     |> String.split(",")
     |> Enum.map(fn s ->
       case String.split(s, ":") do
         [condition, next] ->
           <<cat::binary-size(1), op::binary-size(1), val::binary>> = condition
           {cat, op, String.to_integer(val), next}

         [next] ->
           {next}
       end
     end)}
  end

  def parse_workflows(s) do
    String.split(s, "\n", trim: true)
    |> Enum.map(&parse_workflow(&1))
    |> Map.new()
  end

  def parse_part(s) do
    String.trim_leading(s, "{")
    |> String.trim_trailing("}")
    |> String.split(",")
    |> Enum.map(fn s ->
      String.split(s, "=")
      |> then(fn [c, n] -> {c, String.to_integer(n)} end)
    end)
    |> Map.new()
  end

  def get_next(_, {next}) do
    next
  end

  def get_next(part, {cat, op, val, next}) do
    case op do
      "<" -> if part[cat] < val, do: next, else: nil
      ">" -> if part[cat] > val, do: next, else: nil
    end
  end

  def eval_part(part, workflows, [head | tail]) do
    case get_next(part, head) do
      nil -> eval_part(part, workflows, tail)
      "A" -> true
      "R" -> false
      next -> eval_part(part, workflows, Map.get(workflows, next))
    end
  end

  def part1(file) do
    [workflows, parts] =
      File.read!(file)
      |> String.split("\n\n", trim: true)

    workflows = parse_workflows(workflows)

    String.split(parts, "\n", trim: true)
    |> Enum.map(&parse_part(&1))
    |> Enum.filter(&eval_part(&1, workflows, Map.get(workflows, "in")))
    |> Enum.map(&(Map.values(&1) |> Enum.sum()))
    |> Enum.sum()
  end

  def next_poss({next}, ranges) do
    [{next, ranges}]
  end

  def next_poss({cat, op, val, next}, ranges) do
    a..b = Map.get(ranges, cat)

    cond do
      (val > b and op == "<") or (val < a and op == ">") ->
        [{next, ranges}]

      (val < a and op == "<") or (val > b and op == ">") ->
        [{nil, ranges}]

      op == "<" ->
        [{next, Map.put(ranges, cat, a..(val - 1))}, {nil, Map.put(ranges, cat, val..b)}]

      op == ">" ->
        [{nil, Map.put(ranges, cat, a..val)}, {next, Map.put(ranges, cat, (val + 1)..b)}]
    end
  end

  def count_combs(workflows, ranges, [head | tail]) do
    next_poss(head, ranges)
    |> Enum.map(fn {next, nranges} ->
      case next do
        "A" ->
          Map.values(nranges)
          |> Enum.map(&Range.size(&1))
          |> Enum.product()

        "R" ->
          0

        nil ->
          count_combs(workflows, nranges, tail)

        next ->
          count_combs(workflows, nranges, Map.get(workflows, next))
      end
    end)
    |> Enum.sum()
  end

  def part2(file) do
    [workflows, _] =
      File.read!(file)
      |> String.split("\n\n", trim: true)

    workflows = parse_workflows(workflows)

    count_combs(
      workflows,
      Enum.map(["x", "m", "a", "s"], &{&1, 1..4000}) |> Map.new(),
      Map.get(workflows, "in")
    )
  end
end
