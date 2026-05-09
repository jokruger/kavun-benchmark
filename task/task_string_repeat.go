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
text = import("text")
s = text.repeat("x", 200)
res = len(s)
`,
			EngineTengo: `
text := import("text")
s := text.repeat("x", 200)
res = len(s)
`,
			EngineGoja: `
"x".repeat(200).length;
`,
			EngineGoLua: `
return #string.rep("x", 200)
`,
			EngineGopher: `
return #string.rep("x", 200)
`,
			EngineRisor: `
len(strings.repeat("x", 200))
`,
			EngineStarlark: `
res = len("x" * 200)
`,
		},
	}
}
