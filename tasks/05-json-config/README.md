# 05. Strict JSON Configuration

**Time:** 20–25 minutes.

Implement `LoadConfig(r io.Reader)`. Decode JSON into `Config`, reject unknown fields, and validate all values.

## Format

```json
{"address":"127.0.0.1:8080","workers":4,"timeout":"750ms"}
```

`timeout` is a string in the format accepted by `time.ParseDuration`. Implement `Duration.UnmarshalJSON` to support it.

## Contract

- `address` must not be empty.
- `workers` must be in the range 1..64.
- `timeout` must be positive.
- An unknown field is an error (`json.Decoder.DisallowUnknownFields`).
- No second value or trailing garbage may follow the first JSON object.
- Errors must include context about the field or processing stage and preserve the underlying cause where one exists.

Run: `go test ./tasks/05-json-config -v`.

Unlike a dynamic dictionary, a typed configuration documents the format and moves some errors to compile time.

## Going Further

Add default values, but apply them only to fields that are absent, not to fields explicitly set to zero values.
