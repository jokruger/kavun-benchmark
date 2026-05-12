# Kavun Benchmark Report

Comparison of embeddable scripting engines for Go, focused on repeated execution of pre-compiled scripts.
Compilation cost is excluded from measurements.

## Host

| Field        | Value                                    |
| ------------ | ---------------------------------------- |
| OS           | darwin                                   |
| Arch         | amd64                                    |
| CPU          | Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz |
| Logical CPUs | 12                                       |
| GOMAXPROCS   | 12                                       |
| Total memory | 16.00 GB                                 |
| Go version   | go1.26.2                                 |
| Run at (UTC) | 2026-05-12T06:33:19Z                     |

## Summary

Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.
Lower is better; 1.00× means "fastest on every task".

Missing-task penalty: ratio = (worst observed ratio on the task) × 2.00, rank = last + 1.

| Rank | Engine    | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |
| ---- | --------- | ----------- | -------- | ----------- | ---- | ----------- | --------- | ------- |
| 1    | kavun0    | 1.04×       | 1.44     | 1.20×       | 5    | 1.20×       | 9         | 0       |
| 2    | kavun     | 1.06×       | 1.89     | 1.24×       | 3    | 1.04×       | 9         | 0       |
| 3    | gopherlua | 1.58×       | 3.56     | 4.00×       | 1    | 208.63×     | 9         | 0       |
| 4    | golua     | 1.63×       | 4.22     | 2.40×       | 0    | 291.49×     | 9         | 0       |
| 5    | starlark  | 2.55×       | 5.33     | 5.11×       | 0    | 202.66×     | 9         | 0       |
| 6    | tengo     | 3.17×       | 5.33     | 57.80×      | 0    | 1502.91×    | 9         | 0       |
| 7    | goja      | 5.23×       | 7.22     | 11.09×      | 0    | 379.43×     | 9         | 0       |
| 8    | risor     | 6.43×       | 7.00     | 180.82×     | 0    | 3958.94×    | 9         | 0       |

## Per-task results

### fib

_Naive recursive Fibonacci(25). Stresses function-call overhead and recursion._

| Rank | Engine    | ns/op    | B/op     | allocs/op | vs best |
| ---- | --------- | -------- | -------- | --------- | ------- |
| 1    | kavun0    | 14.07 ms | 8 B      | 1         | 1.00×   |
| 2    | kavun     | 14.26 ms | 8 B      | 1         | 1.01×   |
| 3    | tengo     | 17.43 ms | 2.51 MB  | 317814    | 1.24×   |
| 4    | golua     | 20.03 ms | 2.42 MB  | 317817    | 1.42×   |
| 5    | gopherlua | 21.38 ms | 10.52 KB | 42        | 1.52×   |
| 6    | goja      | 36.42 ms | 7.96 KB  | 392       | 2.59×   |
| 7    | risor     | 46.11 ms | 4.01 MB  | 243204    | 3.28×   |
| 8    | starlark  | 53.29 ms | 18.53 MB | 242828    | 3.79×   |

### fib_tail

_Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun0    | 1.58 µs  | 8 B       | 1         | 1.00×   |
| 2    | kavun     | 1.60 µs  | 8 B       | 1         | 1.01×   |
| 3    | golua     | 2.51 µs  | 544 B     | 47        | 1.59×   |
| 4    | gopherlua | 3.12 µs  | 2.89 KB   | 12        | 1.98×   |
| 5    | goja      | 5.56 µs  | 5.08 KB   | 23        | 3.52×   |
| 6    | starlark  | 8.06 µs  | 5.10 KB   | 58        | 5.11×   |
| 7    | tengo     | 17.59 µs | 88.48 KB  | 44        | 11.14×  |
| 8    | risor     | 55.13 µs | 309.24 KB | 74        | 34.92×  |

### loop_sum

_Sum of integers 1..10000. Stresses loops and integer arithmetic._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | gopherlua | 378.79 µs | 234.98 KB | 938       | 1.00×   |
| 2    | kavun0    | 396.53 µs | 8 B       | 1         | 1.05×   |
| 3    | kavun     | 470.21 µs | 8 B       | 1         | 1.24×   |
| 4    | golua     | 480.85 µs | 234.54 KB | 30006     | 1.27×   |
| 5    | starlark  | 627.45 µs | 1016 B    | 19        | 1.66×   |
| 6    | tengo     | 761.13 µs | 244.44 KB | 20006     | 2.01×   |
| 7    | risor     | 1.42 ms   | 615.82 KB | 19762     | 3.76×   |
| 8    | goja      | 4.18 ms   | 1.07 MB   | 39729     | 11.02×  |

### sum_pow

_Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication._

