package task

// sumPow stresses tight-loop integer arithmetic with a multiplication on the hot path.
// Sum of squares i^2 for i in [1, 10000) = 9999*10000*19999/6 = 333_283_335_000.
func sumPow() Task {
	return Task{
		Name:        "sum_pow",
		Description: "Sum of squares 1..9999 in a tight loop. Stresses loops plus integer multiplication.",
		Expected:    333283335000,
		Sources: map[string]string{
			EngineKavun: `
s = 0
for i = 1; i < 10000; i++ { s += i * i }
__result = s
`,
			EngineTengo: `
s := 0
for i := 1; i < 10000; i++ { s += i * i }
__result = s
`,
			EngineGoja: `
var s = 0;
for (var i = 1; i < 10000; i++) { s += i * i; }
s;
`,
			EngineGoLua: `
local s = 0
for i = 1, 9999 do s = s + i * i end
return s
`,
			EngineGopher: `
local s = 0
for i = 1, 9999 do s = s + i * i end
return s
`,
		},
	}
}
