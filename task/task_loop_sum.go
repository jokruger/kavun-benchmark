package task

// loopSum stresses loops and integer arithmetic.
func loopSum() Task {
	return Task{
		Name:        "loop_sum",
		Description: "Sum of integers 1..10000. Stresses loops and integer arithmetic.",
		Expected:    50005000,
		Sources: map[string]string{
			EngineKavun: `
s = 0
for i = 1; i <= 10000; i++ { s += i }
__result = s
`,
			EngineTengo: `
s := 0
for i := 1; i <= 10000; i++ { s += i }
__result = s
`,
			EngineGoja: `
var s = 0;
for (var i = 1; i <= 10000; i++) { s += i; }
s;
`,
			EngineGoLua: `
local s = 0
for i = 1, 10000 do s = s + i end
return s
`,
			EngineGopher: `
local s = 0
for i = 1, 10000 do s = s + i end
return s
`,
		},
	}
}
