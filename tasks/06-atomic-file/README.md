# 06. Atomic File Write

**Time:** 20–25 minutes.

Implement `WriteFileAtomic(path, data, perm)`: after a successful return, `path` must contain the complete new file, and an observer must never see partially written content.

## Algorithm and Contract

1. Create a temporary file **in the same directory** as the target.
2. Write the data, call `Sync`, close the file, set its permissions, and replace the target with `os.Rename`.
3. If any operation fails, close and remove the temporary file.
4. You do not need to create a missing parent directory.
5. Every error must include the operation and path, and the underlying cause must be wrapped with `%w`.

Using the same directory matters: a rename within one file system is usually atomic. Think carefully about the order of `defer` calls so cleanup does not remove a file that has already been renamed.

Run: `go test ./tasks/06-atomic-file -v`.

## Going Further

Investigate why you may need to call `Sync` on the parent directory after the rename to protect against power loss, and add a platform-specific implementation.
