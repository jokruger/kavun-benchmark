package task

// fib stresses function-call overhead and recursion via naive Fibonacci.
func fib() Task {
	return Task{
		Name:        "fib",
		Description: "Naive recursive Fibonacci(25). Stresses function-call overhead and recursion.",
		Expected:    75025,
		Sources: map[string]string{
			EngineKavun: `
fib = func(n) {
  if n < 2 { return n }
  return fib(n-1) + fib(n-2)
}
res = fib(25)
`,
			EngineTengo: `
fib := func(n) {
  if n < 2 { return n }
  return fib(n-1) + fib(n-2)
}
res = fib(25)
`,
			EngineGoja: `
function fib(n) {
  if (n < 2) return n;
  return fib(n-1) + fib(n-2);
}
fib(25);
`,
			EngineGoLua: `
local function fib(n)
  if n < 2 then return n end
  return fib(n-1) + fib(n-2)
end
return fib(25)
`,
			EngineGopher: `
local function fib(n)
  if n < 2 then return n end
  return fib(n-1) + fib(n-2)
end
return fib(25)
`,
			EngineRisor: `
func fib(n) {
  if n < 2 { return n }
  return fib(n-1) + fib(n-2)
}
fib(25)
`,
			EngineStarlark: `
def fib(n):
    if n < 2: return n
    return fib(n-1) + fib(n-2)
res = fib(25)
`,
		},
	}
}
