defmodule AOCDay20 do

  def parse_module(s) do
    [label, children] = String.split(s, "->")
    {type, name} = case String.trim(label) do
      "broadcaster" -> {"B", "broadcaster"}
      "%" <> rest -> {"%", rest}
      "&" <> rest -> {"&", rest}
    end

    children = String.trim(children)
    |> String.split(", ")

    {name, %{type: type, children: children, parents: [], state: (if type == "&", do: %{}, else: 0)}}
  end

  def parse_modules(s) do
    modules = String.split(s, "\n", trim: true) 
    |> Enum.map(&parse_module(&1))
    |> Map.new()


    modules = Enum.reduce(Map.keys(modules), modules, fn mod, modules -> 
      Enum.reduce(modules[mod].children, modules, fn child, modules ->
        if Map.has_key?(modules, child) do 
          vals = modules[child] 
          parents = vals.parents
          Map.put(modules, child, Map.put(vals, :parents, [mod | parents]))
        else 
          modules
        end
      end)
    end) 

    Enum.reduce(Map.keys(modules), modules, fn mod, modules -> 
      val = modules[mod]
      if val.type == "&" do
        update_in(modules[mod][:state], fn _ -> Map.new(val.parents, fn c -> {c, 0} end) end)
      else 
        modules
      end
    end)
  end

  def update_modules(modules, name, pulse, from) do
    case modules[name].type do
      "%" -> update_in(modules[name][:state], &(if pulse == 0, do: 1 - &1, else: &1))
      "&" -> update_in(modules[name][:state][from], fn _ -> pulse end)
      _ -> modules
    end
  end
  
  def propagate(queue, children, val, from) do
    Enum.reduce(children, queue, &:queue.in({&1, val, from}, &2))
  end

  def send_pulse(modules, queue, name, pulse) do
    case modules[name].type do
      "B" -> propagate(queue, modules[name].children, 0, name)
      "%" -> 
        case pulse do  
          0 -> propagate(queue, modules[name].children, modules[name].state, name)
          1 -> queue
        end
      "&" -> propagate(queue, modules[name].children, (if Enum.all?(Map.values(modules[name].state), &(&1 == 1)), do: 0, else: 1), name) 
    end
  end

  def simulate(modules, {[], []}, {l, h}, _) do
    {modules, {l, h}}
  end

  def simulate(modules, queue, {l, h}, b) do
    {{:value, {name, pulse, from}}, queue} = :queue.out(queue)
    if Map.has_key?(modules, name) do
      modules = if pulse == 0, do: Map.put(modules, name, Map.put_new(modules[name], :first, b)), else: modules
      modules = update_modules(modules, name, pulse, from)
      queue = send_pulse(modules, queue, name, pulse)
      simulate(modules, queue, {l + 1 - pulse, h + pulse}, b)
    else 
      simulate(modules, queue, {l + 1 - pulse, h + pulse}, b) 
    end
  end

  def part1(file) do
    modules = File.read!(file)
    |> parse_modules()

    {_, {l, h}} = Enum.reduce(1..1000, {modules, {0,0}}, fn b, {modules, counts} -> 
      simulate(modules, :queue.from_list([{"broadcaster", 0, "button"}]), counts, b) 
    end)

    l * h
  end
  
  def lcm(a, b) do
    div(abs(a * b), Integer.gcd(a, b))
  end
  
  def part2(file) do
    modules = File.read!(file)
    |> parse_modules()

    {modules, _} = Enum.reduce(1..100000, {modules, {0,0}}, fn b, {modules, counts} -> 
      simulate(modules, :queue.from_list([{"broadcaster", 0, "button"}]), counts, b) 
    end)
   
    # module names taken from input file
    # parents of parents of rx
    [modules["ft"].first, modules["jz"].first, modules["sv"].first, modules["ng"].first]
    |> Enum.reduce(&lcm(&1, &2))
  end
end




