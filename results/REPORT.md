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
| Run at (UTC) | 2026-05-08T15:33:37Z                     |

## Summary

Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.
Lower is better; 1.00× means "fastest on every task".

Missing-task penalty: ratio = (worst observed ratio on the task) × 2.00, rank = last + 1.

| Rank | Engine    | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |
| ---- | --------- | ----------- | -------- | ----------- | ---- | ----------- | --------- | ------- |
| 1    | kavun0    | 1.07×       | 1.90     | 1.23×       | 3    | 1.28×       | 10        | 0       |
| 2    | kavun     | 1.11×       | 2.00     | 1.30×       | 4    | 1.00×       | 10        | 0       |
| 3    | gopherlua | 1.41×       | 3.00     | 2.75×       | 3    | 262.81×     | 10        | 0       |
| 4    | golua     | 1.56×       | 3.90     | 2.35×       | 0    | 415.74×     | 10        | 0       |
| 5    | tengo     | 2.33×       | 5.30     | 10.81×      | 0    | 1385.35×    | 10        | 0       |
| 6    | starlark  | 3.86×       | 6.10     | 29.23×      | 0    | 879.92×     | 8         | 2       |
| 7    | risor     | 4.41×       | 6.50     | 34.23×      | 0    | 3416.42×    | 10        | 0       |
| 8    | goja      | 6.34×       | 7.30     | 14.61×      | 0    | 806.58×     | 10        | 0       |

## Per-task results

### fib

_Naive recursive Fibonacci(25). Stresses function-call overhead and recursion._

| Rank | Engine    | ns/op    | B/op     | allocs/op | vs best |
| ---- | --------- | -------- | -------- | --------- | ------- |
| 1    | kavun     | 12.31 ms | 8 B      | 1         | 1.00×   |
| 2    | kavun0    | 12.58 ms | 8 B      | 1         | 1.02×   |
| 3    | tengo     | 17.11 ms | 2.51 MB  | 317814    | 1.39×   |
| 4    | golua     | 19.98 ms | 2.42 MB  | 317817    | 1.62×   |
| 5    | gopherlua | 21.30 ms | 10.53 KB | 42        | 1.73×   |
| 6    | goja      | 36.21 ms | 7.96 KB  | 392       | 2.94×   |
| 7    | risor     | 42.29 ms | 4.01 MB  | 243204    | 3.44×   |
| 8    | starlark  | 53.38 ms | 18.53 MB | 242828    | 4.34×   |

### fib_tail

_Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun     | 1.62 µs  | 8 B       | 1         | 1.00×   |
| 2    | kavun0    | 1.64 µs  | 8 B       | 1         | 1.01×   |
| 3    | golua     | 2.53 µs  | 544 B     | 47        | 1.56×   |
| 4    | gopherlua | 3.12 µs  | 2.89 KB   | 12        | 1.92×   |
| 5    | goja      | 5.75 µs  | 5.08 KB   | 23        | 3.55×   |
| 6    | starlark  | 8.02 µs  | 5.10 KB   | 58        | 4.95×   |
| 7    | tengo     | 17.52 µs | 88.48 KB  | 44        | 10.81×  |
| 8    | risor     | 55.46 µs | 309.24 KB | 74        | 34.23×  |

### loop_sum

_Sum of integers 1..10000. Stresses loops and integer arithmetic._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | gopherlua | 382.06 µs | 234.98 KB | 938       | 1.00×   |
| 2    | kavun0    | 407.29 µs | 8 B       | 1         | 1.07×   |
| 3    | kavun     | 474.82 µs | 8 B       | 1         | 1.24×   |
| 4    | golua     | 482.25 µs | 234.54 KB | 30006     | 1.26×   |
| 5    | starlark  | 621.32 µs | 1.01 KB   | 19        | 1.63×   |
| 6    | tengo     | 742.50 µs | 244.44 KB | 20006     | 1.94×   |
| 7    | risor     | 1.47 ms   | 615.82 KB | 19762     | 3.85×   |
| 8    | goja      | 4.53 ms   | 1.07 MB   | 39729     | 11.86×  |

### sum_pow

_Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication._

| Rank | Engine    | ns/op     | B/op       | allocs/op | vs best |
| ---- | --------- | --------- | ---------- | --------- | ------- |
| 1    | kavun0    | 480.14 µs | 8 B        | 1         | 1.00×   |
| 2    | gopherlua | 528.41 µs | 313.06 KB  | 1250      | 1.10×   |
| 3    | kavun     | 547.99 µs | 8 B        | 1         | 1.14×   |
| 4    | golua     | 657.95 µs | 312.63 KB  | 40002     | 1.37×   |
| 5    | tengo     | 1.02 ms   | 322.53 KB  | 30002     | 2.12×   |
| 6    | risor     | 1.92 ms   | 772.02 KB  | 29758     | 4.00×   |
| 7    | starlark  | 2.25 ms   | 1018.30 KB | 40711     | 4.68×   |
| 8    | goja      | 5.35 ms   | 1.14 MB    | 49723     | 11.15×  |

### closure_counter

