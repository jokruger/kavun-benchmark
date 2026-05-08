package task

// arraySum stresses array creation, indexed access, and iteration.
func arraySum() Task {
	return Task{
		Name:        "array_sum",
		Description: "Build an array of 0..499 then sum its elements.",
		Expected:    124750,
		Sources: map[string]string{
			EngineKavun: `
a = []
for i = 0; i < 500; i++ { a = append(a, i) }
s = 0
for i = 0; i < len(a); i++ { s += a[i] }
__result = s
`,
			EngineTengo: `
a := []
for i := 0; i < 500; i++ { a = append(a, i) }
s := 0
for i := 0; i < len(a); i++ { s += a[i] }
__result = s
`,
			EngineGoja: `
var a = [];
for (var i = 0; i < 500; i++) { a.push(i); }
var s = 0;
for (var i = 0; i < a.length; i++) { s += a[i]; }
s;
`,
			EngineGoLua: `
local a = {}
for i = 0, 499 do a[#a+1] = i end
local s = 0
for i = 1, #a do s = s + a[i] end
return s
`,
		},
	}
}
