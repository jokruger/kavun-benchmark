// Kavun benchmarks. Naming scheme: BenchmarkKavun/<task> (parsed by cmd/report).
package kavun_bench

import (
	"testing"

	"github.com/jokruger/kavun"
	"github.com/jokruger/kavun-benchmark/task"
	"github.com/jokruger/kavun/core"
	"github.com/jokruger/kavun/stdlib"
	"github.com/jokruger/kavun/vm"
)

// kavunCompile prepares a Kavun script for repeated execution. The runner re-runs the compiled bytecode and returns
// the value stored in the script's __result global.
func kavunCompile(b *testing.B, source string) func() any {
	b.Helper()
	s := kavun.NewScript([]byte(source))
	s.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	s.Add("__result", core.Undefined)
	cta := core.NewArena(nil)
	rta := core.NewArena(nil)
	machine := vm.NewVM(vm.DefaultMaxFrames, vm.DefaultStackSize)
	compiled, err := s.Compile(cta)
	if err != nil {
		b.Fatalf("kavun compile: %v", err)
	}
	return func() any {
		if err := compiled.Run(rta, machine); err != nil {
			b.Fatalf("kavun run: %v", err)
		}
		v := compiled.GetValue("__result")
		if n, ok := v.AsInt(); ok {
			return n
		}
		if str, ok := v.AsString(); ok {
			return str
		}
		if f, ok := v.AsFloat(); ok {
			return f
		}
		return nil
	}
}

func runKavunTask(b *testing.B, taskName, source string) {
	b.Helper()
	run := kavunCompile(b, source)
	if err := task.Validate(taskName, run()); err != nil {
		b.Fatalf("validate: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = run()
	}
}

func BenchmarkKavun(b *testing.B) {
	for _, t := range task.All {
		src, ok := t.Sources[task.EngineKavun]
		if !ok {
			continue
		}
		b.Run(t.Name, func(b *testing.B) {
			runKavunTask(b, t.Name, src)
		})
	}
}
