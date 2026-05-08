// Goja benchmarks. Naming scheme: BenchmarkGoja/<task> (parsed by cmd/report).
package goja_bench

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/jokruger/kavun-benchmark/task"
)

// gojaCompile prepares a Goja program for repeated execution. The returned runner re-runs the whole compiled program
// and returns the program's last-expression value.
func gojaCompile(b *testing.B, name, source string) func() any {
	b.Helper()
	prog, err := goja.Compile(name, source, false)
	if err != nil {
		b.Fatalf("goja compile: %v", err)
	}
	vm := goja.New()
	return func() any {
		v, err := vm.RunProgram(prog)
		if err != nil {
			b.Fatalf("goja run: %v", err)
		}
		if v == nil {
			return nil
		}
		return v.Export()
	}
}

func runGojaTask(b *testing.B, taskName, source string) {
	b.Helper()
	run := gojaCompile(b, taskName, source)
	if err := task.Validate(taskName, run()); err != nil {
		b.Fatalf("validate: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = run()
	}
}

func BenchmarkGoja(b *testing.B) {
	for _, t := range task.All {
		src, ok := t.Sources[task.EngineGoja]
		if !ok {
			continue
		}
		b.Run(t.Name, func(b *testing.B) {
			runGojaTask(b, t.Name, src)
		})
	}
}
