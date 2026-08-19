# Coming to Go from Other Languages

This is a practical field guide for developers moving to Go from Python, JavaScript/TypeScript, Java/C#, or C++. It focuses on the differences most likely to affect the exercises, not on syntax.

## The Go Mental Model

- **Assignment copies values.** Copying a struct copies its fields. Copying a slice copies its descriptor, so two slices may still share a backing array. Maps also refer to shared runtime state.
- **Zero values are part of API design.** A `nil` slice can be read and appended to. A `nil` map can be read but not written. Prefer types whose zero value is useful when practical.
- **Strings contain bytes.** `len(s)` reports bytes, indexing returns a byte, and `range` decodes UTF-8 code points as runes. A rune is not necessarily a user-perceived character.
- **Errors are values.** Return them, wrap causes with `fmt.Errorf("...: %w", err)`, and inspect the chain with `errors.Is` or `errors.As`. Do not use `panic` for routine validation.
- **Composition replaces class hierarchies.** Types satisfy interfaces implicitly by implementing their methods. Small interfaces such as `io.Reader` make code easier to reuse and test.
- **Cleanup is explicit.** Garbage collection manages memory, not files, sockets, or goroutines. Use `defer` after successfully acquiring a resource.
- **Concurrency needs ownership.** Define how every goroutine stops, who closes each channel, and how cancellation reaches every blocking operation. Use bounded concurrency for external resources and test with `go test -race`.

## If You Come from Python

- A slice is not an independent Python list: two slices can expose the same array, and `append` may reuse that array or allocate a new one.
- Explicit loops are often clearer and more idiomatic than recreating list comprehensions through helpers.
- Replace exception-driven control flow with explicit `value, err` handling and preserve wrapped error causes.
- Interfaces provide compile-time structural behavior; generics express algorithms over related types. Neither is unrestricted duck typing.
- Do not rely on interpreter-level serialization. Goroutines may run simultaneously, so shared memory requires synchronization.

## If You Come from JavaScript or TypeScript

- A goroutine is not an `async` function or Promise. Starting one does not create a result handle; result delivery, cancellation, and completion are explicit.
- Channel sends and receives can block. Use `select` with `ctx.Done()` when an operation must be cancelable.
- Go has no general truthiness or `undefined`. Conditions are boolean, conversions are explicit, and every variable has a zero value.
- Interfaces are checked by the compiler and do not describe arbitrary JSON shapes; decoding and validation remain explicit.
- Goroutines may execute in parallel, so data races are possible. Run concurrent tests with the race detector.

## If You Come from Java or C#

- Go has structs and methods, but no class inheritance hierarchy. Build behavior with composition and small interfaces.
- Struct values are copied unless you use pointers. Do not assume Java-style reference semantics for ordinary structs.
- Constructors such as `NewClient` are conventions, not language features, and are unnecessary when the zero value already works.
- Expected failures are returned as errors rather than thrown as exceptions. `defer` often handles nearby cleanup that would otherwise live in `finally`.
- Goroutines are lighter than platform threads, but they are not free. Bound concurrency, propagate cancellation, and wait for every goroutine to finish.

## If You Come from C++

- Go uses garbage collection and does not expose pointer arithmetic, but files, sockets, and goroutines still need explicit lifecycle management.
- Go has no deterministic destructors tied to lexical scope. Use `defer` for function-scoped cleanup; garbage collection does not close resources for you.
- Slices and maps are not direct equivalents of STL containers. Learn slice capacity, shared backing arrays, map iteration rules, and concurrency constraints.
- Generics are intentionally smaller than C++ templates: there is no template specialization, user-defined operator overloading, or inheritance-based generic design.
- Return and wrap errors instead of relying on exceptions, and use the race detector in addition to reasoning about locks and lifetimes.

## Where to Start

| Background | Suggested first exercises |
|---|---|
| Python | [Slices](../tasks/01-slices), [Unicode](../tasks/02-unicode), [Errors](../tasks/03-errors), [Generics](../tasks/09-generic-set) |
| JavaScript/TypeScript | [Errors](../tasks/03-errors), [JSON](../tasks/05-json-config), [HTTP client](../tasks/10-http-client), [Cancellable pipeline](../tasks/12-cancellable-pipeline) |
| Java/C# | [Errors](../tasks/03-errors), [Interfaces and I/O](../tasks/04-stream-io), [Generics](../tasks/09-generic-set), [TTL cache](../tasks/14-ttl-cache) |
| C++ | [Slices](../tasks/01-slices), [Atomic files](../tasks/06-atomic-file), [Worker pool](../tasks/11-worker-pool), [Graceful shutdown](../tasks/18-graceful-shutdown) |

For interview-focused concurrency practice, continue with [exercises 11–20](../tasks/11-worker-pool).
