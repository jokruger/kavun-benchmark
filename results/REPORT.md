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
| Run at (UTC) | 2026-05-09T06:56:28Z                     |

## Summary

Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.
Lower is better; 1.00× means "fastest on every task".

Missing-task penalty: ratio = (worst observed ratio on the task) × 2.00, rank = last + 1.

| Rank | Engine    | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |
| ---- | --------- | ----------- | -------- | ----------- | ---- | ----------- | --------- | ------- |
| 1    | kavun0    | 1.03×       | 1.44     | 1.16×       | 5    | 1.20×       | 9         | 0       |
| 2    | kavun     | 1.09×       | 1.78     | 1.25×       | 3    | 1.04×       | 9         | 0       |
| 3    | gopherlua | 1.63×       | 3.78     | 3.98×       | 1    | 208.63×     | 9         | 0       |
| 4    | golua     | 1.68×       | 4.33     | 2.40×       | 0    | 291.49×     | 9         | 0       |
| 5    | starlark  | 2.62×       | 5.11     | 5.15×       | 0    | 202.66×     | 9         | 0       |
| 6    | tengo     | 3.30×       | 5.33     | 59.98×      | 0    | 1502.91×    | 9         | 0       |
| 7    | goja      | 5.35×       | 7.22     | 11.08×      | 0    | 379.43×     | 9         | 0       |
| 8    | risor     | 6.55×       | 7.00     | 180.74×     | 0    | 3958.94×    | 9         | 0       |

## Per-task results

### fib

_Naive recursive Fibonacci(25). Stresses function-call overhead and recursion._

| Rank | Engine    | ns/op    | B/op     | allocs/op | vs best |
| ---- | --------- | -------- | -------- | --------- | ------- |
| 1    | kavun0    | 11.83 ms | 8 B      | 1         | 1.00×   |
| 2    | kavun     | 11.88 ms | 8 B      | 1         | 1.00×   |
| 3    | tengo     | 16.75 ms | 2.51 MB  | 317814    | 1.42×   |
| 4    | golua     | 19.09 ms | 2.42 MB  | 317817    | 1.61×   |
| 5    | gopherlua | 19.93 ms | 10.52 KB | 42        | 1.69×   |
| 6    | goja      | 33.85 ms | 7.96 KB  | 392       | 2.86×   |
| 7    | risor     | 41.12 ms | 4.01 MB  | 243204    | 3.48×   |
| 8    | starlark  | 51.83 ms | 18.53 MB | 242828    | 4.38×   |

### fib_tail

_Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun0    | 1.53 µs  | 8 B       | 1         | 1.00×   |
| 2    | kavun     | 1.59 µs  | 8 B       | 1         | 1.04×   |
| 3    | golua     | 2.42 µs  | 544 B     | 47        | 1.58×   |
| 4    | gopherlua | 3.07 µs  | 2.89 KB   | 12        | 2.01×   |
| 5    | goja      | 5.45 µs  | 5.08 KB   | 23        | 3.57×   |
| 6    | starlark  | 7.86 µs  | 5.10 KB   | 58        | 5.15×   |
| 7    | tengo     | 17.27 µs | 88.48 KB  | 44        | 11.32×  |
| 8    | risor     | 52.81 µs | 309.24 KB | 74        | 34.61×  |

### loop_sum

_Sum of integers 1..10000. Stresses loops and integer arithmetic._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | gopherlua | 370.76 µs | 234.98 KB | 938       | 1.00×   |
| 2    | kavun0    | 380.68 µs | 8 B       | 1         | 1.03×   |
| 3    | kavun     | 464.90 µs | 8 B       | 1         | 1.25×   |
| 4    | golua     | 470.66 µs | 234.54 KB | 30006     | 1.27×   |
| 5    | starlark  | 606.22 µs | 1016 B    | 19        | 1.64×   |
| 6    | tengo     | 734.64 µs | 244.44 KB | 20006     | 1.98×   |
| 7    | risor     | 1.42 ms   | 615.82 KB | 19762     | 3.83×   |
| 8    | goja      | 4.01 ms   | 1.07 MB   | 39729     | 10.81×  |

### sum_pow

_Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication._

| Rank | Engine    | ns/op     | B/op       | allocs/op | vs best |
| ---- | --------- | --------- | ---------- | --------- | ------- |
| 1    | kavun0    | 445.97 µs | 8 B        | 1         | 1.00×   |
| 2    | kavun     | 512.02 µs | 8 B        | 1         | 1.15×   |
| 3    | gopherlua | 519.92 µs | 313.06 KB  | 1250      | 1.17×   |
| 4    | golua     | 641.61 µs | 312.63 KB  | 40002     | 1.44×   |
| 5    | tengo     | 1.01 ms   | 322.53 KB  | 30002     | 2.26×   |
| 6    | risor     | 1.84 ms   | 772.02 KB  | 29758     | 4.13×   |
| 7    | starlark  | 2.18 ms   | 1018.29 KB | 40711     | 4.88×   |
| 8    | goja      | 4.94 ms   | 1.14 MB    | 49723     | 11.08×  |

