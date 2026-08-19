package task15

import "context"

type Group[K comparable, V any] struct {
	// TODO: add synchronization and in-flight call storage.
}

func (g *Group[K, V]) Do(ctx context.Context, key K, fn func() (V, error)) (V, error) {
	// TODO: implement.
	var zero V
	return zero, nil
}