| Rank | Engine    | ns/op     | B/op       | allocs/op | vs best |
| ---- | --------- | --------- | ---------- | --------- | ------- |
| 1    | kavun0    | 468.39 µs | 8 B        | 1         | 1.00×   |
| 2    | gopherlua | 534.62 µs | 313.06 KB  | 1250      | 1.14×   |
| 3    | kavun     | 539.59 µs | 8 B        | 1         | 1.15×   |
| 4    | golua     | 660.66 µs | 312.63 KB  | 40002     | 1.41×   |
| 5    | tengo     | 1.04 ms   | 322.53 KB  | 30002     | 2.21×   |
| 6    | risor     | 1.89 ms   | 772.02 KB  | 29758     | 4.03×   |
| 7    | starlark  | 2.23 ms   | 1018.29 KB | 40711     | 4.76×   |
| 8    | goja      | 5.19 ms   | 1.14 MB    | 49723     | 11.09×  |

### closure_counter

_Closure that increments a captured counter, called 1000 times._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | kavun0    | 72.43 µs  | 120 B     | 4         | 1.00×   |
| 2    | kavun     | 79.23 µs  | 40 B      | 3         | 1.09×   |
| 3    | golua     | 106.61 µs | 23.69 KB  | 3011      | 1.47×   |
| 4    | gopherlua | 110.23 µs | 23.35 KB  | 94        | 1.52×   |
| 5    | tengo     | 119.31 µs | 103.92 KB | 2010      | 1.65×   |
| 6    | starlark  | 245.29 µs | 63.79 KB  | 1028      | 3.39×   |
| 7    | risor     | 294.13 µs | 331.79 KB | 1542      | 4.06×   |
| 8    | goja      | 452.12 µs | 107.18 KB | 3514      | 6.24×   |

### string_concat

_Build a string of 200 characters via repeated concatenation; return its length._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun     | 22.02 µs | 20.97 KB  | 199       | 1.00×   |
| 2    | kavun0    | 26.32 µs | 24.09 KB  | 399       | 1.20×   |
| 3    | gopherlua | 29.15 µs | 27.98 KB  | 413       | 1.32×   |
| 4    | starlark  | 29.31 µs | 25.12 KB  | 420       | 1.33×   |
| 5    | tengo     | 48.56 µs | 120.13 KB | 606       | 2.21×   |
| 6    | golua     | 52.85 µs | 39.85 KB  | 1403      | 2.40×   |
| 7    | risor     | 85.39 µs | 333.35 KB | 438       | 3.88×   |
| 8    | goja      | 92.43 µs | 43.05 KB  | 804       | 4.20×   |

### string_repeat

_Build a string of 200 'x' characters using the language's idiomatic repeat builtin; return its length._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun     | 280 ns   | 208 B     | 1         | 1.00×   |
| 2    | kavun0    | 303 ns   | 224 B     | 2         | 1.08×   |
| 3    | golua     | 538 ns   | 400 B     | 8         | 1.92×   |
| 4    | starlark  | 964 ns   | 1.07 KB   | 17        | 3.45×   |
| 5    | gopherlua | 1.12 µs  | 2.93 KB   | 11        | 4.00×   |
| 6    | goja      | 1.28 µs  | 416 B     | 8         | 4.59×   |
| 7    | tengo     | 16.17 µs | 88.48 KB  | 10        | 57.80×  |
| 8    | risor     | 50.57 µs | 307.85 KB | 42        | 180.82× |

### str_contains

_1000 substring searches over a 1000-char haystack, alternating an 8-char hit and miss needle. Stresses substring search._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | kavun     | 346.01 µs | 1.01 KB   | 2         | 1.00×   |
| 2    | kavun0    | 356.47 µs | 1.02 KB   | 3         | 1.03×   |
| 3    | starlark  | 440.30 µs | 2.15 KB   | 24        | 1.27×   |
| 4    | tengo     | 530.19 µs | 140.03 KB | 3507      | 1.53×   |
| 5    | gopherlua | 551.43 µs | 28.09 KB  | 1077      | 1.59×   |
| 6    | risor     | 583.05 µs | 355.80 KB | 2035      | 1.69×   |
| 7    | golua     | 659.95 µs | 63.69 KB  | 6008      | 1.91×   |
| 8    | goja      | 1.02 ms   | 134.12 KB | 6502      | 2.94×   |

### array_dot

_Pre-allocate two 500-element arrays, fill via indexed loops, then compute the integer dot product via an indexed loop._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | kavun0    | 95.03 µs  | 24.15 KB  | 9         | 1.00×   |
| 2    | kavun     | 99.56 µs  | 24.01 KB  | 3         | 1.05×   |
| 3    | gopherlua | 132.68 µs | 67.23 KB  | 135       | 1.40×   |
| 4    | golua     | 142.84 µs | 72.27 KB  | 4527      | 1.50×   |
| 5    | starlark  | 144.54 µs | 17.30 KB  | 32        | 1.52×   |
| 6    | tengo     | 154.45 µs | 152.34 KB | 3527      | 1.63×   |
| 7    | risor     | 216.17 µs | 352.83 KB | 1890      | 2.27×   |
| 8    | goja      | 699.96 µs | 135.68 KB | 5160      | 7.37×   |
