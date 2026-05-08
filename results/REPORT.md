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
| Run at (UTC) | 2026-05-08T14:49:21Z                     |

## Summary

Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.
Lower is better; 1.00× means "fastest on every task".

Missing-task penalty: ratio = (worst observed ratio on the task) × 2.00, rank = last + 1.

| Rank | Engine | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |
| ---- | ------ | ----------- | -------- | ----------- | ---- | ----------- | --------- | ------- |
| 1    | kavun0 | 1.05×       | 1.40     | 1.19×       | 6    | 1.28×       | 10        | 0       |
| 2    | kavun  | 1.10×       | 1.80     | 1.26×       | 3    | 1.00×       | 10        | 0       |
| 3    | golua  | 1.55×       | 3.00     | 2.39×       | 1    | 415.74×     | 10        | 0       |
| 4    | tengo  | 2.35×       | 3.90     | 11.62×      | 0    | 1385.35×    | 10        | 0       |
| 5    | goja   | 6.23×       | 4.90     | 14.64×      | 0    | 806.58×     | 10        | 0       |

## Per-task results

### fib

_Naive recursive Fibonacci(25). Stresses function-call overhead and recursion._

| Rank | Engine | ns/op    | B/op    | allocs/op | vs best |
| ---- | ------ | -------- | ------- | --------- | ------- |
| 1    | kavun0 | 12.63 ms | 8 B     | 1         | 1.00×   |
| 2    | kavun  | 13.46 ms | 8 B     | 1         | 1.07×   |
| 3    | tengo  | 18.15 ms | 2.51 MB | 317814    | 1.44×   |
| 4    | golua  | 20.40 ms | 2.42 MB | 317817    | 1.62×   |
| 5    | goja   | 36.31 ms | 7.96 KB | 392       | 2.87×   |

### fib_tail

_Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth._

| Rank | Engine | ns/op    | B/op     | allocs/op | vs best |
| ---- | ------ | -------- | -------- | --------- | ------- |
| 1    | kavun0 | 1.61 µs  | 8 B      | 1         | 1.00×   |
| 2    | kavun  | 1.67 µs  | 8 B      | 1         | 1.04×   |
| 3    | golua  | 2.54 µs  | 544 B    | 47        | 1.58×   |
| 4    | goja   | 5.41 µs  | 5.08 KB  | 23        | 3.36×   |
| 5    | tengo  | 18.70 µs | 88.48 KB | 44        | 11.62×  |

### loop_sum

_Sum of integers 1..10000. Stresses loops and integer arithmetic._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 406.73 µs | 8 B       | 1         | 1.00×   |
| 2    | kavun  | 474.65 µs | 8 B       | 1         | 1.17×   |
| 3    | golua  | 481.66 µs | 234.54 KB | 30006     | 1.18×   |
| 4    | tengo  | 745.73 µs | 244.44 KB | 20006     | 1.83×   |
| 5    | goja   | 4.59 ms   | 1.07 MB   | 39729     | 11.29×  |

### sum_pow

_Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 472.09 µs | 8 B       | 1         | 1.00×   |
| 2    | kavun  | 543.09 µs | 8 B       | 1         | 1.15×   |
| 3    | golua  | 660.49 µs | 312.63 KB | 40002     | 1.40×   |
| 4    | tengo  | 1.02 ms   | 322.53 KB | 30002     | 2.17×   |
| 5    | goja   | 5.41 ms   | 1.14 MB   | 49723     | 11.47×  |

### closure_counter

_Closure that increments a captured counter, called 1000 times._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 67.06 µs  | 120 B     | 4         | 1.00×   |
| 2    | kavun  | 78.85 µs  | 40 B      | 3         | 1.18×   |
| 3    | golua  | 108.16 µs | 23.69 KB  | 3011      | 1.61×   |
| 4    | tengo  | 117.56 µs | 103.92 KB | 2010      | 1.75×   |
| 5    | goja   | 455.86 µs | 107.18 KB | 3514      | 6.80×   |

### closures_iife

_Create-and-invoke a fresh closure 1000 times, accumulating the loop index into a captured variable._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 61.99 µs  | 8 B       | 1         | 1.00×   |
| 2    | kavun  | 71.23 µs  | 8 B       | 1         | 1.15×   |
| 3    | golua  | 114.19 µs | 23.67 KB  | 3009      | 1.84×   |
| 4    | tengo  | 117.46 µs | 103.80 KB | 2005      | 1.89×   |
| 5    | goja   | 907.52 µs | 763.78 KB | 10728     | 14.64×  |

### string_concat

_Build a string of 200 characters via repeated concatenation; return its length._

| Rank | Engine | ns/op    | B/op      | allocs/op | vs best |
| ---- | ------ | -------- | --------- | --------- | ------- |
| 1    | kavun  | 22.25 µs | 20.97 KB  | 199       | 1.00×   |
| 2    | kavun0 | 25.42 µs | 24.09 KB  | 399       | 1.14×   |
| 3    | tengo  | 50.08 µs | 120.13 KB | 606       | 2.25×   |
| 4    | golua  | 53.08 µs | 39.85 KB  | 1403      | 2.39×   |
| 5    | goja   | 91.77 µs | 43.05 KB  | 804       | 4.13×   |

### str_contains

_Build a 100-char string of even-coded printable bytes, then probe contains() for each code in [32,232). Stresses substring search._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun  | 78.45 µs  | 8.03 KB   | 399       | 1.00×   |
| 2    | kavun0 | 87.12 µs  | 14.28 KB  | 799       | 1.11×   |
| 3    | golua  | 179.79 µs | 54.91 KB  | 3006      | 2.29×   |
| 4    | tengo  | 179.82 µs | 142.69 KB | 3302      | 2.29×   |
| 5    | goja   | 315.31 µs | 59.76 KB  | 1863      | 4.02×   |

### array_sum

_Build an array of 0..499 then sum its elements._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun  | 86.45 µs  | 22.76 KB  | 4         | 1.00×   |
| 2    | kavun0 | 94.55 µs  | 43.10 KB  | 512       | 1.09×   |
| 3    | golua  | 103.88 µs | 45.99 KB  | 3518      | 1.20×   |
| 4    | tengo  | 185.53 µs | 157.37 KB | 3518      | 2.15×   |
| 5    | goja   | 562.28 µs | 111.05 KB | 4368      | 6.50×   |

### nested_loop

_100x100 nested loop with array index read+write per inner iteration._

| Rank | Engine | ns/op   | B/op      | allocs/op | vs best |
| ---- | ------ | ------- | --------- | --------- | ------- |
| 1    | golua  | 1.05 ms | 324.18 KB | 40916     | 1.00×   |
| 2    | kavun0 | 1.24 ms | 2.68 KB   | 4         | 1.19×   |
| 3    | kavun  | 1.32 ms | 2.63 KB   | 2         | 1.26×   |
| 4    | tengo  | 2.20 ms | 572.28 KB | 50717     | 2.10×   |
| 5    | goja   | 6.91 ms | 1.08 MB   | 39560     | 6.60×   |
