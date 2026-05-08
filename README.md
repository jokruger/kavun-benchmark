# Kavun Benchmark

Comparison of embeddable scripting engines for Go, focused on the repeated-execution use case (game NPC logic, smart
contracts, rule engines) where the host application compiles a script once and then runs it many times.

Compilation cost is intentionally excluded from measurements.

## Engines

- [Kavun](https://github.com/jokruger/kavun)
  - Kavun = default configuration
  - Kavun0 = zero pre-allocs configuration
- [Tengo](https://github.com/d5/tengo)
- [Goja](https://github.com/dop251/goja)
- [Go-lua](https://github.com/Shopify/go-lua)
- [Gopher-lua](https://github.com/yuin/gopher-lua)
- [Risor](https://github.com/risor-io/risor)
- [Starlark](https://github.com/google/starlark-go)

### Note on Kavun vs Kavun0

Kavun is optimized for short-lived scripts with a small number of variables and assignments. Internally it uses an
arena allocator with pre-allocated, typed value buffers (decimals, times, strings, arrays, dicts, iterators,
compiled/builtin functions, etc.). For workloads that fit inside those pre-allocated capacities, value construction
is effectively zero-allocation — the arena hands out slots from a pool that gets reset between runs rather than
asking the Go heap.

The two Kavun benchmarks isolate the impact of that pool:

- Kavun uses `core.NewArena(nil)`, i.e. `core.DefaultArenaOptions()` — the default pre-allocated capacities.
- Kavun0 sets every per-type capacity in `core.ArenaOptions` to `0` (decimals, times, bytes, runes, arrays,
  builtin/compiled functions, error/string/runes/bytes/array/dict values, int-range values, and all iterator pools).
  With zero capacity the arena cannot satisfy allocations from its pool, so every value falls through to the Go
  heap.

Comparing the two columns shows how much of Kavun's throughput on each task comes from the arena pool versus the
core VM and bytecode design.

## Tasks

Tasks are declared once in `task/task.go` (name + expected value + description). Each engine implements whichever subset
it can; missing tasks are penalized in the summary.

| Task              | Stresses                                                      |
| ----------------- | ------------------------------------------------------------- |
| `fib`             | Recursion, function-call overhead (exponential call tree)     |
| `fib_tail`        | Function-call overhead at linear depth (tail-recursion shape) |
| `loop_sum`        | Loops, integer arithmetic                                     |
| `sum_pow`         | Loops + integer multiplication on the hot path                |
| `closure_counter` | Closure creation + repeated invocation of one closure         |
| `closures_iife`   | Repeated closure creation-and-invoke per iteration            |
| `string_concat`   | String allocation/concatenation                               |
| `str_contains`    | Repeated substring search via the `text` stdlib module        |
| `array_sum`       | Array creation, indexed access                                |
| `nested_loop`     | Nested loop with in-place array index read+write              |

### A note on idiomatic vs portable implementations

Some engines expose richer language features than others. Kavun, for example, has fluent collection methods and arrow
lambdas that let many of these tasks be written as one-liners — `sum_pow` could be
`range(1, 10000, 1).reduce(0, (a, b) => a + b * b)`, and `string_concat` / `str_contains` have similarly compact
`.map(...)` / `.filter(...)` / `.upper()` formulations. So, an idiomatic Kavun version would have no fair counterpart
in the other engines.

To keep the comparison apples-to-apples, every task uses the same shape of implementation across all engines —
typically explicit `for` loops, indexed access, and small functions — even when an engine could express the workload
more concisely. The benchmark therefore measures the cost of executing the same algorithm on each runtime, not the
cost of each language's preferred idiom for the same problem. Showcasing engine-specific idioms is left as a
potential future "idiom" suite.

## Running

```sh
make all                       # bench + report
make bench COUNT=5             # 5 runs per (task, engine), into results/raw.txt
make report PENALTY=2.0        # parse results/raw.txt -> results/REPORT.md
```

## Methodology

- `go test -bench` with `-benchmem`. Sub-bench naming: `Benchmark<Engine>/<task>` — the report parses the engine name
  from the top-level benchmark identifier and the task from the sub-bench label.
- Each (engine, task) pair is validated for correctness once before the timed loop starts (via `task.Validate`).
- Compilation happens before `b.ResetTimer()`; the hot loop is the per-iteration script execution.

### Aggregation

Per-task tables are sorted by `ns/op` ascending, with a `vs best` ratio relative to the fastest engine on that task.

The summary uses two complementary metrics, since tasks are not equally expensive and a plain average across tasks would
be dominated by the slowest:

- **CPU geomean** — geometric mean of per-task `ns/op` ratios. Lower is better; `1.00×` means "fastest on every task".
- **Avg rank** — mean rank across tasks. Robust but discards magnitude.

We also report worst-case ratio, win count (rank-1), and a memory geomean.

### Missing-task penalty

If an engine doesn't implement a task, the report still includes the engine in that task's table but marks it
`_missing_` and substitutes a penalty for both the geomean and the avg-rank computation:

- penalty ratio = (worst observed ratio on this task) × `-penalty` (default `2.0`)
- penalty rank = (number of engines that ran this task) + 1

This means skipping an easy task (where everyone is close to the best) costs less than skipping a hard one (where the
worst engine is already far from the best) — proportional to the difficulty of the task.
