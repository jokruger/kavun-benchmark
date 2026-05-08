// gopher-lua benchmarks. Naming scheme: BenchmarkGopherLua/<task> (parsed by cmd/report).
package gopherlua_bench

import (
	"testing"

	lua "github.com/yuin/gopher-lua"
	"github.com/jokruger/kavun-benchmark/task"
)

// gopherCompile prepares a Lua chunk for repeated execution. The compiled *LFunction is reused for every iteration:
// each call pushes it on the stack and PCalls it, returning the chunk's single return value.
func gopherCompile(b *testing.B, source string) func() any {
	b.Helper()
	L := lua.NewState()
	fn, err := L.LoadString(source)
	if err != nil {
		b.Fatalf("gopher-lua load: %v", err)
	}
	return func() any {
		L.Push(fn)
		if err := L.PCall(0, 1, nil); err != nil {
			b.Fatalf("gopher-lua run: %v", err)
		}
		v := L.Get(-1)
		L.Pop(1)
		switch x := v.(type) {
		case lua.LNumber:
			return float64(x)
		case lua.LString:
			return string(x)
		case lua.LBool:
			return bool(x)
		case *lua.LNilType:
			return nil
		}
		return v
	}
}

func runGopherTask(b *testing.B, taskName, source string) {
	b.Helper()
	run := gopherCompile(b, source)
	if err := task.Validate(taskName, run()); err != nil {
		b.Fatalf("validate: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = run()
	}
}

func BenchmarkGopherLua(b *testing.B) {
	for _, t := range task.All {
		src, ok := t.Sources[task.EngineGopher]
		if !ok {
			continue
		}
		b.Run(t.Name, func(b *testing.B) {
			runGopherTask(b, t.Name, src)
		})
	}
}
