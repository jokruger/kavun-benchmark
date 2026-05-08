package task

// closuresIIFE stresses repeated closure *creation* (in addition to invocation): each iteration constructs a fresh
// anonymous function that captures the outer `out` variable and is immediately invoked. Contrast with closure_counter,
// which builds the closure once and only calls it many times.
func closuresIIFE() Task {
	return Task{
		Name:        "closures_iife",
		Description: "Create-and-invoke a fresh closure 1000 times, accumulating the loop index into a captured variable.",
		Expected:    499500,
		Sources: map[string]string{
			EngineKavun: `
__result = 0
for i = 0; i < 1000; i++ {
  func(x) { __result += x }(i)
}
`,
			EngineTengo: `
__result = 0
for i := 0; i < 1000; i++ {
  func(x) { __result += x }(i)
}
`,
			EngineGoja: `
var out = 0;
for (var i = 0; i < 1000; i++) {
  (function(x) { out += x; })(i);
}
out;
`,
			EngineGoLua: `
local out = 0
for i = 0, 999 do
  (function(x) out = out + x end)(i)
end
return out
`,
		},
	}
}