### closure_counter

_Closure that increments a captured counter, called 1000 times._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | kavun0    | 64.77 µs  | 120 B     | 4         | 1.00×   |
| 2    | kavun     | 80.74 µs  | 40 B      | 3         | 1.25×   |
| 3    | golua     | 103.32 µs | 23.69 KB  | 3011      | 1.60×   |
| 4    | gopherlua | 106.05 µs | 23.35 KB  | 94        | 1.64×   |
| 5    | tengo     | 117.83 µs | 103.92 KB | 2010      | 1.82×   |
| 6    | starlark  | 236.95 µs | 63.79 KB  | 1028      | 3.66×   |
| 7    | risor     | 267.12 µs | 331.79 KB | 1542      | 4.12×   |
| 8    | goja      | 436.05 µs | 107.18 KB | 3514      | 6.73×   |

### string_concat

_Build a string of 200 characters via repeated concatenation; return its length._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun     | 21.82 µs | 20.97 KB  | 199       | 1.00×   |
| 2    | kavun0    | 25.25 µs | 24.09 KB  | 399       | 1.16×   |
| 3    | starlark  | 28.64 µs | 25.12 KB  | 420       | 1.31×   |
| 4    | gopherlua | 28.93 µs | 27.98 KB  | 413       | 1.33×   |
| 5    | tengo     | 48.04 µs | 120.13 KB | 606       | 2.20×   |
| 6    | golua     | 52.41 µs | 39.85 KB  | 1403      | 2.40×   |
| 7    | risor     | 81.94 µs | 333.35 KB | 438       | 3.76×   |
| 8    | goja      | 91.14 µs | 43.05 KB  | 804       | 4.18×   |

### string_repeat

_Build a string of 200 'x' characters using the language's idiomatic repeat builtin; return its length._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun     | 270 ns   | 208 B     | 1         | 1.00×   |
| 2    | kavun0    | 289 ns   | 224 B     | 2         | 1.07×   |
| 3    | golua     | 526 ns   | 400 B     | 8         | 1.94×   |
| 4    | starlark  | 922 ns   | 1.07 KB   | 17        | 3.41×   |
| 5    | gopherlua | 1.08 µs  | 2.93 KB   | 11        | 3.98×   |
| 6    | goja      | 1.25 µs  | 416 B     | 8         | 4.64×   |
| 7    | tengo     | 16.22 µs | 88.48 KB  | 10        | 59.98×  |
| 8    | risor     | 48.88 µs | 307.85 KB | 42        | 180.74× |

### str_contains

_1000 substring searches over a 1000-char haystack, alternating an 8-char hit and miss needle. Stresses substring search._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | kavun     | 327.80 µs | 1.01 KB   | 2         | 1.00×   |
| 2    | kavun0    | 338.73 µs | 1.02 KB   | 3         | 1.03×   |
| 3    | starlark  | 418.55 µs | 2.15 KB   | 24        | 1.28×   |
| 4    | tengo     | 512.97 µs | 140.03 KB | 3507      | 1.56×   |
| 5    | gopherlua | 529.31 µs | 28.09 KB  | 1077      | 1.61×   |
| 6    | risor     | 569.14 µs | 355.80 KB | 2035      | 1.74×   |
| 7    | golua     | 629.44 µs | 63.69 KB  | 6008      | 1.92×   |
| 8    | goja      | 975.72 µs | 134.12 KB | 6502      | 2.98×   |

### array_dot

_Pre-allocate two 500-element arrays, fill via indexed loops, then compute the integer dot product via an indexed loop._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | kavun0    | 88.74 µs  | 24.15 KB  | 9         | 1.00×   |
| 2    | kavun     | 100.39 µs | 24.01 KB  | 3         | 1.13×   |
| 3    | gopherlua | 129.03 µs | 67.23 KB  | 135       | 1.45×   |
| 4    | starlark  | 135.07 µs | 17.30 KB  | 32        | 1.52×   |
| 5    | golua     | 139.12 µs | 72.27 KB  | 4527      | 1.57×   |
| 6    | tengo     | 150.64 µs | 152.34 KB | 3527      | 1.70×   |
| 7    | risor     | 212.50 µs | 352.83 KB | 1890      | 2.39×   |
| 8    | goja      | 675.43 µs | 135.68 KB | 5160      | 7.61×   |
