// Tengo benchmarks.
// Naming scheme: BenchmarkTengo/<task> (parsed by cmd/report).
package tengo_bench

import (
	"context"
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/jokruger/kavun-benchmark/task"
)

// tengoCompile prepares a Tengo script for repeated execution. Returned runner performs one full execution and returns
// the value stored in the script's res global.
func tengoCompile(b *testing.B, source string) func() any {
	b.Helper()
	s := tengo.NewScript([]byte(source))
	s.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	if err := s.Add("res", int64(0)); err != nil {
		b.Fatalf("tengo add global: %v", err)
	}
	c, err := s.Compile()
	if err != nil {
		b.Fatalf("tengo compile: %v", err)
	}
	ctx := context.Background()
	return func() any {
		if err := c.RunContext(ctx); err != nil {
			b.Fatalf("tengo run: %v", err)
		}
		v := c.Get("res")
		if v == nil {
			return nil
		}
		return v.Value()
	}
}

func runTengoTask(b *testing.B, taskName, source string) {
	b.Helper()
	run := tengoCompile(b, source)
	if err := task.Validate(taskName, run()); err != nil {
		b.Fatalf("validate: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = run()
	}
}

func BenchmarkTengo(b *testing.B) {
	for _, t := range task.All {
		src, ok := t.Sources[task.EngineTengo]
		if !ok {
			continue
		}
		b.Run(t.Name, func(b *testing.B) {
			runTengoTask(b, t.Name, src)
		})
	}
}
