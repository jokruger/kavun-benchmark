// Risor benchmarks. Naming scheme: BenchmarkRisor/<task> (parsed by cmd/report).
package risor_bench

import (
	"context"
	"testing"

	"github.com/risor-io/risor"
	"github.com/risor-io/risor/compiler"
	"github.com/risor-io/risor/parser"
	"github.com/risor-io/risor/vm"
	"github.com/jokruger/kavun-benchmark/task"
)

// risorCompile parses + compiles the script once with Risor's default config (so the standard module globals like
// `strings` and builtins like `chr` are available), then returns a runner that invokes vm.Run on the compiled bytecode
// and surfaces the top-of-stack value as a Go native via Object.Interface().
func risorCompile(b *testing.B, source string) func() any {
	b.Helper()
	ctx := context.Background()
	cfg := risor.NewConfig()
	ast, err := parser.Parse(ctx, source)
	if err != nil {
		b.Fatalf("risor parse: %v", err)
	}
	code, err := compiler.Compile(ast, cfg.CompilerOpts()...)
	if err != nil {
		b.Fatalf("risor compile: %v", err)
	}
	vmOpts := cfg.VMOpts()
	return func() any {
		result, err := vm.Run(ctx, code, vmOpts...)
		if err != nil {
			b.Fatalf("risor run: %v", err)
		}
		if result == nil {
			return nil
		}
		return result.Interface()
	}
}

func runRisorTask(b *testing.B, taskName, source string) {
	b.Helper()
	run := risorCompile(b, source)
	if err := task.Validate(taskName, run()); err != nil {
		b.Fatalf("validate: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = run()
	}
}

func BenchmarkRisor(b *testing.B) {
	for _, t := range task.All {
		src, ok := t.Sources[task.EngineRisor]
		if !ok {
			continue
		}
		b.Run(t.Name, func(b *testing.B) {
			runRisorTask(b, t.Name, src)
		})
	}
}
