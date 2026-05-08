// go-lua benchmarks. Naming scheme: BenchmarkGoLua/<task> (parsed by cmd/report).
package golua_bench

import (
	"testing"

	lua "github.com/Shopify/go-lua"
	"github.com/jokruger/kavun-benchmark/task"
)

// goluaCompile prepares a Lua chunk for repeated execution. The compiled chunk is loaded once and left on the Lua
// stack; each invocation duplicates it (PushValue) and ProtectedCalls it, returning the chunk's single return value.
func goluaCompile(b *testing.B, source string) func() any {
	b.Helper()
	l := lua.NewState()
	lua.OpenLibraries(l)
	if err := lua.LoadString(l, source); err != nil {
		b.Fatalf("go-lua load: %v", err)
	}
	chunkIdx := l.Top()
	return func() any {
		l.PushValue(chunkIdx)
		if err := l.ProtectedCall(0, 1, 0); err != nil {
			b.Fatalf("go-lua run: %v", err)
		}
		var r any
		if l.IsNumber(-1) {
			f, _ := l.ToNumber(-1)
			r = f
		} else if s, ok := l.ToString(-1); ok {
			r = s
		} else if l.IsBoolean(-1) {
			r = l.ToBoolean(-1)
		}
		l.Pop(1)
		return r
	}
}

func runGoLuaTask(b *testing.B, taskName, source string) {
	b.Helper()
	run := goluaCompile(b, source)
	if err := task.Validate(taskName, run()); err != nil {
		b.Fatalf("validate: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = run()
	}
}

func BenchmarkGoLua(b *testing.B) {
	for _, t := range task.All {
		src, ok := t.Sources[task.EngineGoLua]
		if !ok {
			continue
		}
		b.Run(t.Name, func(b *testing.B) {
			runGoLuaTask(b, t.Name, src)
		})
	}
}
