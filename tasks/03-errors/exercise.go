package task03

import "errors"

var ErrOutOfRange = errors.New("port out of range")

// ParsePort parses a decimal TCP port in the inclusive range [1, 65535].
func ParsePort(raw string) (int, error) {
	// TODO: implement.
	return 0, errors.New("TODO")
}
