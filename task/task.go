// Package task is the registry of benchmark workloads AND the single home for every script source.
package task

import "fmt"

// Engine identifiers
const (
	EngineKavun = "kavun"
	EngineTengo = "tengo"
	EngineGoja  = "goja"
	EngineGoLua = "golua"
)

// Task is a single canonical workload.
type Task struct {
	Name        string
	Description string
	Expected    any
	Sources     map[string]string
}

// All tasks recognized by the benchmark suite, in canonical report order.
var All = []Task{
	fib(),
	fibTail(),
	loopSum(),
	sumPow(),
	closureCounter(),
	closuresIIFE(),
	stringConcat(),
	strContains(),
	arraySum(),
	nestedLoop(),
}

// Names returns the canonical task names in registry order.
func Names() []string {
	out := make([]string, len(All))
	for i, t := range All {
		out[i] = t.Name
	}
	return out
}

// Lookup returns the canonical Task for the given name, or false.
func Lookup(name string) (Task, bool) {
	for _, t := range All {
		if t.Name == name {
			return t, true
		}
	}
	return Task{}, false
}

// Source returns the script source for (task, engine), or false if the engine does not implement the task.
func Source(taskName, engine string) (string, bool) {
	t, ok := Lookup(taskName)
	if !ok {
		return "", false
	}
	s, ok := t.Sources[engine]
	return s, ok
}

// EqualNumeric compares two values that may have different concrete numeric types across engines. Falls back to
// fmt.Sprint equality for non-numerics.
func EqualNumeric(got, want any) bool {
	gf, gok := toFloat(got)
	wf, wok := toFloat(want)
	if gok && wok {
		return gf == wf
	}
	return fmt.Sprint(got) == fmt.Sprint(want)
}

// Validate returns nil if got matches the expected value for the named task.
func Validate(name string, got any) error {
	t, ok := Lookup(name)
	if !ok {
		return fmt.Errorf("unknown task %q", name)
	}
	if !EqualNumeric(got, t.Expected) {
		return fmt.Errorf("task %q: got %v (%T), want %v", name, got, got, t.Expected)
	}
	return nil
}

func toFloat(v any) (float64, bool) {
	switch n := v.(type) {
	case int:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case uint:
		return float64(n), true
	case uint32:
		return float64(n), true
	case uint64:
		return float64(n), true
	case float32:
		return float64(n), true
	case float64:
		return n, true
	}
	return 0, false
}