_Closure that increments a captured counter, called 1000 times._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best          |
| ---- | --------- | --------- | --------- | --------- | ---------------- |
| 1    | kavun0    | 69.27 µs  | 120 B     | 4         | 1.00×            |
| 2    | kavun     | 81.95 µs  | 40 B      | 3         | 1.18×            |
| 3    | golua     | 108.51 µs | 23.69 KB  | 3011      | 1.57×            |
| 4    | gopherlua | 109.86 µs | 23.35 KB  | 94        | 1.59×            |
| 5    | tengo     | 117.22 µs | 103.92 KB | 2010      | 1.69×            |
| 6    | risor     | 276.95 µs | 331.79 KB | 1542      | 4.00×            |
| 7    | goja      | 451.51 µs | 107.18 KB | 3514      | 6.52×            |
| 8    | starlark  | _missing_ | _missing_ | _missing_ | 13.04× (penalty) |

### closures_iife

_Create-and-invoke a fresh closure 1000 times, accumulating the loop index into a captured variable._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best          |
| ---- | --------- | --------- | --------- | --------- | ---------------- |
| 1    | kavun0    | 62.05 µs  | 8 B       | 1         | 1.00×            |
| 2    | kavun     | 74.62 µs  | 8 B       | 1         | 1.20×            |
| 3    | golua     | 112.96 µs | 23.67 KB  | 3009      | 1.82×            |
| 4    | tengo     | 115.44 µs | 103.80 KB | 2005      | 1.86×            |
| 5    | gopherlua | 170.37 µs | 94.39 KB  | 2095      | 2.75×            |
| 6    | risor     | 311.17 µs | 350.48 KB | 2765      | 5.02×            |
| 7    | goja      | 906.77 µs | 763.78 KB | 10728     | 14.61×           |
| 8    | starlark  | _missing_ | _missing_ | _missing_ | 29.23× (penalty) |

### string_concat

_Build a string of 200 characters via repeated concatenation; return its length._

| Rank | Engine    | ns/op    | B/op      | allocs/op | vs best |
| ---- | --------- | -------- | --------- | --------- | ------- |
| 1    | kavun     | 22.85 µs | 20.97 KB  | 199       | 1.00×   |
| 2    | kavun0    | 25.73 µs | 24.09 KB  | 399       | 1.13×   |
| 3    | starlark  | 28.58 µs | 25.12 KB  | 420       | 1.25×   |
| 4    | gopherlua | 29.24 µs | 27.98 KB  | 413       | 1.28×   |
| 5    | tengo     | 48.47 µs | 120.13 KB | 606       | 2.12×   |
| 6    | golua     | 53.72 µs | 39.85 KB  | 1403      | 2.35×   |
| 7    | risor     | 85.63 µs | 333.35 KB | 438       | 3.75×   |
| 8    | goja      | 93.01 µs | 43.05 KB  | 804       | 4.07×   |

### str_contains

_Build a 100-char string of even-coded printable bytes, then probe contains() for each code in [32,232). Stresses substring search._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | kavun     | 78.89 µs  | 8.03 KB   | 399       | 1.00×   |
| 2    | kavun0    | 88.87 µs  | 14.28 KB  | 799       | 1.13×   |
| 3    | gopherlua | 122.56 µs | 19.61 KB  | 723       | 1.55×   |
| 4    | starlark  | 129.65 µs | 20.16 KB  | 1122      | 1.64×   |
| 5    | risor     | 174.57 µs | 336.34 KB | 1337      | 2.21×   |
| 6    | golua     | 176.12 µs | 54.91 KB  | 3006      | 2.23×   |
| 7    | tengo     | 179.44 µs | 142.69 KB | 3302      | 2.27×   |
| 8    | goja      | 320.38 µs | 59.76 KB  | 1863      | 4.06×   |

### array_sum

_Build an array of 0..499 then sum its elements._

| Rank | Engine    | ns/op     | B/op      | allocs/op | vs best |
| ---- | --------- | --------- | --------- | --------- | ------- |
| 1    | gopherlua | 79.69 µs  | 41.95 KB  | 100       | 1.00×   |
| 2    | kavun     | 89.09 µs  | 22.76 KB  | 4         | 1.12×   |
| 3    | kavun0    | 95.20 µs  | 43.10 KB  | 512       | 1.19×   |
| 4    | golua     | 101.98 µs | 45.99 KB  | 3518      | 1.28×   |
| 5    | starlark  | 158.52 µs | 58.59 KB  | 1537      | 1.99×   |
| 6    | tengo     | 183.77 µs | 157.37 KB | 3518      | 2.31×   |
| 7    | risor     | 258.85 µs | 403.80 KB | 3519      | 3.25×   |
| 8    | goja      | 564.72 µs | 111.05 KB | 4368      | 7.09×   |

### nested_loop

_100x100 nested loop with array index read+write per inner iteration._

| Rank | Engine    | ns/op   | B/op      | allocs/op | vs best |
| ---- | --------- | ------- | --------- | --------- | ------- |
| 1    | gopherlua | 1.02 ms | 158.78 KB | 621       | 1.00×   |
| 2    | golua     | 1.04 ms | 324.18 KB | 40916     | 1.03×   |
| 3    | kavun0    | 1.25 ms | 2.68 KB   | 4         | 1.23×   |
| 4    | kavun     | 1.32 ms | 2.63 KB   | 2         | 1.30×   |
| 5    | tengo     | 2.20 ms | 572.28 KB | 50717     | 2.17×   |
| 6    | risor     | 2.86 ms | 781.66 KB | 29787     | 2.82×   |
| 7    | starlark  | 2.93 ms | 1.04 MB   | 25173     | 2.89×   |
| 8    | goja      | 6.96 ms | 1.08 MB   | 39560     | 6.86×   |
