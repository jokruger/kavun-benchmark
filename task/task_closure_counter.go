package task

// closureCounter stresses closure creation and repeated invocation.
func closureCounter() Task {
	return Task{
		Name:        "closure_counter",
		Description: "Closure that increments a captured counter, called 1000 times.",
		Expected:    1000,
		Sources: map[string]string{
			EngineKavun: `
make = func() {
  c = 0
  return func() { c += 1; return c }
}
counter = make()
res = 0
for i = 0; i < 1000; i++ { res = counter() }
`,
			EngineTengo: `
make := func() {
  c := 0
  return func() { c += 1; return c }
}
counter := make()
res = 0
for i := 0; i < 1000; i++ { res = counter() }
`,
			EngineGoja: `
function make() {
  var c = 0;
  return function() { c += 1; return c; };
}
var counter = make();
var res = 0;
for (var i = 0; i < 1000; i++) { res = counter(); }
res;
`,
			EngineGoLua: `
local function make()
  local c = 0
  return function() c = c + 1 return c end
end
local counter = make()
local res = 0
for i = 1, 1000 do res = counter() end
return res
`,
			EngineGopher: `
local function make()
  local c = 0
  return function() c = c + 1 return c end
end
local counter = make()
local res = 0
for i = 1, 1000 do res = counter() end
return res
`,
			EngineRisor: `
mk := func() {
  c := 0
  return func() { c++; return c }
}
counter := mk()
res := 0
for i := 0; i < 1000; i++ { res = counter() }
res
`,
			EngineStarlark: `
def make():
    c = [0]
    def inc():
        c[0] += 1
        return c[0]
    return inc
counter = make()
res = 0
for i in range(1000):
    res = counter()
`,
		},
	}
}
