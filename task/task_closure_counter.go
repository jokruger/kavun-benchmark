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
last = 0
for i = 0; i < 1000; i++ { last = counter() }
__result = last
`,
			EngineTengo: `
make := func() {
  c := 0
  return func() { c += 1; return c }
}
counter := make()
last := 0
for i := 0; i < 1000; i++ { last = counter() }
__result = last
`,
			EngineGoja: `
function make() { var c = 0; return function() { c += 1; return c; }; }
var counter = make();
var last = 0;
for (var i = 0; i < 1000; i++) { last = counter(); }
last;
`,
		},
	}
}
