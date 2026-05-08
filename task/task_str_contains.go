package task

// strContains stresses repeated substring search. Builds a 100-character string of even-coded printable bytes
// (codes 32, 34, ..., 230), then for each code in [32, 232) asks whether that single character is present.
// Even codes hit (100), odd codes miss (100), so the answer is 100. Tengo and Kavun use their `text` stdlib module;
// Goja uses String.prototype.includes.
func strContains() Task {
	return Task{
		Name:        "str_contains",
		Description: "Build a 100-char string of even-coded printable bytes, then probe contains() for each code in [32,232). Stresses substring search.",
		Expected:    100,
		Sources: map[string]string{
			EngineKavun: `
text = import("text")
size = 100
s = ""
for r = 0; r < size*2; r++ {
  if r%2 == 0 { s += string(rune(32+r)) }
}
n = 0
for r = 0; r < size*2; r++ {
  if text.contains(s, string(rune(32+r))) { n += 1 }
}
__result = n
`,
			EngineTengo: `
text := import("text")
size := 100
s := ""
for r := 0; r < size*2; r++ {
  if r%2 == 0 { s += string(char(32+r)) }
}
n := 0
for r := 0; r < size*2; r++ {
  if text.contains(s, string(char(32+r))) { n += 1 }
}
__result = n
`,
			EngineGoja: `
var size = 100;
var s = "";
for (var r = 0; r < size*2; r++) {
  if (r % 2 === 0) { s += String.fromCharCode(32+r); }
}
var n = 0;
for (var r = 0; r < size*2; r++) {
  if (s.includes(String.fromCharCode(32+r))) { n += 1; }
}
n;
`,
			EngineGoLua: `
local size = 100
local s = ""
for r = 0, size*2-1 do
  if r % 2 == 0 then s = s .. string.char(32+r) end
end
local n = 0
for r = 0, size*2-1 do
  if string.find(s, string.char(32+r), 1, true) then n = n + 1 end
end
return n
`,
			EngineGopher: `
local size = 100
local s = ""
for r = 0, size*2-1 do
  if r % 2 == 0 then s = s .. string.char(32+r) end
end
local n = 0
for r = 0, size*2-1 do
  if string.find(s, string.char(32+r), 1, true) then n = n + 1 end
end
return n
`,
		},
	}
}
