# 01. Stable Unique: Slices Without Python Habits

**Time:** 5–10 minutes.

Implement `UniqueStable`: the function takes a slice of strings and returns a new slice with duplicates removed, preserving the order of first occurrence.

## Contract

- `[]string{"go", "py", "go"}` becomes `[]string{"go", "py"}`.
- Comparison is case-sensitive: `"Go"` and `"go"` are different values.
- The input slice must not be modified or reused as the output buffer.
- Return `nil` for a `nil` input. For a non-empty input whose values are all duplicates of one another, return a regular non-empty result.
- Expected complexity: `O(n)` time and `O(n)` additional memory.

Run: `go test ./tasks/01-slices -v`.

Pay attention to the `map[string]struct{}` idiom for representing a set and to preallocating capacity. In Go, an explicit loop is more natural here than trying to imitate a comprehension.

## Going Further

Add `UniqueStableBy`, which accepts a function that computes an element's key, and write table-driven tests.
