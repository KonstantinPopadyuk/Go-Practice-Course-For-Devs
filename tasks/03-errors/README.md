# 03. Errors as Values

**Time:** 10–15 minutes.

Implement `ParsePort`: convert a string to a TCP port and return informative errors.

## Contract

- Valid inputs are decimal numbers from 1 through 65535.
- A syntax error must wrap the original `strconv.NumError` so that `errors.As` can find it.
- A number outside the valid range must return the sentinel `ErrOutOfRange`, discoverable through `errors.Is`.
- Every error message must include the original value for diagnostic purposes.
- Do not inspect errors by their text or use `panic`.

Run: `go test ./tasks/03-errors -v`.

Practice using `fmt.Errorf("...: %w", err)`, `errors.Is`, and `errors.As`. `%v` includes the text but breaks the chain of causes.

## Going Further

Add a `RangeError` type with `Value`, `Min`, and `Max` fields while preserving compatibility with `errors.Is(err, ErrOutOfRange)` through an `Unwrap` method.
