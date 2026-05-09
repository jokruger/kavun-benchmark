package task

// arrayDot stresses indexed array access in a loop with realistic per-iteration work: a dot product
// of two integer vectors. Each engine pre-allocates two 500-element arrays (using its idiomatic
// primitive), fills them via indexed loops (a[i] = i, b[i] = 2*i), then computes
// res += a[i] * b[i] in a tight loop. Result: 2 * sum(i^2 for i in 0..499) = 83_083_500.
func arrayDot() Task {
	return Task{
		Name:        "array_dot",
		Description: "Pre-allocate two 500-element arrays, fill via indexed loops, then compute the integer dot product via an indexed loop.",
		Expected:    83083500,
		Sources: map[string]string{
			EngineKavun: `
a = [0].repeat(500)
b = [0].repeat(500)
for i = 0; i < 500; i++ { a[i] = i; b[i] = 2*i }
res = 0
for i = 0; i < 500; i++ { res += a[i] * b[i] }
`,
			EngineTengo: `
a := range(0, 500)
b := range(0, 500)
for i := 0; i < 500; i++ { a[i] = i; b[i] = 2*i }
res = 0
for i := 0; i < 500; i++ { res += a[i] * b[i] }
`,
			EngineGoja: `
var a = new Array(500);
var b = new Array(500);
for (var i = 0; i < 500; i++) { a[i] = i; b[i] = 2*i; }
var res = 0;
for (var i = 0; i < 500; i++) { res += a[i] * b[i]; }
res;
`,
			EngineGoLua: `
local a = {}
local b = {}
for i = 1, 500 do a[i] = i - 1; b[i] = 2 * (i - 1) end
local res = 0
for i = 1, 500 do res = res + a[i] * b[i] end
return res
`,
			EngineGopher: `
local a = {}
local b = {}
for i = 1, 500 do a[i] = i - 1; b[i] = 2 * (i - 1) end
local res = 0
for i = 1, 500 do res = res + a[i] * b[i] end
return res
`,
			EngineRisor: `
a := list(500)
b := list(500)
for i := 0; i < 500; i++ { a[i] = i; b[i] = 2*i }
res := 0
for i := 0; i < 500; i++ { res += a[i] * b[i] }
res
`,
			EngineStarlark: `
a = [0] * 500
b = [0] * 500
for i in range(500):
    a[i] = i
    b[i] = 2 * i
res = 0
for i in range(500):
    res = res + a[i] * b[i]
`,
		},
	}
}
