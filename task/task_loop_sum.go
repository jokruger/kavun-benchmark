package task

// loopSum stresses loops and integer arithmetic.
func loopSum() Task {
	return Task{
		Name:        "loop_sum",
		Description: "Sum of integers 1..10000. Stresses loops and integer arithmetic.",
		Expected:    50005000,
		Sources: map[string]string{
			EngineKavun: `
res = 0
for i := 1; i <= 10000; i++ { res += i }
`,
			EngineTengo: `
res = 0
for i := 1; i <= 10000; i++ { res += i }
`,
			EngineGoja: `
var res = 0;
for (var i = 1; i <= 10000; i++) { res += i; }
res;
`,
			EngineGoLua: `
local res = 0
for i = 1, 10000 do res = res + i end
return res
`,
			EngineGopher: `
local res = 0
for i = 1, 10000 do res = res + i end
return res
`,
			EngineRisor: `
res := 0
for i := 1; i <= 10000; i++ { res += i }
res
`,
			EngineStarlark: `
res = 0
for i in range(1, 10001):
    res = res + i
`,
		},
	}
}
