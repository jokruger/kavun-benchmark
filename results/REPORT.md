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
| Run at (UTC) | 2026-05-08T08:24:44Z                     |

## Summary

Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.
Lower is better; 1.00× means "fastest on every task".

Missing-task penalty: ratio = (worst observed ratio on the task) × 2.00, rank = last + 1.

| Rank | Engine | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |
| ---- | ------ | ----------- | -------- | ----------- | ---- | ----------- | --------- | ------- |
| 1    | kavun0 | 1.04×       | 1.40     | 1.15×       | 6    | 1.28×       | 10        | 0       |
| 2    | kavun  | 1.06×       | 1.60     | 1.17×       | 4    | 1.00×       | 10        | 0       |
| 3    | tengo  | 2.33×       | 3.10     | 11.18×      | 0    | 1385.35×    | 10        | 0       |
| 4    | goja   | 6.15×       | 3.90     | 15.48×      | 0    | 806.58×     | 10        | 0       |

## Per-task results

### fib

_Naive recursive Fibonacci(25). Stresses function-call overhead and recursion._

| Rank | Engine | ns/op    | B/op    | allocs/op | vs best |
| ---- | ------ | -------- | ------- | --------- | ------- |
| 1    | kavun  | 11.69 ms | 8 B     | 1         | 1.00×   |
| 2    | kavun0 | 11.84 ms | 8 B     | 1         | 1.01×   |
| 3    | tengo  | 16.52 ms | 2.51 MB | 317814    | 1.41×   |
| 4    | goja   | 33.58 ms | 7.96 KB | 392       | 2.87×   |

### fib_tail

_Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth._

| Rank | Engine | ns/op    | B/op     | allocs/op | vs best |
| ---- | ------ | -------- | -------- | --------- | ------- |
| 1    | kavun0 | 1.52 µs  | 8 B      | 1         | 1.00×   |
| 2    | kavun  | 1.56 µs  | 8 B      | 1         | 1.02×   |
| 3    | goja   | 5.34 µs  | 5.08 KB  | 23        | 3.50×   |
| 4    | tengo  | 17.06 µs | 88.48 KB | 44        | 11.18×  |

### loop_sum

_Sum of integers 1..10000. Stresses loops and integer arithmetic._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 386.45 µs | 8 B       | 1         | 1.00×   |
| 2    | kavun  | 442.09 µs | 8 B       | 1         | 1.14×   |
| 3    | tengo  | 718.44 µs | 244.44 KB | 20006     | 1.86×   |
| 4    | goja   | 4.22 ms   | 1.07 MB   | 39729     | 10.92×  |

### sum_pow

_Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 447.72 µs | 8 B       | 1         | 1.00×   |
| 2    | kavun  | 510.78 µs | 8 B       | 1         | 1.14×   |
| 3    | tengo  | 987.31 µs | 322.53 KB | 30002     | 2.21×   |
| 4    | goja   | 5.12 ms   | 1.14 MB   | 49723     | 11.43×  |

### closure_counter

_Closure that increments a captured counter, called 1000 times._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 63.22 µs  | 120 B     | 4         | 1.00×   |
| 2    | kavun  | 73.67 µs  | 40 B      | 3         | 1.17×   |
| 3    | tengo  | 116.40 µs | 103.92 KB | 2010      | 1.84×   |
| 4    | goja   | 429.20 µs | 107.18 KB | 3514      | 6.79×   |

### closures_iife

_Create-and-invoke a fresh closure 1000 times, accumulating the loop index into a captured variable._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun0 | 57.44 µs  | 8 B       | 1         | 1.00×   |
| 2    | kavun  | 66.64 µs  | 8 B       | 1         | 1.16×   |
| 3    | tengo  | 114.76 µs | 103.80 KB | 2005      | 2.00×   |
| 4    | goja   | 889.21 µs | 763.78 KB | 10728     | 15.48×  |

### string_concat

_Build a string of 200 characters via repeated concatenation; return its length._

| Rank | Engine | ns/op    | B/op      | allocs/op | vs best |
| ---- | ------ | -------- | --------- | --------- | ------- |
| 1    | kavun  | 21.88 µs | 20.97 KB  | 199       | 1.00×   |
| 2    | kavun0 | 25.27 µs | 24.09 KB  | 399       | 1.15×   |
| 3    | tengo  | 47.68 µs | 120.13 KB | 606       | 2.18×   |
| 4    | goja   | 89.82 µs | 43.05 KB  | 804       | 4.10×   |

### str_contains

_Build a 100-char string of even-coded printable bytes, then probe contains() for each code in [32,232). Stresses substring search._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun  | 74.20 µs  | 8.03 KB   | 399       | 1.00×   |
| 2    | kavun0 | 83.21 µs  | 14.28 KB  | 799       | 1.12×   |
| 3    | tengo  | 176.08 µs | 142.69 KB | 3302      | 2.37×   |
| 4    | goja   | 300.84 µs | 59.76 KB  | 1863      | 4.05×   |

### array_sum

_Build an array of 0..499 then sum its elements._

| Rank | Engine | ns/op     | B/op      | allocs/op | vs best |
| ---- | ------ | --------- | --------- | --------- | ------- |
| 1    | kavun  | 83.56 µs  | 22.76 KB  | 4         | 1.00×   |
| 2    | kavun0 | 91.96 µs  | 43.10 KB  | 512       | 1.10×   |
| 3    | tengo  | 179.81 µs | 157.37 KB | 3518      | 2.15×   |
| 4    | goja   | 536.33 µs | 111.05 KB | 4368      | 6.42×   |

### nested_loop

_100x100 nested loop with array index read+write per inner iteration._

| Rank | Engine | ns/op   | B/op      | allocs/op | vs best |
| ---- | ------ | ------- | --------- | --------- | ------- |
| 1    | kavun0 | 1.19 ms | 2.68 KB   | 4         | 1.00×   |
| 2    | kavun  | 1.24 ms | 2.63 KB   | 2         | 1.04×   |
| 3    | tengo  | 2.12 ms | 572.28 KB | 50717     | 1.78×   |
| 4    | goja   | 6.56 ms | 1.08 MB   | 39560     | 5.51×   |
