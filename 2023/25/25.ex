defmodule AOCDay25 do
  def inc(v) do
    if v == nil, do: 1, else: v + 1
  end

  def add_edges(s, graph) do
    [src, rest] = String.split(s, ": ")

    String.split(rest)
    |> Enum.reduce(graph, fn tgt, graph ->
      graph
      |> Map.put_new(src, %{})
      |> Map.put_new(tgt, %{})
      |> update_in([src, tgt], &inc(&1))
      |> update_in([tgt, src], &inc(&1))
    end)
  end

  def mincut(graph, nodes) do
    case map_size(graph) do
      2 ->
        {Enum.random(graph)
         |> elem(1)
         |> Enum.random()
         |> elem(1),
         Map.values(nodes)
         |> Enum.product()}

      _ ->
        {src, rest} = Enum.random(graph)
        {tgt, _} = Enum.random(rest)

        graph =
          Map.keys(graph[tgt])
          |> Enum.reduce(graph, fn ngh, graph ->
            sz = graph[ngh][tgt]

            graph
            |> Map.update!(ngh, &Map.delete(&1, tgt))
            |> update_in([ngh, src], fn old_sz ->
              if old_sz == nil, do: sz, else: old_sz + sz
            end)
          end)
          |> Map.put(
            src,
            Map.merge(graph[src], graph[tgt], fn _, v1, v2 -> v1 + v2 end)
            |> Map.delete(tgt)
            |> Map.delete(src)
          )
          |> Map.delete(tgt)

        nodes =
          Map.update!(nodes, src, &(&1 + nodes[tgt]))
          |> Map.delete(tgt)

        mincut(graph, nodes)
    end
  end

  def solve(graph, nodes) do
    case mincut(graph, nodes) do
      {3, value} -> value
      _ -> solve(graph, nodes)
    end
  end

  def part1(file) do
    graph =
      File.read!(file)
      |> String.split("\n", trim: true)
      |> Enum.reduce(%{}, &add_edges(&1, &2))

    solve(
      graph,
      Map.keys(graph)
      |> Map.new(&{&1, 1})
    )
  end
end
