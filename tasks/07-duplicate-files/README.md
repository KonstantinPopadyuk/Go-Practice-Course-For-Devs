# 07. Finding Duplicate Files

**Time:** 25–30 minutes.

Implement `FindDuplicates(root, minSize)`. Traverse the directory recursively and group identical regular files by SHA-256 hash.

## Contract

- Each result key is a lowercase hexadecimal SHA-256 hash, and its value is a slice of file paths.
- Return only groups containing at least two files.
- Ignore files smaller than `minSize`, directories, and symbolic links.
- Do not read an entire file into memory: stream it into `sha256` with `io.Copy`.
- Paths within each group and the groups themselves must be deterministic. For a map result, sorting each value slice is sufficient.
- A traversal, open, or read error must stop processing and include the affected path.

Run: `go test ./tasks/07-duplicate-files -v`.

Use `filepath.WalkDir` or `fs.WalkDir`. You can first group files by size to avoid hashing files that are definitely unique, but this optimization is optional.

## Going Further

Add a bounded worker pool for hashing and verify it with `go test -race`.
