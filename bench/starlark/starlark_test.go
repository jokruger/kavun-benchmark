// Starlark benchmarks. Naming scheme: BenchmarkStarlark/<task> (parsed by cmd/report).
package starlark_bench

import (
	"testing"

	"github.com/jokruger/kavun-benchmark/task"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

// starlarkOptions enables the language features the benchmark scripts rely on:
//   - Recursion: needed for fib / fib_tail
//   - GlobalReassign + TopLevelControl: scripts execute as top-level statements that reassign module globals
//     (e.g. `s = 0; for i in range(...): s = s + i`)
var starlarkOptions = &syntax.FileOptions{
	Set:             true,
	While:           true,
	TopLevelControl: true,
	GlobalReassign:  true,
	Recursion:       true,
}

// starlarkCompile parses and compiles a Starlark program once. Each invocation of the returned runner calls
// program.Init on a fresh thread, which re-executes the top level and returns a new globals dictionary; the value of
// res` is then unwrapped to a Go native via toGo.
func starlarkCompile(b *testing.B, source string) func() any {
	b.Helper()
	_, prog, err := starlark.SourceProgramOptions(starlarkOptions, "bench.star", source, func(string) bool { return false })
	if err != nil {
		b.Fatalf("starlark compile: %v", err)
	}
	return func() any {
		thread := &starlark.Thread{Name: "bench"}
		globals, err := prog.Init(thread, nil)
		if err != nil {
			b.Fatalf("starlark run: %v", err)
		}
		return toGo(globals["res"])
	}
}

func toGo(v starlark.Value) any {
	switch x := v.(type) {
	case nil:
		return nil
	case starlark.Int:
		if i, ok := x.Int64(); ok {
			return i
		}
		return x.BigInt()
	case starlark.Float:
		return float64(x)
	case starlark.String:
		return string(x)
	case starlark.Bool:
		return bool(x)
	}
	return v
}

func runStarlarkTask(b *testing.B, taskName, source string) {
	b.Helper()
	run := starlarkCompile(b, source)
	if err := task.Validate(taskName, run()); err != nil {
		b.Fatalf("validate: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = run()
	}
}

func BenchmarkStarlark(b *testing.B) {
	for _, t := range task.All {
		src, ok := t.Sources[task.EngineStarlark]
		if !ok {
			continue
		}
		b.Run(t.Name, func(b *testing.B) {
			runStarlarkTask(b, t.Name, src)
		})
	}
}
