# 09. A Generic Set and a Useful Zero Value

**Time:** 15–20 minutes.

Implement a generic `Set[T comparable]` backed by a map.

## Contract

- Methods: `Add`, `Remove`, `Contains`, `Len`, `Values`, and `Clone`.
- The zero value `var s Set[string]` must be ready for `Add` without a constructor.
- Adding the same value again does not increase the length.
- `Values` returns a new slice; its order is unspecified.
- `Clone` is independent: changes to the clone do not affect the original, and vice versa.
- The type must work with any `comparable` type, including structs whose fields are comparable.

Run: `go test ./tasks/09-generic-set -v`.

A nil map in Go can be read, but writing to it causes a panic. A mutating method must therefore initialize the storage lazily. Pay attention to when a pointer receiver is required.

## Going Further

Implement `Union`, `Intersection`, and `Difference` without modifying their operands.
