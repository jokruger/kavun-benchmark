package task

// strContains stresses repeated multi-character substring search.
func strContains() Task {
	return Task{
		Name:        "str_contains",
		Description: "1000 substring searches over a 1000-char haystack, alternating an 8-char hit and miss needle. Stresses substring search.",
		Expected:    500,
		Sources: map[string]string{
			EngineKavun: `
h = "abcdefghijABCDEFGHIJ".repeat(50)
hit = "ijABCDEF"
miss = "ijABCDEZ"
res = 0
for i = 0; i < 1000; i++ {
  if i % 2 == 0 {
    if hit in h { res += 1 }
  } else {
    if miss in h { res += 1 }
  }
}
`,
			EngineTengo: `
text := import("text")
h := text.repeat("abcdefghijABCDEFGHIJ", 50)
hit := "ijABCDEF"
miss := "ijABCDEZ"
res = 0
for i := 0; i < 1000; i++ {
  if i % 2 == 0 {
    if text.contains(h, hit) { res += 1 }
  } else {
    if text.contains(h, miss) { res += 1 }
  }
}
`,
			EngineGoja: `
var h = "abcdefghijABCDEFGHIJ".repeat(50);
var hit = "ijABCDEF";
var miss = "ijABCDEZ";
var res = 0;
for (var i = 0; i < 1000; i++) {
  if (i % 2 === 0) {
    if (h.includes(hit)) { res += 1; }
  } else {
    if (h.includes(miss)) { res += 1; }
  }
}
res;
`,
			EngineGoLua: `
local h = string.rep("abcdefghijABCDEFGHIJ", 50)
local hit = "ijABCDEF"
local miss = "ijABCDEZ"
local res = 0
for i = 0, 999 do
  if i % 2 == 0 then
    if string.find(h, hit, 1, true) then res = res + 1 end
  else
    if string.find(h, miss, 1, true) then res = res + 1 end
  end
end
return res
`,
			EngineGopher: `
local h = string.rep("abcdefghijABCDEFGHIJ", 50)
local hit = "ijABCDEF"
local miss = "ijABCDEZ"
local res = 0
for i = 0, 999 do
  if i % 2 == 0 then
    if string.find(h, hit, 1, true) then res = res + 1 end
  else
    if string.find(h, miss, 1, true) then res = res + 1 end
  end
end
return res
`,
			EngineRisor: `
h := strings.repeat("abcdefghijABCDEFGHIJ", 50)
hit := "ijABCDEF"
miss := "ijABCDEZ"
res := 0
for i := 0; i < 1000; i++ {
  if i % 2 == 0 {
    if strings.contains(h, hit) { res++ }
  } else {
    if strings.contains(h, miss) { res++ }
  }
}
res
`,
			EngineStarlark: `
h = "abcdefghijABCDEFGHIJ" * 50
hit = "ijABCDEF"
miss = "ijABCDEZ"
res = 0
for i in range(1000):
    if i % 2 == 0:
        if hit in h:
            res = res + 1
    else:
        if miss in h:
            res = res + 1
`,
		},
	}
}
