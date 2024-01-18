defmodule AOCDay24 do
  def parse_line(s) do
    String.split(s, "@")
    |> Enum.map(fn s ->
      String.split(s, ",")
      |> Enum.map(&String.trim(&1))
      |> Enum.map(&String.to_integer(&1))
    end)
  end

  def cross_product({x1, y1}, {x2, y2}), do: x1 * y2 - y1 * x2
  def add({x1, y1}, {x2, y2}), do: {x1 + x2, y1 + y2}
  def subtract({x1, y1}, {x2, y2}), do: {x1 - x2, y1 - y2}
  def scalar_mult({x, y}, s), do: {s * x, s * y}

  def compute_time(p1, p2, v, denom) do
    subtract(p2, p1)
    |> cross_product(v)
    |> Kernel./(denom)
  end

  def compute_intersection({p1, v1}, {p2, v2}) do
    case cross_product(v1, v2) do
      0 ->
        {-1, -1}

      denom ->
        t = compute_time(p1, p2, v2, denom)
        u = compute_time(p1, p2, v1, denom)
        if t >= 0 and u >= 0, do: add(p1, scalar_mult(v1, t)), else: {-1, -1}
    end
  end

  def cross_inside({a, b}, mn, mx) do
    {xi, yi} =
      compute_intersection(a, b)

    mn <= xi and xi <= mx and mn <= yi and yi <= mx
  end

  def part1(file) do
    mn = 200_000_000_000_000
    mx = 400_000_000_000_000
    # mn = 7
    # mx = 27

    hailstones =
      File.read!(file)
      |> String.split("\n", trim: true)
      |> Enum.map(&parse_line(&1))
      |> Enum.map(fn [[x, y, _], [dx, dy, _]] -> {{x, y}, {dx, dy}} end)

    for(x <- hailstones, y <- hailstones, x != y, do: {x, y})
    |> Enum.count(&cross_inside(&1, mn, mx))
    |> div(2)
  end

  # solved part2 externally by solving the linear equation system
end
