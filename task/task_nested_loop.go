package task

// nestedLoop stresses indexed array access and in-place mutation inside a hot nested loop. The reported value is just
// the iteration count (100*100 = 10000) so engines aren't penalized for differing integer overflow semantics on the
// growing array elements; the array mutation is there as work, not as the result.
func nestedLoop() Task {
	return Task{
		Name:        "nested_loop",
		Description: "100x100 nested loop with array index read+write per inner iteration.",
		Expected:    10000,
		Sources: map[string]string{
			EngineKavun: `
s = range(0, 100, 1).array()
n = 0
for i = 0; i < len(s); i++ {
  for j = 0; j < len(s); j++ {
    s[j] += s[i]
    n += 1
  }
}
__result = n
`,
			EngineTengo: `
s := []
for i := 0; i < 100; i++ { s = append(s, i) }
n := 0
for i := 0; i < len(s); i++ {
  for j := 0; j < len(s); j++ {
    s[j] += s[i]
    n += 1
  }
}
__result = n
`,
			EngineGoja: `
var s = [];
for (var i = 0; i < 100; i++) { s.push(i); }
var n = 0;
for (var i = 0; i < s.length; i++) {
  for (var j = 0; j < s.length; j++) {
    s[j] += s[i];
    n += 1;
  }
}
n;
`,
			EngineGoLua: `
local s = {}
for i = 0, 99 do s[#s+1] = i end
local n = 0
for i = 1, #s do
  for j = 1, #s do
    s[j] = s[j] + s[i]
    n = n + 1
  end
end
return n
`,
			EngineGopher: `
local s = {}
for i = 0, 99 do s[#s+1] = i end
local n = 0
for i = 1, #s do
  for j = 1, #s do
    s[j] = s[j] + s[i]
    n = n + 1
  end
end
return n
`,
			EngineRisor: `
s := []
for i := 0; i < 100; i++ { s.append(i) }
n := 0
for i := 0; i < len(s); i++ {
  for j := 0; j < len(s); j++ {
    s[j] += s[i]
    n++
  }
}
n
`,
			EngineStarlark: `
s = []
for i in range(100):
    s.append(i)
n = 0
for i in range(len(s)):
    for j in range(len(s)):
        s[j] = s[j] + s[i]
        n = n + 1
__result = n
`,
		},
	}
}
