package task

// fibTail stresses function-call overhead with linear recursion depth (tail-recursion shape) rather than the
// exponential call tree of naive fib. Same arithmetic result as fib(20) = 6765, but ~O(n) calls instead of O(phi^n).
func fibTail() Task {
	return Task{
		Name:        "fib_tail",
		Description: "Tail-recursion-shaped Fibonacci(20) with accumulator. Stresses function-call overhead at linear depth.",
		Expected:    6765,
		Sources: map[string]string{
			EngineKavun: `
fib = func(x, a, b) {
  if x == 0 { return a }
  if x == 1 { return b }
  return fib(x-1, b, a+b)
}
__result = fib(20, 0, 1)
`,
			EngineTengo: `
fib := func(x, a, b) {
  if x == 0 { return a }
  if x == 1 { return b }
  return fib(x-1, b, a+b)
}
__result = fib(20, 0, 1)
`,
			EngineGoja: `
function fib(x, a, b) {
  if (x === 0) return a;
  if (x === 1) return b;
  return fib(x-1, b, a+b);
}
fib(20, 0, 1);
`,
			EngineGoLua: `
local function fib(x, a, b)
  if x == 0 then return a end
  if x == 1 then return b end
  return fib(x-1, b, a+b)
end
return fib(20, 0, 1)
`,
			EngineGopher: `
local function fib(x, a, b)
  if x == 0 then return a end
  if x == 1 then return b end
  return fib(x-1, b, a+b)
end
return fib(20, 0, 1)
`,
		},
	}
}
