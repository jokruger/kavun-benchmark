package task

// stringRepeat exercises each engine's idiomatic string-building primitive (repeat / rep / `*`) rather than the
// loop-concat path measured by string_concat. This stresses the C-side fast path and shows what users actually pay
// when they use the natural form for the language.
func stringRepeat() Task {
	return Task{
		Name:        "string_repeat",
		Description: "Build a string of 200 'x' characters using the language's idiomatic repeat builtin; return its length.",
		Expected:    200,
		Sources: map[string]string{
			EngineKavun: `
s = "x".repeat(200)
res = len(s)
`,
			EngineTengo: `
text := import("text")
s := text.repeat("x", 200)
res = len(s)
`,
			EngineGoja: `
var s = "x".repeat(200)
s.length;
`,
			EngineGoLua: `
local s = string.rep("x", 200)
return #s
`,
			EngineGopher: `
local s = string.rep("x", 200)
return #s
`,
			EngineRisor: `
s := strings.repeat("x", 200)
len(s)
`,
			EngineStarlark: `
s = "x" * 200
res = len(s)
`,
		},
	}
}
