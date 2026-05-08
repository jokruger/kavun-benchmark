package task

// stringConcat stresses repeated string allocation and concatenation.
func stringConcat() Task {
	return Task{
		Name:        "string_concat",
		Description: "Build a string of 200 characters via repeated concatenation; return its length.",
		Expected:    200,
		Sources: map[string]string{
			EngineKavun: `
s = ""
for i = 0; i < 200; i++ { s += "x" }
__result = len(s)
`,
			EngineTengo: `
s := ""
for i := 0; i < 200; i++ { s += "x" }
__result = len(s)
`,
			EngineGoja: `
var s = "";
for (var i = 0; i < 200; i++) { s += "x"; }
s.length;
`,
			EngineGoLua: `
local s = ""
for i = 1, 200 do s = s .. "x" end
return #s
`,
			EngineGopher: `
local s = ""
for i = 1, 200 do s = s .. "x" end
return #s
`,
			EngineRisor: `
s := ""
for i := 0; i < 200; i++ { s += "x" }
len(s)
`,
		},
	}
}
