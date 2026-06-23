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
| Run at (UTC) | 2026-06-23T06:12:00Z |

## Summary

Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.
Lower is better; 1.00× means "fastest on every task".

Missing-task penalty: ratio = (worst observed ratio on the task) × 2.00, rank = last + 1.

| Rank | Engine | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |
|------|--------|-------------|----------|-------------|------|-------------|-----------|---------|
| 1 | kavun | 1.01× | 1.11 | 1.05× | 8 | 1.04× | 9 | 0 |
| 2 | gopherlua | 1.59× | 2.78 | 4.87× | 1 | 180.03× | 9 | 0 |
| 3 | golua | 1.61× | 3.33 | 2.35× | 0 | 251.53× | 9 | 0 |
| 4 | starlark | 2.55× | 4.22 | 5.17× | 0 | 174.88× | 9 | 0 |
| 5 | tengo | 3.15× | 4.33 | 69.56× | 0 | 1296.87× | 9 | 0 |
| 6 | goja | 5.16× | 6.22 | 10.85× | 0 | 327.41× | 9 | 0 |
| 7 | risor | 6.32× | 6.00 | 230.43× | 0 | 3416.20× | 9 | 0 |

## Per-task results

### fib

_Naive recursive Fibonacci(25). Stresses function-call overhead and recursion._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 14.90 ms | 8 B | 1 | 1.00× |
| 2 | tengo | 17.22 ms | 2.51 MB | 317814 | 1.16× |
| 3 | golua | 19.26 ms | 2.42 MB | 317817 | 1.29× |
| 4 | gopherlua | 21.06 ms | 10.53 KB | 42 | 1.41× |
| 5 | goja | 33.92 ms | 7.96 KB | 392 | 2.28× |
| 6 | risor | 42.04 ms | 4.01 MB | 243204 | 2.82× |
| 7 | starlark | 52.83 ms | 18.53 MB | 242828 | 3.55× |

### fib_tail

_Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 1.54 µs | 8 B | 1 | 1.00× |
| 2 | golua | 2.45 µs | 544 B | 47 | 1.59× |
| 3 | gopherlua | 3.12 µs | 2.89 KB | 12 | 2.03× |
| 4 | goja | 5.55 µs | 5.08 KB | 23 | 3.61× |
| 5 | starlark | 7.96 µs | 5.10 KB | 58 | 5.17× |
| 6 | tengo | 17.47 µs | 88.48 KB | 44 | 11.35× |
| 7 | risor | 54.27 µs | 309.24 KB | 74 | 35.26× |

### loop_sum

_Sum of integers 1..10000. Stresses loops and integer arithmetic._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | gopherlua | 388.75 µs | 234.98 KB | 938 | 1.00× |
| 2 | kavun | 407.65 µs | 8 B | 1 | 1.05× |
| 3 | golua | 465.94 µs | 234.54 KB | 30006 | 1.20× |
| 4 | starlark | 626.82 µs | 1016 B | 19 | 1.61× |
| 5 | tengo | 754.86 µs | 244.44 KB | 20006 | 1.94× |
| 6 | risor | 1.42 ms | 615.82 KB | 19762 | 3.65× |
| 7 | goja | 4.22 ms | 1.07 MB | 39729 | 10.85× |

### sum_pow

_Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 476.51 µs | 8 B | 1 | 1.00× |
| 2 | gopherlua | 526.12 µs | 313.06 KB | 1250 | 1.10× |
| 3 | golua | 641.63 µs | 312.63 KB | 40002 | 1.35× |
| 4 | tengo | 1.04 ms | 322.53 KB | 30002 | 2.18× |
| 5 | risor | 1.81 ms | 772.02 KB | 29758 | 3.81× |
| 6 | starlark | 2.20 ms | 1018.29 KB | 40711 | 4.62× |
| 7 | goja | 5.11 ms | 1.14 MB | 49723 | 10.73× |

### closure_counter

_Closure that increments a captured counter, called 1000 times._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 73.67 µs | 120 B | 4 | 1.00× |
| 2 | golua | 101.56 µs | 23.69 KB | 3011 | 1.38× |
| 3 | gopherlua | 108.92 µs | 23.35 KB | 94 | 1.48× |
| 4 | tengo | 115.21 µs | 103.92 KB | 2010 | 1.56× |
| 5 | starlark | 251.67 µs | 63.79 KB | 1028 | 3.42× |
| 6 | risor | 268.50 µs | 331.79 KB | 1542 | 3.64× |
| 7 | goja | 432.17 µs | 107.18 KB | 3514 | 5.87× |

### string_concat

_Build a string of 200 characters via repeated concatenation; return its length._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 25.39 µs | 24.09 KB | 399 | 1.00× |
| 2 | gopherlua | 28.30 µs | 27.98 KB | 413 | 1.11× |
| 3 | starlark | 28.94 µs | 25.12 KB | 420 | 1.14× |
| 4 | tengo | 48.42 µs | 120.13 KB | 606 | 1.91× |
| 5 | golua | 52.76 µs | 39.85 KB | 1403 | 2.08× |
| 6 | risor | 84.68 µs | 333.35 KB | 438 | 3.34× |
| 7 | goja | 91.74 µs | 43.05 KB | 804 | 3.61× |

### string_repeat

_Build a string of 200 'x' characters using the language's idiomatic repeat builtin; return its length._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 230 ns | 224 B | 2 | 1.00× |
| 2 | golua | 542 ns | 400 B | 8 | 2.35× |
| 3 | starlark | 925 ns | 1.07 KB | 17 | 4.02× |
| 4 | gopherlua | 1.12 µs | 2.93 KB | 11 | 4.87× |
| 5 | goja | 1.29 µs | 416 B | 8 | 5.60× |
| 6 | tengo | 16.02 µs | 88.48 KB | 10 | 69.56× |
| 7 | risor | 53.06 µs | 307.85 KB | 42 | 230.43× |

### str_contains

_1000 substring searches over a 1000-char haystack, alternating an 8-char hit and miss needle. Stresses substring search._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 338.02 µs | 1.02 KB | 3 | 1.00× |
| 2 | starlark | 443.43 µs | 2.15 KB | 24 | 1.31× |
| 3 | tengo | 518.24 µs | 140.03 KB | 3507 | 1.53× |
| 4 | gopherlua | 565.99 µs | 28.09 KB | 1077 | 1.67× |
| 5 | risor | 575.50 µs | 355.80 KB | 2035 | 1.70× |
| 6 | golua | 672.93 µs | 63.69 KB | 6008 | 1.99× |
| 7 | goja | 1.00 ms | 134.12 KB | 6502 | 2.97× |

### array_dot

_Pre-allocate two 500-element arrays, fill via indexed loops, then compute the integer dot product via an indexed loop._

| Rank | Engine | ns/op | B/op | allocs/op | vs best |
|------|--------|-------|------|-----------|---------|
| 1 | kavun | 87.84 µs | 24.15 KB | 9 | 1.00× |
| 2 | gopherlua | 131.78 µs | 67.23 KB | 135 | 1.50× |
| 3 | starlark | 142.47 µs | 17.30 KB | 32 | 1.62× |
| 4 | golua | 146.30 µs | 72.27 KB | 4527 | 1.67× |
| 5 | tengo | 150.27 µs | 152.34 KB | 3527 | 1.71× |
| 6 | risor | 214.37 µs | 352.83 KB | 1890 | 2.44× |
| 7 | goja | 680.81 µs | 135.68 KB | 5160 | 7.75× |

