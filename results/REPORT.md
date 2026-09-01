# Kavun Benchmark Report

Comparison of embeddable scripting engines for Go, focused on repeated execution of pre-compiled scripts.
Compilation cost is excluded from measurements.

## Host

| Field | Value |
|-------|-------|
| OS | darwin |
| Arch | amd64 |
| CPU | Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz |
| Logical CPUs | 12 |
| GOMAXPROCS | 12 |
| Total memory | 16.00 GB |
| Go version | go1.26.4 |
| Run at (UTC) | 2026-09-01T12:30:32Z |

## Summary

Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.
Lower is better; 1.00× means "fastest on every task".

Missing-task penalty: ratio = (worst observed ratio on the task) × 2.00, rank = last + 1.

| Rank | Engine | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |
|------|--------|-------------|----------|-------------|------|-------------|-----------|---------|
| 1 | kavun | 1.05× | 1.44 | 1.33× | 6 | 1.40× | 9 | 0 |
| 2 | gopherlua | 1.44× | 2.56 | 3.03× | 3 | 163.50× | 9 | 0 |
| 3 | golua | 1.51× | 3.33 | 1.85× | 0 | 228.44× | 9 | 0 |
| 4 | starlark | 2.32× | 4.11 | 5.56× | 0 | 158.82× | 9 | 0 |
| 5 | tengo | 2.94× | 4.33 | 44.34× | 0 | 1177.80× | 9 | 0 |
| 6 | goja | 4.87× | 6.22 | 11.91× | 0 | 297.35× | 9 | 0 |
| 7 | risor | 5.82× | 6.00 | 136.10× | 0 | 3102.55× | 9 | 0 |

## Per-task results

### fib

_Naive recursive Fibonacci(25). Stresses function-call overhead and recursion._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 13.86 ms | 8 B | 1 | 1.00× |
| 2 | tengo | 16.49 ms | 2.51 MB | 317814 | 1.19× |
| 3 | golua | 19.24 ms | 2.42 MB | 317817 | 1.39× |
| 4 | gopherlua | 19.67 ms | 10.52 KB | 42 | 1.42× |
| 5 | goja | 34.67 ms | 7.96 KB | 392 | 2.50× |
| 6 | risor | 39.55 ms | 4.01 MB | 243204 | 2.85× |
| 7 | starlark | 50.80 ms | 18.53 MB | 242828 | 3.67× |

### fib_tail

_Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 1.39 µs | 8 B | 1 | 1.00× |
| 2 | golua | 2.47 µs | 544 B | 47 | 1.77× |
| 3 | gopherlua | 3.11 µs | 2.89 KB | 12 | 2.23× |
| 4 | goja | 5.48 µs | 5.08 KB | 23 | 3.94× |
| 5 | starlark | 7.75 µs | 5.10 KB | 58 | 5.56× |
| 6 | tengo | 17.30 µs | 88.48 KB | 44 | 12.42× |
| 7 | risor | 52.84 µs | 309.24 KB | 74 | 37.95× |

### loop_sum

_Sum of integers 1..10000. Stresses loops and integer arithmetic._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | gopherlua | 355.47 µs | 234.98 KB | 938 | 1.00× |
| 2 | kavun | 392.49 µs | 8 B | 1 | 1.10× |
| 3 | golua | 464.16 µs | 234.54 KB | 30006 | 1.31× |
| 4 | starlark | 617.41 µs | 1016 B | 19 | 1.74× |
| 5 | tengo | 730.30 µs | 244.44 KB | 20006 | 2.05× |
| 6 | risor | 1.43 ms | 615.82 KB | 19762 | 4.03× |
| 7 | goja | 4.23 ms | 1.07 MB | 39729 | 11.91× |

### sum_pow

_Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | gopherlua | 493.47 µs | 313.06 KB | 1250 | 1.00× |
| 2 | kavun | 544.17 µs | 8 B | 1 | 1.10× |
| 3 | golua | 635.83 µs | 312.63 KB | 40002 | 1.29× |
| 4 | tengo | 997.81 µs | 322.53 KB | 30002 | 2.02× |
| 5 | risor | 1.80 ms | 772.02 KB | 29758 | 3.64× |
| 6 | starlark | 2.16 ms | 1018.29 KB | 40711 | 4.37× |
| 7 | goja | 5.19 ms | 1.14 MB | 49723 | 10.52× |

### closure_counter

_Closure that increments a captured counter, called 1000 times._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 68.42 µs | 136 B | 4 | 1.00× |
| 2 | golua | 101.94 µs | 23.69 KB | 3011 | 1.49× |
| 3 | gopherlua | 105.04 µs | 23.35 KB | 94 | 1.54× |
| 4 | tengo | 117.98 µs | 103.92 KB | 2010 | 1.72× |
| 5 | starlark | 229.93 µs | 63.79 KB | 1028 | 3.36× |
| 6 | risor | 268.11 µs | 331.79 KB | 1542 | 3.92× |
| 7 | goja | 430.71 µs | 107.18 KB | 3514 | 6.29× |

### string_concat

_Build a string of 200 characters via repeated concatenation; return its length._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | gopherlua | 28.48 µs | 27.98 KB | 413 | 1.00× |
| 2 | starlark | 28.51 µs | 25.12 KB | 420 | 1.00× |
| 3 | kavun | 37.75 µs | 24.09 KB | 399 | 1.33× |
| 4 | tengo | 49.92 µs | 120.13 KB | 606 | 1.75× |
| 5 | golua | 52.70 µs | 39.85 KB | 1403 | 1.85× |
| 6 | risor | 82.33 µs | 333.35 KB | 438 | 2.89× |
| 7 | goja | 90.35 µs | 43.05 KB | 804 | 3.17× |

### string_repeat

_Build a string of 200 'x' characters using the language's idiomatic repeat builtin; return its length._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 366 ns | 224 B | 2 | 1.00× |
| 2 | golua | 531 ns | 400 B | 8 | 1.45× |
| 3 | starlark | 919 ns | 1.07 KB | 17 | 2.51× |
| 4 | gopherlua | 1.11 µs | 2.93 KB | 11 | 3.03× |
| 5 | goja | 1.28 µs | 416 B | 8 | 3.50× |
| 6 | tengo | 16.23 µs | 88.48 KB | 10 | 44.34× |
| 7 | risor | 49.80 µs | 307.85 KB | 42 | 136.10× |

### str_contains

_1000 substring searches over a 1000-char haystack, alternating an 8-char hit and miss needle. Stresses substring search._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 412.86 µs | 32.27 KB | 1003 | 1.00× |
| 2 | starlark | 414.98 µs | 2.15 KB | 24 | 1.01× |
| 3 | tengo | 506.38 µs | 140.03 KB | 3507 | 1.23× |
| 4 | gopherlua | 532.02 µs | 28.09 KB | 1077 | 1.29× |
| 5 | risor | 563.17 µs | 355.80 KB | 2035 | 1.36× |
| 6 | golua | 635.29 µs | 63.69 KB | 6008 | 1.54× |
| 7 | goja | 993.16 µs | 134.12 KB | 6502 | 2.41× |

### array_dot

_Pre-allocate two 500-element arrays, fill via indexed loops, then compute the integer dot product via an indexed loop._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 91.15 µs | 24.23 KB | 11 | 1.00× |
| 2 | gopherlua | 129.29 µs | 67.23 KB | 135 | 1.42× |
| 3 | starlark | 135.60 µs | 17.30 KB | 32 | 1.49× |
| 4 | golua | 142.36 µs | 72.27 KB | 4527 | 1.56× |
| 5 | tengo | 149.64 µs | 152.34 KB | 3527 | 1.64× |
| 6 | risor | 208.92 µs | 352.83 KB | 1890 | 2.29× |
| 7 | goja | 677.41 µs | 135.68 KB | 5160 | 7.43× |

