// Kavun (0 pre-allocs) benchmarks. Naming scheme: BenchmarkKavun/<task> (parsed by cmd/report).
package kavun_bench

import (
	"testing"

	"github.com/jokruger/kavun"
	"github.com/jokruger/kavun-benchmark/task"
	"github.com/jokruger/kavun/core"
	"github.com/jokruger/kavun/stdlib"
	"github.com/jokruger/kavun/vm"
)

func allocOptions() *core.ArenaOptions {
	opts := core.DefaultArenaOptions()
	opts.Decimals = 0
	opts.Times = 0
	opts.BytesNum = 0
	opts.RunesNum = 0
	opts.ArraysNum = 0
	opts.BuiltinFunctions = 0
	opts.CompiledFunctions = 0
	opts.ErrorValues = 0
	opts.StringValues = 0
	opts.RunesValues = 0
	opts.BytesValues = 0
	opts.ArrayValues = 0
	opts.DictValues = 0
	opts.IntRangeValues = 0
	opts.RunesIterators = 0
	opts.BytesIterators = 0
	opts.ArrayIterators = 0
	opts.DictIterators = 0
	opts.IntRangeIterators = 0
	return opts
}

// kavunCompile prepares a Kavun script for repeated execution. The runner re-runs the compiled bytecode and returns
// the value stored in the script's res global.
func kavunCompile(b *testing.B, source string) func() any {
	b.Helper()
	s := kavun.NewScript([]byte(source))
	s.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	s.Add("res", core.Undefined)
	cta := core.NewArena(nil)
	rta := core.NewArena(allocOptions())
	machine := vm.NewVM(vm.DefaultMaxFrames, vm.DefaultStackSize)
	compiled, err := s.Compile(cta)
	if err != nil {
		b.Fatalf("kavun compile: %v", err)
	}
	return func() any {
		if err := compiled.Run(rta, machine); err != nil {
			b.Fatalf("kavun run: %v", err)
		}
		v := compiled.GetValue("res")
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

func BenchmarkKavun0(b *testing.B) {
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
